/*Package generatepdf
* Contiene la función para generar el pdf
* */
package generatepdf

import (
	"fmt"
	"io"
	pdfstruct "pdfgen/struct"
	"strconv"
	"strings"

	"codeberg.org/go-pdf/fpdf"
	"github.com/araddon/dateparse"
)

func GenSimplePDF(w io.Writer, input *pdfstruct.SimpleInput) error {
	pdf := fpdf.New("P", "mm", "A4", "")
	var fontPath string
	fontDesired := strings.ToLower(input.Options.Font)

	if fontDesired != "arial" {
		fontPath = "./fonts/times.ttf"
	} else {
		fontPath = "./fonts/arial.ttf"
	}

	pdf.SetMargins(float64(input.Options.Margin), float64(input.Options.Margin), float64(input.Options.Margin))
	pdf.AddPage()
	pdf.AddUTF8Font(fontDesired, "", fontPath)
	pdf.SetFont(fontDesired, "B", 16)
	pdf.CellFormat(40, 10, input.Title, "0", 1, "", false, 0, "")
	pdf.SetFont(fontDesired, "", 12)
	pdf.MultiCell(0, 6, input.Content, "0", "J", false)
	err := pdf.Output(w)
	if err != nil {
		return fmt.Errorf("ocurrió un error al escrbir el pdf: %v", err)
	}

	return nil
}

func GenInvoicePDF(w io.Writer, input *pdfstruct.InvoiceInput) error {
	pdf := fpdf.New("P", "mm", "A4", "")
	fontPath := "./fonts/arial.ttf"
	marginLeft := 15.0
	marginRight := 15.0
	marginTop := 15.0
	marginBottom := 15.0
	pdf.SetMargins(marginLeft, marginTop, marginRight)
	pdf.AddPage()
	pdf.AddUTF8Font("Arial", "", fontPath)
	pdf.SetAutoPageBreak(true, marginBottom)

	pageWidth, _ := pdf.GetPageSize()
	contentWidth := pageWidth - marginLeft - marginRight

	pdf.SetFont("Arial", "B", 12)

	companyInfoWidth := contentWidth * 0.6
	invoiceInfoWidth := contentWidth * 0.4

	pdf.CellFormat(companyInfoWidth, 7, input.Company, "0", 0, "L", false, 0, "")
	pdf.CellFormat(invoiceInfoWidth, 7, fmt.Sprintf("Factura Nro. %d", input.InvoiceNr), "0", 1, "R", false, 0, "")

	pdf.SetFont("Arial", "", 10)

	parsedDate, err := dateparse.ParseAny(input.Date)
	if err != nil {

		fmt.Println("error al parsear la fecha")
		return fmt.Errorf("error al parsear la fecha '%s': %w", input.Date, err)
	}
	const pdfDateFormat = "02/01/2006"
	dateText := "Fecha: " + parsedDate.Format(pdfDateFormat)

	pdf.CellFormat(companyInfoWidth, 6, input.Rif, "0", 0, "L", false, 0, "")
	pdf.CellFormat(invoiceInfoWidth, 6, dateText, "0", 1, "R", false, 0, "")

	pdf.MultiCell(contentWidth, 6, "DIRECCIÓN: "+input.Direction+", "+input.State+".", "0", "L", false)
	pdf.CellFormat(0, 6, "TELÉFONO:", "0", 1, "L", false, 0, "")

	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(0, 8, "MOSTRADOR", "0", 1, "L", false, 0, "")
	pdf.Ln(2)

	pdf.SetFont("Arial", "", 10)
	clientLabelWidth := 25.0
	clientDataWidth := pageWidth - marginLeft - marginRight - clientLabelWidth

	pdf.CellFormat(clientLabelWidth, 6, "CLIENTE:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(clientDataWidth, 6, input.RifClient, "0", 1, "L", false, 0, "")

	pdf.CellFormat(clientLabelWidth, 6, "NOMBRE:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(clientDataWidth, 6, input.CompanyClient, "0", 1, "L", false, 0, "")

	pdf.CellFormat(clientLabelWidth, 6, "DIRECCIÓN:", "0", 0, "L", false, 0, "")
	pdf.MultiCell(clientDataWidth, 6, input.DirectionClient+".", "0", "L", false)

	pdf.CellFormat(clientLabelWidth, 6, "TELÉFONO:", "0", 1, "L", false, 0, "")

	pdf.Ln(10)

	colWidths := map[string]float64{
		"CODIGO":      18,
		"DESCRIPCION": 75,
		"CANTIDAD":    17,
		"PRECIO":      25,
		"DESCUENTO":   20,
		"TOTAL":       25,
	}
	headerHeight := 7.0
	rowHeight := 6.0
	border := "1"
	startX := pdf.GetX()

	pdf.SetFont("Arial", "B", 8)
	pdf.SetFillColor(220, 220, 220)
	fillHeader := true
	headers := []string{"CODIGO", "DESCRIPCION", "CANTIDAD", "PRECIO", "DESCUENTO", "TOTAL"}
	for _, header := range headers {
		pdf.CellFormat(colWidths[header], headerHeight, header, border, 0, "C", fillHeader, 0, "")
	}
	pdf.Ln(headerHeight)

	pdf.SetFont("Arial", "", 8)
	fillRow := false
	subtotal := 0.0
	totalDiscount := 0.0

	for _, item := range input.Products {
		pdf.SetX(startX)
		itemTotalRaw := float64(item.Quantity) * item.PriceEachUnit
		itemDiscountAmount := itemTotalRaw * (item.Discount / 100.0)
		itemTotalNet := itemTotalRaw - itemDiscountAmount

		subtotal += itemTotalRaw
		totalDiscount += itemDiscountAmount

		pdf.CellFormat(colWidths["CODIGO"], rowHeight, strconv.Itoa(item.Code), border, 0, "C", fillRow, 0, "")
		pdf.CellFormat(colWidths["DESCRIPCION"], rowHeight, item.Description, border, 0, "L", fillRow, 0, "")
		pdf.CellFormat(colWidths["CANTIDAD"], rowHeight, strconv.Itoa(item.Quantity), border, 0, "C", fillRow, 0, "")
		pdf.CellFormat(colWidths["PRECIO"], rowHeight, fmt.Sprintf("%.2f%s", item.PriceEachUnit, input.Currency), border, 0, "R", fillRow, 0, "")
		pdf.CellFormat(colWidths["DESCUENTO"], rowHeight, fmt.Sprintf("%.0f%%", item.Discount), border, 0, "R", fillRow, 0, "")
		pdf.CellFormat(colWidths["TOTAL"], rowHeight, fmt.Sprintf("%.2f%s", itemTotalNet, input.Currency), border, 1, "R", fillRow, 0, "")
	}

	pdf.Ln(2)
	pdf.SetFont("Arial", "", 9)
	pdf.SetFillColor(240, 240, 240)

	widthBeforeTotals := colWidths["CODIGO"] + colWidths["DESCRIPCION"] + colWidths["CANTIDAD"]

	pdf.CellFormat(widthBeforeTotals, rowHeight, "", "0", 0, "", false, 0, "")
	pdf.CellFormat(colWidths["PRECIO"]+colWidths["DESCUENTO"], rowHeight, "Subtotal:", border, 0, "R", true, 0, "")
	pdf.CellFormat(colWidths["TOTAL"], rowHeight, fmt.Sprintf("%.2f%s", subtotal, input.Currency), border, 1, "R", true, 0, "")

	pdf.CellFormat(widthBeforeTotals, rowHeight, "", "0", 0, "", false, 0, "")
	pdf.CellFormat(colWidths["PRECIO"]+colWidths["DESCUENTO"], rowHeight, "Total Descuento:", border, 0, "R", true, 0, "")
	pdf.CellFormat(colWidths["TOTAL"], rowHeight, fmt.Sprintf("%.2f%s", totalDiscount, input.Currency), border, 1, "R", true, 0, "")

	ivaRate := float64(input.IVA) / 100.0
	ivaAmount := (subtotal - totalDiscount) * ivaRate

	pdf.CellFormat(widthBeforeTotals, rowHeight, "", "0", 0, "", false, 0, "")
	pdf.CellFormat(colWidths["PRECIO"]+colWidths["DESCUENTO"], rowHeight, fmt.Sprintf("IVA (%.0f%%):", ivaRate*100), border, 0, "R", true, 0, "")

	pdf.CellFormat(colWidths["TOTAL"], rowHeight, fmt.Sprintf("%.2f%s", ivaAmount, input.Currency), border, 1, "R", true, 0, "")

	grandTotal := (subtotal - totalDiscount) + ivaAmount
	pdf.SetFont("Arial", "", 10)
	pdf.SetFillColor(190, 220, 255)

	pdf.CellFormat(widthBeforeTotals, rowHeight+2, "", "0", 0, "", false, 0, "")
	pdf.CellFormat(colWidths["PRECIO"]+colWidths["DESCUENTO"], rowHeight+2, "TOTAL A PAGAR:", border, 0, "R", true, 0, "")
	pdf.CellFormat(colWidths["TOTAL"], rowHeight+2, fmt.Sprintf("%.2f%s", grandTotal, input.Currency), border, 1, "R", true, 0, "")

	return pdf.Output(w)
}

/* Package pdfstruct
*	La estructura que debe mandar el frontend para el procesamiento del pdf
* */
package pdfstruct

type Options struct {
	Font   string `json:"font"`
	Margin int    `json:"margin"`
}

type SimpleInput struct {
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Options Options `json:"options"`
}

type Item struct {
	Code          int     `json:"code"`
	Description   string  `json:"description"`
	Quantity      int     `json:"quantity"`
	PriceEachUnit float64 `json:"priceEachUnit"`
	Discount      float64 `json:"discount"`
}

type InvoiceInput struct {
	Company         string `json:"company"`
	Rif             string `json:"rif"`
	Direction       string `json:"direction"`
	State           string `json:"state"`
	InvoiceNr       int    `json:"invoiceNr"`
	Date            string `json:"date"`
	RifClient       string `json:"rifClient"`
	CompanyClient   string `json:"companyClient"`
	DirectionClient string `json:"directionClient"`
	Products        []Item `json:"products"`
	Currency        string `json:"currency"`
	IVA             int    `json:"iva"`
}

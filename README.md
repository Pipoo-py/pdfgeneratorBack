# izpdfgenerator - Simple PDF Generator

Una aplicación web Full-Stack con un backend de alto rendimiento en Go y un frontend reactivo en Astro + Svelte para generar PDFs a partir de plantillas.

**[➡️ Ver Demo en Vivo](https://izpdfgenerator.netlify.app/)**

---

### 🇬🇧 About This Project

izpdfgenerator is a web service that allows users to generate PDF documents from two predefined templates: a basic note and a professional invoice. This project demonstrates my ability to build a complete Full-Stack application, with a strong focus on backend performance, API design, and a seamless user experience on the front end.

#### ✨ Features

* **High-Performance Go Backend:** A robust API built in Go to handle PDF generation efficiently.
* **Two PDF Templates:**
    1.  **Basic:** A simple document with a title and content body.
    2.  **Invoice:** A B2B invoice with fields for company, client, items, and totals.
* **Reactive Frontend:** A clean and intuitive UI built with Astro and Svelte that communicates with the backend API.

#### 🛠️ Architecture & Tech Stack

The application is architected as a decoupled frontend and backend.

**Backend:**
<p>
  <img alt="Go" src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"/>
</p>

**Frontend:**
<p>
  <img alt="Astro" src="https://img.shields.io/badge/Astro-FF5D01?style=for-the-badge&logo=astro&logoColor=white"/>
  <img alt="Svelte" src="https://img.shields.io/badge/Svelte-FF3E00?style=for-the-badge&logo=svelte&logoColor=white"/>
  <img alt="JavaScript" src="https://img.shields.io/badge/JavaScript-007ACC?style=for-the-badge&logo=javascript&logoColor=white"/>
</p>

**Philosophy:** **Go** was chosen for the backend due to its incredible performance, low memory footprint, and first-class support for concurrency. This makes it ideal for a service that performs I/O-bound tasks like generating a file. The result is a fast, reliable, and efficient API that can handle many requests with minimal server resources.
---
🇪🇸 Sobre el Proyecto
izpdfgenerator es un servicio web que permite a los usuarios generar documentos PDF a partir de dos plantillas predefinidas: una nota básica y una factura profesional. Este proyecto demuestra mi habilidad para construir una aplicación Full-Stack completa, con un fuerte enfoque en el rendimiento del backend, el diseño de APIs y una experiencia de usuario fluida en el frontend.

#### ✨ Características
 * **Backend de Alto Rendimiento en Go**: Una API robusta construida en Go para manejar la generación de PDFs de manera eficiente.
* **Dos Plantillas de PDF**:
    1. Básica: Un documento simple con título y cuerpo de contenido.
    2. Factura: Una factura B2B con campos para emisor, cliente, ítems y totales.
* **Frontend Reactivo**: Una UI limpia e intuitiva construida con Astro y Svelte que se comunica con la API del backend.

#### 🛠️ Arquitectura y Stack Tecnológico
La aplicación está diseñada con un frontend y un backend desacoplados.

* **Backend**: Go
* **Frontend**: Astro, Svelte, TypeScript

### Filosofía: Se eligió Go para el backend por su increíble rendimiento, bajo consumo de memoria y soporte nativo de primer nivel para la concurrencia. Esto lo hace ideal para un servicio que realiza tareas ligadas a I/O como la generación de un archivo. El resultado es una API rápida, fiable y eficiente que puede manejar múltiples peticiones con mínimos recursos de servidor.

Made with ❤️ by Lizardiz Nexus.

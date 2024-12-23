package main

import (
	"fmt"
	"generatorPDF/htmlParser"
	"generatorPDF/pdfGenerator"
)

type Data struct {
	Name string
}

func main() {

	h := htmlParser.New("tmp")
	wk := pdfGenerator.NewWkHtmlToPdf("tmp")

	dataHTML := Data{
		Name: "Spock",
	}

	htmlGenerated, err := h.Create("templates/example.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("HTML Generated: ", htmlGenerated)

	filePDFName, err := wk.Create(htmlGenerated)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("PDF Generated: ", filePDFName)
}

package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	_ "github.com/mattn/go-sqlite3"
)

func GeneratePDF(filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.CellFormat(190, 7, "BÁO CÁO KẾT QUẢ ĐẦU TƯ", "0", 0, "CM", false, 0, "")
	fmt.Println("BÁO CÁO KẾT QUẢ ĐẦU TƯ")
	return pdf.OutputFileAndClose(filename)

}

func main() {
	err := GeneratePDF("hello.pdf")
	if err != nil {
		panic(err)
	}
}

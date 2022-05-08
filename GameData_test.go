package main

import (
	"github.com/xuri/excelize/v2"
	"log"
	"testing"
)

//tests if the data returned from getExcelName is correct
func TestGetExcelNameContent(t *testing.T) {
	excelFile, err := excelize.OpenFile("games-features.xlsx")
	if err != nil {
		log.Fatalln(err)
	}
	all_rows, err := excelFile.GetRows("games-features")
	if err != nil {
		log.Fatalln(err)
	}
	nameRows := getExcelName(all_rows)
	nameSlice := sanitizeData(nameRows)

	firstName := nameSlice[1]
	if firstName != "Counter-Strike" {
		t.Error("Test Failed: name slice does not have correct data")
	}
}

//tests if getExcelName returns all required data
func TestGetExcelNameSize(t *testing.T) {
	excelFile, err := excelize.OpenFile("games-features.xlsx")
	if err != nil {
		log.Fatalln(err)
	}
	all_rows, err := excelFile.GetRows("games-features")
	if err != nil {
		log.Fatalln(err)
	}
	nameRows := getExcelName(all_rows)
	nameSlice := sanitizeData(nameRows)

	size := len(nameSlice)
	if size != 13357 {
		t.Error("Test Failed: name slice is not correct size")
	}
}

//tests if the data returned from getExcelMetacritic is correct
func TestGetExcelMetacriticContent(t *testing.T) {
	excelFile, err := excelize.OpenFile("games-features.xlsx")
	if err != nil {
		log.Fatalln(err)
	}
	all_rows, err := excelFile.GetRows("games-features")
	if err != nil {
		log.Fatalln(err)
	}
	metaRows := getExcelMetacritic(all_rows)
	metaSlice := sanitizeData(metaRows)

	firstName := metaSlice[1]
	if firstName != "88" {
		t.Error("Test Failed: metacritic slice does not have correct data")
	}
}

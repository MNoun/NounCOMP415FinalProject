/************************
Created by: Mitchell Noun
Date created: 4/29/22
Class: COMP415 Emerging Languages
Assignment: Final Project: Project 4
*************************/
package main

import (
	"database/sql"
	"github.com/xuri/excelize/v2"
	"log"
)

func main() {
	//opens Excel file and gets all data
	excelFile, err := excelize.OpenFile("games-features.xlsx")
	if err != nil {
		log.Fatalln(err)
	}
	all_rows, err := excelFile.GetRows("")
	if err != nil {
		log.Fatalln(err)
	}

	//Excel data function calls

	//database function calls

	//GUI function calls
}

//creates database
func OpenDatabase(dbfile string) *sql.DB {
	database, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
	}
	return database
}

//creates database table
func tableSetup() {

}

//run once, inserts excel data into database
func populateDatabase() {

}

/*
Below are separate functions to get required excel data
*/

func getExcelName() {

}

func getExcelAge() {

}

func getExcelDLC() {

}

func getExcelMetacritic() {

}

func getExcelRecCount() {

}

func getExcelSteamOwners() {

}

func getExcelSteamPlayers() {

}

func getExcelPLinux() {

}

func getExcelPMac() {

}

func getExcelPWindows() {

}

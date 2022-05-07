/************************
Created by: Mitchell Noun
Date created: 4/29/22
Class: COMP415 Emerging Languages
Assignment: Final Project: Project 4
*************************/
package main

import (
	"database/sql"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
)

func main() {
	//opens Excel file and gets all data
	excelFile, err := excelize.OpenFile("games-features.xlsx")
	if err != nil {
		log.Fatalln(err)
	}
	all_rows, err := excelFile.GetRows("games-features")
	if err != nil {
		log.Fatalln(err)
	}

	//Excel data function calls
	nameRows := getExcelName(all_rows)
	nameSlice := sanitizeData(nameRows)

	ageRows := getExcelAge(all_rows)
	ageSlice := sanitizeData(ageRows)

	dlcRows := getExcelDLC(all_rows)
	dlcSlice := sanitizeData(dlcRows)

	metaRows := getExcelMetacritic(all_rows)
	metaSlice := sanitizeData(metaRows)

	recCountRows := getExcelRecCount(all_rows)
	recCountSlice := sanitizeData(recCountRows)

	steamOwnersRows := getExcelSteamOwners(all_rows)
	steamOwnersSlice := sanitizeData(steamOwnersRows)

	steamPlayersRows := getExcelSteamPlayers(all_rows)
	steamPlayersSlice := sanitizeData(steamPlayersRows)

	pLinuxRows := getExcelPLinux(all_rows)
	pLinuxSlice := sanitizeData(pLinuxRows)

	pMacRows := getExcelPMac(all_rows)
	pMacSlice := sanitizeData(pMacRows)

	pWindowsRows := getExcelPWindows(all_rows)
	pWindowsSlice := sanitizeData(pWindowsRows)

	//database function calls
	gameDatabase := OpenDatabase("games-features.db")
	tableSetup(gameDatabase)

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
func tableSetup(gameDatabase *sql.DB) {
	createStatement := "CREATE TABLE IF NOT EXISTS GameFeatures(    " +
		"name TEXT PRIMARY KEY," +
		"age TEXT NOT NULL," +
		"dlc TEXT NOT NULL," +
		"metacritic TEXT NOT NULL," +
		"recCount TEXT NOT NULL," +
		"steamOwners TEXT NOT NULL," +
		"steamPlayers TEXT NOT NULL," +
		"platformLinux TEXT NOT NULL," +
		"platformMac TEXT NOT NULL," +
		"platformWindows TEXT NOT NULL);"
	_, err := gameDatabase.Exec(createStatement)
	if err != nil {
		log.Println(err)
	}
}

//run once, inserts excel data into database
func populateDatabase() {

}

func sanitizeData(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

/*
Below are separate functions to get required excel data
*/

func getExcelName(r [][]string) []string {
	nameRows := make([]string, 13358)

	for _, row := range r {
		temp_string := fmt.Sprintln(row[2])
		temp_slice := strings.Split(temp_string, "\n")
		for _, s := range temp_slice {
			nameRows = append(nameRows, s) //returns slice of the name column
		}
	}
	return nameRows
}

func getExcelAge(r [][]string) []string {
	ageRows := make([]string, 13358)

	for _, row := range r {
		temp_string := fmt.Sprintln(row[5])
		temp_slice := strings.Split(temp_string, "\n")
		for _, s := range temp_slice {
			ageRows = append(ageRows, s) //returns slice of the age column
		}
	}
	return ageRows
}

func getExcelDLC(r [][]string) []string {
	dlcRows := make([]string, 13358)

	for _, row := range r {
		temp_string := fmt.Sprintln(row[8])
		temp_slice := strings.Split(temp_string, "\n")
		for _, s := range temp_slice {
			dlcRows = append(dlcRows, s) //returns slice of the DLCCount column
		}
	}
	return dlcRows
}

func getExcelMetacritic(r [][]string) []string {
	metaRows := make([]string, 13358)

	for _, row := range r {
		temp_string := fmt.Sprintln(row[9])
		temp_slice := strings.Split(temp_string, "\n")
		for _, s := range temp_slice {
			metaRows = append(metaRows, s) //returns slice of the Metacritic column
		}
	}
	return metaRows
}

func getExcelRecCount(r [][]string) []string {
	recCountRows := make([]string, 13358)

	for _, row := range r {
		temp_string := fmt.Sprintln(row[12])
		temp_slice := strings.Split(temp_string, "\n")
		for _, s := range temp_slice {
			recCountRows = append(recCountRows, s) //returns slice of the RecommendationCount column
		}
	}
	return recCountRows
}

func getExcelSteamOwners(r [][]string) []string {
	steamOwnersRows := make([]string, 13358)

	for _, row := range r {
		temp_string := fmt.Sprintln(row[15])
		temp_slice := strings.Split(temp_string, "\n")
		for _, s := range temp_slice {
			steamOwnersRows = append(steamOwnersRows, s) //returns slice of the SteamSpyOwners column
		}
	}
	return steamOwnersRows
}

func getExcelSteamPlayers(r [][]string) []string {
	steamPlayersRows := make([]string, 13358)

	for _, row := range r {
		temp_string := fmt.Sprintln(row[17])
		temp_slice := strings.Split(temp_string, "\n")
		for _, s := range temp_slice {
			steamPlayersRows = append(steamPlayersRows, s) //returns slice of the SteamSpyPlayersEstimate column
		}
	}
	return steamPlayersRows
}

func getExcelPLinux(r [][]string) []string {
	pLinuxRows := make([]string, 13358)

	for _, row := range r {
		temp_string := fmt.Sprintln(row[27])
		temp_slice := strings.Split(temp_string, "\n")
		for _, s := range temp_slice {
			pLinuxRows = append(pLinuxRows, s) //returns slice of the PlatformLinux column
		}
	}
	return pLinuxRows
}

func getExcelPMac(r [][]string) []string {
	pMacRows := make([]string, 13358)

	for _, row := range r {
		temp_string := fmt.Sprintln(row[28])
		temp_slice := strings.Split(temp_string, "\n")
		for _, s := range temp_slice {
			pMacRows = append(pMacRows, s) //returns slice of the PlatformMac column
		}
	}
	return pMacRows
}

func getExcelPWindows(r [][]string) []string {
	pWindowsRows := make([]string, 13358)

	for _, row := range r {
		temp_string := fmt.Sprintln(row[26])
		temp_slice := strings.Split(temp_string, "\n")
		for _, s := range temp_slice {
			pWindowsRows = append(pWindowsRows, s) //returns slice of the PlatformWindows column
		}
	}
	return pWindowsRows
}

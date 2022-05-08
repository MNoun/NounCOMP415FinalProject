/************************
Created by: Mitchell Noun
Date created: 4/29/22
Class: COMP415 Emerging Languages
Assignment: Final Project: Project 4
*************************/
package main

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xuri/excelize/v2"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"image/png"
	"log"
	"strings"
)

//go:embed graphics/*
var EmbeddedAssets embed.FS
var gameApp GuiApp

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

	print(all_rows)

	//Excel data function calls, run once
	//nameRows := getExcelName(all_rows)
	//nameSlice := sanitizeData(nameRows)
	//
	//ageRows := getExcelAge(all_rows)
	//ageSlice := sanitizeData(ageRows)
	//
	//dlcRows := getExcelDLC(all_rows)
	//dlcSlice := sanitizeData(dlcRows)
	//
	//metaRows := getExcelMetacritic(all_rows)
	//metaSlice := sanitizeData(metaRows)
	//
	//recCountRows := getExcelRecCount(all_rows)
	//recCountSlice := sanitizeData(recCountRows)
	//
	//steamOwnersRows := getExcelSteamOwners(all_rows)
	//steamOwnersSlice := sanitizeData(steamOwnersRows)
	//
	//steamPlayersRows := getExcelSteamPlayers(all_rows)
	//steamPlayersSlice := sanitizeData(steamPlayersRows)
	//
	//pLinuxRows := getExcelPLinux(all_rows)
	//pLinuxSlice := sanitizeData(pLinuxRows)
	//
	//pMacRows := getExcelPMac(all_rows)
	//pMacSlice := sanitizeData(pMacRows)
	//
	//pWindowsRows := getExcelPWindows(all_rows)
	//pWindowsSlice := sanitizeData(pWindowsRows)
	//
	////database function calls
	//gameDatabase := OpenDatabase("./games-features.db")
	//tableSetup(gameDatabase) //run once
	//populateDatabase(nameSlice, ageSlice, dlcSlice, metaSlice, recCountSlice, steamOwnersSlice, steamPlayersSlice,
	//	pLinuxSlice, pMacSlice, pWindowsSlice, gameDatabase) //run once

	//GUI function calls
	ebiten.SetWindowSize(900, 800)
	ebiten.SetWindowTitle("Game Search")

	gameApp = GuiApp{AppUI: MakeUIWindow()}

	err = ebiten.RunGame(&gameApp)
	if err != nil {
		log.Fatalln("Error running User Interface", err)
	}

}

func MakeUIWindow() (GUIhandler *ebitenui.UI) {

	background := image.NewNineSliceColor(color.Gray16{})
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true, false}),
			widget.GridLayoutOpts.Padding(widget.Insets{
				Top:    20,
				Bottom: 20,
			}),
			widget.GridLayoutOpts.Spacing(0, 20))),
		widget.ContainerOpts.BackgroundImage(background))

	//search game button
	idle, err := loadImageNineSlice("button-idle.png", 20, 0)
	if err != nil {
		log.Fatalln(err)
	}
	hover, err := loadImageNineSlice("button-hover.png", 20, 0)
	if err != nil {
		log.Fatalln(err)
	}
	pressed, err := loadImageNineSlice("button-pressed.png", 20, 0)
	if err != nil {
		log.Fatalln(err)
	}
	disabled, err := loadImageNineSlice("button-disabled.png", 20, 0)
	if err != nil {
		log.Fatalln(err)
	}
	buttonImage := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}
	button := widget.NewButton(
		widget.ButtonOpts.Image(buttonImage),
		widget.ButtonOpts.Text("Search Game", basicfont.Face7x13, &widget.ButtonTextColor{
			Idle: color.RGBA{0xdf, 0xf4, 0xff, 0xff},
		}),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  30,
			Right: 30,
		}),
		widget.ButtonOpts.ClickedHandler(loadGameData),
	)
	rootContainer.AddChild(button)

	GUIhandler = &ebitenui.UI{Container: rootContainer}
	return GUIhandler
}

func (g GuiApp) Update() error {
	g.AppUI.Update()
	return nil
}

func (g GuiApp) Draw(screen *ebiten.Image) {
	g.AppUI.Draw(screen)
}

func (g GuiApp) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

type GuiApp struct {
	AppUI *ebitenui.UI
}

func loadImageNineSlice(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	i := loadPNGImageFromEmbedded(path)

	w, h := i.Size()
	return image.NewNineSlice(i,
			[3]int{(w - centerWidth) / 2, centerWidth, w - (w-centerWidth)/2 - centerWidth},
			[3]int{(h - centerHeight) / 2, centerHeight, h - (h-centerHeight)/2 - centerHeight}),
		nil
}

func loadPNGImageFromEmbedded(name string) *ebiten.Image {
	pictNames, err := EmbeddedAssets.ReadDir("graphics")
	if err != nil {
		log.Fatal("failed to read embedded dir ", pictNames, " ", err)
	}
	embeddedFile, err := EmbeddedAssets.Open("graphics/" + name)
	if err != nil {
		log.Fatal("failed to load embedded image ", embeddedFile, err)
	}
	rawImage, err := png.Decode(embeddedFile)
	if err != nil {
		log.Fatal("failed to load embedded image ", name, err)
	}
	gameImage := ebiten.NewImageFromImage(rawImage)
	return gameImage
}

func searchGame(searchTerm string) string {
	gameDatabase, err := sql.Open("sqlite3", "games-features.db")
	if err != nil {
		log.Fatal(err)
	}
	var gameData string
	selectStatement := "SELECT * FROM GameFeatures WHERE name = ?"
	row, err := gameDatabase.Query(selectStatement, searchTerm)
	if err != nil {
		log.Fatalln(err)
	}
	err = row.Scan(&gameData)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}
	return gameData
}

func loadGameData(args *widget.ButtonClickedEventArgs) {

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
func populateDatabase(nameSlice []string, ageSlice []string, dlcSlice []string, metaSlice []string,
	recCountSlice []string, steamOwnersSlice []string, steamPlayersSlice []string, pLinuxSlice []string,
	pMacSlice []string, pWindowsSlice []string, gameDatabase *sql.DB) {
	insertStatement := "INSERT INTO GameFeatures (name, age, dlc, metacritic, recCount, steamOwners, steamPlayers," +
		" platformLinux, platformMac, platformWindows) VALUES (?,?,?,?,?,?,?,?,?,?)"
	for i := 1; i < 13357; i++ {
		preppedStatement, err := gameDatabase.Prepare(insertStatement)
		if err != nil {
			log.Fatal(err)
		}
		preppedStatement.Exec(nameSlice[i], ageSlice[i], dlcSlice[i], metaSlice[i], recCountSlice[i],
			steamOwnersSlice[i], steamPlayersSlice[i], pLinuxSlice[i], pMacSlice[i], pWindowsSlice[i])
	}
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

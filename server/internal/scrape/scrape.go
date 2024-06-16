package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocolly/colly"
)

type Team struct {
	Name string
	Link string
}

var Teams []Team

const targetIndex = 0

var allTeams = make([][]string, 0)

var header bool = true

// TODO: improve Code Structure and Readability:
// TODO: Error Handling:
// TODO: Optimize CSV Handling:
// TODO: Use of Go Concurrency:
// TODO: Refactor Repeated Code:

func Scrapper() {
	getTeamsUrl()

	for _, team := range Teams {
		fmt.Println(team.Name)
		teamName := strings.Split(team.Name, "-Stats")[0]
		getTeamData(teamName, team.Link)
	}

	writeToCSV(allTeams)

	removeColumns()
}

func getTeamsUrl() {
	c := colly.NewCollector(
		colly.AllowedDomains("fbref.com"),
	)

	c.OnHTML("table.stats_table", func(h *colly.HTMLElement) {
		if h.Index == targetIndex {
			links := h.ChildAttrs("a", "href")
			filterTeamLink(&links)

			h.Request.Abort()
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://fbref.com/en/comps/9/Premier-League-Stats")
}

func filterTeamLink(links *[]string) {
	for i := 0; i < len(*links); {
		if !strings.Contains((*links)[i], "/squads/") {
			*links = append((*links)[:i], (*links)[i+1:]...)
		} else {
			team := Team{
				Name: strings.Split((*links)[i], "/")[len(strings.Split((*links)[i], "/"))-1],
				Link: (*links)[i],
			}

			Teams = append(Teams, team)
			i++
		}
	}
}

func getTeamData(teamName string, link string) {
	c := colly.NewCollector(colly.AllowedDomains("fbref.com"))

	c.OnHTML("table.stats_table", func(h *colly.HTMLElement) {
		if h.Index == targetIndex {
			rows := make([]string, 0)
			isFirstRow := true
			h.ForEach("tr", func(_ int, row *colly.HTMLElement) {
				//fmt.Println("tr row colly", row)

				if !isFirstRow {
					rowData := make([]string, 0)
					row.ForEach("th, td", func(_ int, cell *colly.HTMLElement) {
						//fmt.Println("th td cell colly", cell)
						//fmt.Println("cell", cell.Text)
						if strings.Contains(cell.Text, "-") {
							cell.Text = strings.Split(cell.Text, "-")[0]
						}
						rowData = append(rowData, strings.ReplaceAll(cell.Text, ",", ""))
						//fmt.Println("yyyyyyyyyyyyyy", y, "yyyyyyyyyyyyyy")
					})
					//fmt.Println("row Data", rowData)
					if rowData[0] == "Player" && header {
						rowData = append(rowData, "Team")
						rows = append(rows, strings.Join(rowData, ","))
						//fmt.Println("xxxxxxxxxxx", x, "xxxxxxxxxxxxx")
						header = false
					}
					if rowData[0] != "Player" {
						rowData = append(rowData, teamName)
						rows = append(rows, strings.Join(rowData, ","))
					}
				} else {
					isFirstRow = false
				}
			})
			// fmt.Println("rows", rows)
			allTeams = append(allTeams, rows)

			h.Request.Abort()
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://fbref.com" + link)

	// writeToCSV(allTeams)
}

func removeColumns() {
	csvfile, err := os.Open("data/stats.csv")
	if err != nil {
		fmt.Println("Could not open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	newfile, err := os.Create("data/stats.csv")
	if err != nil {
		fmt.Println("Error while creating csv file", err)
	}

	defer newfile.Close()

	w := csv.NewWriter(newfile)
	defer w.Flush()

	for _, record := range records {
		record = append(record[:20], record[34:]...)
		record = append(record[:7], record[8:]...)
		record = append(record[:9], record[11:]...)
		record = append(record[:10], record[11:]...)
		record = append(record[:13], record[14:]...)
		record = append(record[:14], record[15:]...)

		if err := w.Write(record); err != nil {
			fmt.Println("Error while writing record to csv:", err)
		}
	}
}

func writeToCSV(allTeams [][]string) {
	dataDir := "data"

	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		err := os.MkdirAll(dataDir, 0755)
		if err != nil {
			log.Fatal("Error creating directory:", err)
		}
	}

	filePath := filepath.Join(dataDir, "stats.csv")
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Error creating CSV file:", err)
	}

	defer file.Close()

	//file, err := os.Create(fileName + ".csv")
	//if err != nil {
	//log.Fatal("Error creating CSV file:", err)
	//}
	//defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, rows := range allTeams {
		for _, row := range rows {
			writer.Write(strings.Split(row, ","))
		}
	}

	fmt.Println(filePath)
	fmt.Println("Data written to stats.csv")
}

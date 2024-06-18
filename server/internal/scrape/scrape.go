package scrape

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

type Team struct {
	Name string
	Link string
}

const (
	targetIndex = 0
	baseURL     = "https://fbref.com"
)

var (
	Teams    []Team
	allTeams = make([][]string, 0)
	header   = true
	dataDir  = "data"
	filePath = filepath.Join(dataDir, "stats.csv")
	mu       sync.Mutex
)

// TODO: Error Handling:

func Scrapper() {

    err := ensureDataDir()
    if err != nil {
        log.Fatalf("Error ensuring data directory: %v", err)
    }

	getTeamsUrl()

	var wg sync.WaitGroup
	for _, team := range Teams {
		wg.Add(1)

		go func(team Team) {
			fmt.Println(team.Name)
			teamName := strings.Split(team.Name, "-Stats")[0]
			getTeamData(teamName, team.Link)
		}(team)
	}
    wg.Wait()

	removeColumns()
}

func ensureDataDir() error {
    if _, err := os.Stat(dataDir); os.IsNotExist(err) {
        err := os.MkdirAll(dataDir, 0755)
        if err != nil {
            return fmt.Errorf("Error creating directory: %v", err)
        }
    }

    file, err := os.Create(filePath)
    if err != nil {
        return fmt.Errorf("Error creating CSV file: %v", err)
    }
    defer file.Close()

    return nil
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

func getTeamData(teamName, link string) {
	c := colly.NewCollector(colly.AllowedDomains("fbref.com"))

	c.OnHTML("table.stats_table", func(h *colly.HTMLElement) {
		if h.Index == targetIndex {
			rows := [][]string{}
			isFirstRow := true
			h.ForEach("tr", func(_ int, row *colly.HTMLElement) {
				if isFirstRow {
					isFirstRow = false
					return
				}
				rowData := []string{}
				row.ForEach("th, td", func(_ int, cell *colly.HTMLElement) {
					text := strings.Split(cell.Text, "-")[0]
					rowData = append(rowData, strings.ReplaceAll(text, ",", ""))
				})
				if len(rowData) > 0 {
					if rowData[0] != "Player" {
						rowData = append(rowData, teamName)
					}
					rows = append(rows, rowData)
				}
			})
			appendToCSV(rows)
			h.Request.Abort()
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(baseURL + link)
}


func appendToCSV(rows [][]string) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Printf("Appending %d rows to CSV\n", len(rows))

	// Open the CSV file for appending, or create it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush() // Ensure all buffered data is written to the file

	// Write rows to the CSV file
	for _, row := range rows {
		fmt.Printf("Writing row: %v\n", row)
		if err := writer.Write(row); err != nil {
			log.Fatalf("Error writing to CSV file: %v", err)
		}
	}
}

func removeColumns() {
	csvfile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Could not open the CSV file: %v", err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)

	tempFilePath := filepath.Join(dataDir, "temp_stats.csv")
	newfile, err := os.Create(tempFilePath)
	if err != nil {
		log.Fatalf("Error while creating temporary CSV file: %v", err)
	}
	defer newfile.Close()

	w := csv.NewWriter(newfile)
	defer w.Flush()

	for {
		record, err := r.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("Error reading record from CSV: %v", err)
		}

		record = append(record[:20], record[34:]...)
		record = append(record[:7], record[8:]...)
		record = append(record[:9], record[11:]...)
		record = append(record[:10], record[11:]...)
		record = append(record[:13], record[14:]...)
		record = append(record[:14], record[15:]...)
		if err := w.Write(record); err != nil {
			log.Fatalf("Error while writing record to CSV: %v", err)
		}
	}

	if err := os.Rename(tempFilePath, filePath); err != nil {
		log.Fatalf("Error replacing original CSV file: %v", err)
	}
}

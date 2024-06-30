package scrape

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

type Scrapper interface {
	Scrape() error
	Teams() []Team
}

type scrape struct{}

func NewScrape() Scrapper {
	return &scrape{}
}

func (s *scrape) Teams() []Team {
	return Teams
}

type Team struct {
	Name string
	Link string
	Logo string
}

const (
	targetIndex = 0
	baseURL     = "https://fbref.com"
)

var (
	Teams    []Team
	header   = true
	dataDir  = "data"
	filePath = filepath.Join(dataDir, "stats.csv")
	mu       sync.Mutex
)

func (s *scrape) Scrape() error {
	err := ensureDataDir()
	if err != nil {
		return fmt.Errorf("error creating data directory: %v", err)
	}

	err = getTeamsUrl()
	if err != nil {
		return fmt.Errorf("error fetching teams' URLs: %v", err)
	}

	// errCh := make(chan error)
	// var wg sync.WaitGroup

	for _, team := range Teams {
		// wg.Add(1)

		// go func(team Team) {
		// 	defer wg.Done()

		// 	teamName := strings.Split(team.Name, "-Stats")[0]
		// 	if err := getTeamData(teamName, team.Link); err != nil {
		// 		errCh <- fmt.Errorf("error scraping team data for %s: %v", teamName, err)
		// 	}
		// }(team)

		teamName := strings.Split(team.Name, "-Stats")[0]
		if err := getTeamData(teamName, team.Link); err != nil {
			return fmt.Errorf("error scraping team data for %s: %v", teamName, err)
		}
	}

	// go func() {
	// 	wg.Wait()
	// 	close(errCh)
	// }()

	// for err := range errCh {
	// 	log.Printf("Error: %v", err)
	// 	return err
	// }

	err = removeColumns()
	if err != nil {
		return fmt.Errorf("error removing columns: %v", err)
	}

	return nil
}

func ensureDataDir() error {
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		err := os.MkdirAll(dataDir, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating CSV file: %v", err)
	}
	defer file.Close()

	return nil
}

func getTeamsUrl() error {
	c := colly.NewCollector(
		colly.AllowedDomains("fbref.com"),
	)

	c.OnHTML("table.stats_table", func(h *colly.HTMLElement) {
		if h.Index == targetIndex {
			teams := h.ChildAttrs("img", "src")
			links := h.ChildAttrs("a", "href")
			filterTeamLink(&links, &teams)

			h.Request.Abort()
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	if err := c.Visit("https://fbref.com/en/comps/9/Premier-League-Stats"); err != nil {
		return fmt.Errorf("error visiting https://fbref.com/en/comps/9/Premier-League-Stats: %v", err)
	}

	return nil
}

func filterTeamLink(links, teams *[]string) {
	for i := 0; i < len(*links); {
		if !strings.Contains((*links)[i], "/squads/") {
			*links = append((*links)[:i], (*links)[i+1:]...)
		} else {
			team := Team{
				Name: strings.Split((*links)[i], "/")[len(strings.Split((*links)[i], "/"))-1],
				Link: (*links)[i],
				Logo: (*teams)[0],
			}

			Teams = append(Teams, team)
			i++
		}
	}
}

func getTeamData(teamName, link string) error {
	c := colly.NewCollector(colly.AllowedDomains("fbref.com"))

	var scrapeError error
	time.Sleep(5 * time.Second)

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
						rows = append(rows, rowData)
					}

					if rowData[0] == "Player" && header {
						rowData = append(rowData, "Team")
						header = false
						rows = append(rows, rowData)
					}
				}
			})
			if err := appendToCSV(rows); err != nil {
				scrapeError = fmt.Errorf("error appending to CSV: %v", err)
			}
			h.Request.Abort()
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	if err := c.Visit(baseURL + link); err != nil {
		return fmt.Errorf("error visiting team URL: %v", err)
	}

	if scrapeError != nil {
		return scrapeError
	}
	return nil
}

func appendToCSV(rows [][]string) error {
	mu.Lock()
	defer mu.Unlock()

	fmt.Printf("Appending %d rows to CSV\n", len(rows))

	// Open the CSV file for appending, or create it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening CSV file: %v", err)
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush() // Ensure all buffered data is written to the file

	// Write rows to the CSV file
	for _, row := range rows {
		fmt.Printf("Writing row: %v\n", row)
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("error writing to CSV file: %v", err)
		}
	}
	return nil
}

func removeColumns() error {
	csvfile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open the CSV file: %v", err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)

	tempFilePath := filepath.Join(dataDir, "temp_stats.csv")
	newfile, err := os.Create(tempFilePath)
	if err != nil {
		return fmt.Errorf("error while creating temporary CSV file: %v", err)
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
			return fmt.Errorf("error reading record from CSV: %v", err)
		}

		record = append(record[:20], record[34:]...)
		record = append(record[:7], record[8:]...)
		record = append(record[:9], record[11:]...)
		record = append(record[:10], record[11:]...)
		record = append(record[:13], record[14:]...)
		record = append(record[:14], record[15:]...)
		if err := w.Write(record); err != nil {
			return fmt.Errorf("error while writing record to CSV: %v", err)
		}
	}

	if err := os.Rename(tempFilePath, filePath); err != nil {
		return fmt.Errorf("error replacing original CSV file: %v", err)
	}

	return nil
}

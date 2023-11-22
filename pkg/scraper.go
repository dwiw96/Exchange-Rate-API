package pkg

import (
	// "fmt"
	"fmt"
	"log"
	"net/http"

	"exchange-rate-api/db"
	"exchange-rate-api/tools"

	"github.com/PuerkitoBio/goquery"
)

func Scrapping() ([]db.Currencies, []db.ExchangeRate, error) {
	// scrape the html
	resp, err := http.Get("https://www.bi.go.id/id/statistik/informasi-kurs/transaksi-bi/default.aspx")
	if err != nil {
		log.Panicln("error (pkg/scraper)[1] - http.Get failed, msg:", err)
		return nil, nil, err
	}
	defer resp.Body.Close()

	// Check the status code response
	if resp.StatusCode != 200 {
		log.Panicln("error (pkg/scraper)[2] - status code isn't OK, status code:", err)
		return nil, nil, fmt.Errorf("status code is wrong, code: %d\n", resp.StatusCode)
	}

	// load response body
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Panicln("error (pkg/scraper)[3] - failed to load response body, msg:", err)
		return nil, nil, err
	}

	// Find currencies name and code data from website inside table
	currencies := make([]db.Currencies, 26)
	id := uint64(1)
	currencies[0].ID = id
	currencies[0].CurrencyCode = "IDR"
	currencies[0].CurrencyName = "INDONESIAN RUPIAH"

	doc.Find(".table-md").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			// fmt.Println("--- indextr:", indextr)
			if indextr > 0 {
				id++
				currencies[indextr].ID = id
				rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
					// fmt.Println("--- indextr:", indextr)
					switch {
					case indexth == 0:
						currencies[indextr].CurrencyCode = deleteSpace(tablecell.Text())
					case indexth == 1:
						currencies[indextr].CurrencyName = deleteSpace(tablecell.Text())
					}
				})
			}
		})
	})
	err = WriteJsonFile("./assets/currencies.json", currencies)
	if err != nil {
		log.Println("error (pkg/scraper)[4] - failed to write currencies scrapping data into json file, msg:", err)
		return nil, nil, err
	}

	// Find exchange rate from website inside table
	exchangeRates := make([]db.ExchangeRate, 25)

	id = uint64(1)
	doc.Find(".table-lg").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			// fmt.Println("--- indextr:", indextr)
			if indextr > 0 {
				exchangeRates[indextr-1].ID = id
				id++
				exchangeRates[indextr-1].CurrencyCodeFrom = "IDR"
				rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
					// fmt.Println("--- indextr:", indextr)
					switch {
					case indexth == 0:
						exchangeRates[indextr-1].CurrencyCodeTo = deleteSpace(tablecell.Text())
					case indexth == 2:
						exchangeRates[indextr-1].Sell = tools.ConvertToFloat(tablecell.Text())
					case indexth == 3:
						exchangeRates[indextr-1].Buy = tools.ConvertToFloat(tablecell.Text())
					}
				})
			}
		})
	})
	err = WriteJsonFile("./assets/exchange_rates.json", exchangeRates)
	if err != nil {
		log.Println("error (pkg/scraper)[5] - failed to write exchange_rates scraping data into json file, msg:", err)
		return nil, nil, err
	}
	return currencies, exchangeRates, nil
}

// Delete blank space after the of the real string data that scraped from website
func deleteSpace(s string) string {
	for i := range s {
		if s[i] == ' ' && s[i-1] == ' ' {
			s = s[:i-1]
			if s[len(s)-1] == '\n' {
				s = s[:len(s)-1]
			}
			return s
		}
	}
	return s
}

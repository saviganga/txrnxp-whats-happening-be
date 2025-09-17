package external

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GoqueryScraper implements the Scraper interface using goquery.
type GoqueryScraper struct{}

// // ScrapeEvents scrapes events from the provided URL.
// func (s *GoqueryScraper) ScrapeEvents(url string) ([]Event, error) {
//     res, err := http.Get(url)
//     if err != nil {
//         return nil, err
//     }
//     defer res.Body.Close()

//     doc, err := goquery.NewDocumentFromReader(res.Body)
//     if err != nil {
//         return nil, err
//     }

//     var events []Event
//     doc.Find(".eds-event-card-content__primary-content").Each(func(i int, s *goquery.Selection) {
//         title := s.Find(".eds-is-hidden-accessible").Text()
//         date := s.Find(".eds-text-bs--fixed").Text()
//         location := s.Find(".card-text--truncated").Text()
//         description := s.Find(".eds-event-card-content__sub-content").Text()

//         events = append(events, Event{
//             Title:       strings.TrimSpace(title),
//             Date:        strings.TrimSpace(date),
//             Location:    strings.TrimSpace(location),
//             Description: strings.TrimSpace(description),
//         })
//     })

//     return events, nil
// }

// ScrapeEvents scrapes events from the provided URL with debug prints.
func (s *GoqueryScraper) ScrapeEvents(url string) ([]Event, error) {
	fmt.Println("Starting ScrapeEvents for URL:", url)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return nil, err
	}
	defer res.Body.Close()
	fmt.Println("Fetched URL successfully, status code:", res.StatusCode)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Error parsing HTML document:", err)
		return nil, err
	}
	fmt.Println("Parsed HTML document successfully")

	var events []Event
	doc.Find(".eds-event-card-content__primary-content").Each(func(i int, s *goquery.Selection) {
		fmt.Println("Processing event index:", i)

		title := s.Find(".eds-is-hidden-accessible").Text()
		date := s.Find(".eds-text-bs--fixed").Text()
		location := s.Find(".card-text--truncated").Text()
		description := s.Find(".eds-event-card-content__sub-content").Text()

		fmt.Println("Raw values extracted:")
		fmt.Println("Title:", title)
		fmt.Println("Date:", date)
		fmt.Println("Location:", location)
		fmt.Println("Description:", description)

		event := Event{
			Title:       strings.TrimSpace(title),
			Date:        strings.TrimSpace(date),
			Location:    strings.TrimSpace(location),
			Description: strings.TrimSpace(description),
		}

		fmt.Println("Trimmed and appended event:", event)
		events = append(events, event)
	})

	fmt.Println("Scraping completed, total events found:", len(events))
	return events, nil
}

package external

// Event represents the structure of an event.
type Event struct {
	Title       string
	Date        string
	Location    string
	Description string
}

// Scraper defines the methods for scraping event data.
type Scraper interface {
	ScrapeEvents(url string) ([]Event, error)
}

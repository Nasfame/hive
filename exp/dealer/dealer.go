package dealer

// Dealer defines the interface for the plugin.
type Dealer interface {
	DealsMatched(dealID string)
	DealsAgreed() <-chan string
}

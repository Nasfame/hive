// Package dealer
//
// The `New` function type is intended to be implemented by your dealer plugin to integrate with the Hive application.
// When creating your plugin, ensure that the constructor function matches the signature of the `New` type. Additionally,
// export the above symbols so that they can be loaded dynamically by the Hive application at runtime.
//
// Example usage of an example plugin:
//
//	package main
//
//	import "context"
//	import "hive/dealer"
//
//	// MyDealer implements the Dealer interface.
//	type MyDealer struct { ... }
//
//	// DealMatched is the method that processes matched deals.
//	func (d *MyDealer) DealMatched(dealID string) { ... }
//
//	// DealsAgreed is the method that provides a channel of agreed deals.
//	func (d *MyDealer) DealsAgreed() <-chan string { ... }
//
//	// New creates a new instance of MyDealer.
//	func New(ctx context.Context) Dealer {
//	    return &MyDealer{ ... }
//	}
//
// For more examples refer:  https://github.com/CoopHive/hive/blob/main/exp/dealer
package dealer

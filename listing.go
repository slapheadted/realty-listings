package main

// Listing is the common interface for a property listing
type Listing interface {
	getId() int
	getName() string
	getPrice() float32
}
package main

import (
  "fmt"
  "encoding/json"
  "github.com/parnurzeal/gorequest"
  "log"
)

type listingItem struct {
  ListingID int    `json:"Listing_Id"`
  Address   string `json:"Address"`
}

type listingsList struct {
  ListingsSummaryVM []listingItem `json:"ListingsSummaryVM"`
}

type listingResponse struct {
  ListingsList listingsList `json:"ListingsList"`
}

func (l listingItem) getName() string {
  return l.Address
}

func fetchListings() []listingItem {
  request := gorequest.New()
  _, body, errs := request.Post("https://www.point2homes.com/CA/Real-Estate-Listings.html").
    Send("&location=Kootenay+Rockies,+BC&search_mode=location&PriceMin=150000&PriceMax=450000&LotSizeMin=5&page=1&SelectedView=listings&LocationGeoId=432040&ajax=1").
    End()

  if len(errs) > 0 {
    fmt.Println(errs)
  }

  var cont listingResponse

  if err := json.Unmarshal([]byte(body), &cont); err != nil {
    log.Fatal(err)
  }

  var singleListing = cont.ListingsList.ListingsSummaryVM[0]
  fmt.Println("single", singleListing.getName())

  return cont.ListingsList.ListingsSummaryVM
}

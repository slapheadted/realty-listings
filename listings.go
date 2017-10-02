package main

import (
	"github.com/parnurzeal/gorequest"
)

// type listing struct {
//   id int
//   name string
// }

func fetchListings() string {
  request := gorequest.New()
  _, body, errs := request.Get("http://www.mocky.io/v2/59d18c831200001307244ed5").End()

  if len(errs) > 0 {
    return "error"
  }

  return body
}

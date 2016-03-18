package main

import (
  "./tsutaya"
  "fmt"
)

func main() {
  url := "http://store-tsutaya.tsite.jp/top/rels/dvd_rental.html"
  pages := tsutaya.GetDVDPages(url)

  itemsUrls := tsutaya.GoGetDVDItemUrls(pages)
  results := tsutaya.GoGetDVDItems(itemsUrls)

  for _, result := range results {
    fmt.Println(result.Title + " : " + result.ReleasedAt)
  }
}

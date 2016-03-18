package tsutaya

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "sync"
)

func GetDVDItem(itemUrl string) ResultItem {
  doc, err := goquery.NewDocument(itemUrl)
  if err != nil {
    fmt.Println("Failed itemUrl scraping")
  }
  title := doc.Find(".header h2 span").First().Text()
  releasedAt := doc.Find(".detailBox li").First().Next().Text()
  result := ResultItem{title, releasedAt}
  return result
}

func GoGetDVDItems(urls []string) []ResultItem {
  results := []ResultItem{}
  var wg sync.WaitGroup
  for _, url := range urls {
    wg.Add(1)
    go func(url string) {
      defer wg.Done()
      results = append(results, GetDVDItem(url))
    }(url)
  }
  wg.Wait()
  return results
}

func GetDVDItemUrls(url string) []string {
  urls := []string{}
  doc, err := goquery.NewDocument(url)
  if err != nil {
    fmt.Println("Failed url scraping")
  }
  doc.Find(".itemGroup .imageBlock a").Each(func(_ int, s *goquery.Selection) {
    url, _ := s.Attr("href")
    urls = append(urls, url)
  })
  return urls
}

func GoGetDVDItemUrls(pages []string) []string {
  itemUrls := []string{}
  // sync.waitGroupはWait()を呼び出すとAdd()を呼び出した回数からDone()を呼び出した回数を引いて0になるまで待機する
  // 全てのgoroutineの修了を待つ場合に使用する
  var wg sync.WaitGroup
  for _, page := range pages {
    wg.Add(1)
    // 並行処理を走らせる
    go func(page string) {
      defer wg.Done()
      itemUrls = append(itemUrls, GetDVDItemUrls(page)...)
    }(page)
  }
  wg.Wait()
  return itemUrls
}

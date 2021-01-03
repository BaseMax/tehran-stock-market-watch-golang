package main

import (
  "log"
  "fmt"
  // "io/ioutil"
  "strings"
  "net/http"
  "io/ioutil"
  "crypto/tls"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func get(link string, refer string) (string, bool) {
  if refer == "" {
    refer = link
  }
  tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }
  client := &http.Client{Transport: tr}

  req, err := http.NewRequest("GET", link, nil)
  if err != nil {
    log.Fatalln(err)
    return "", false
  }
  req.Header.Set("Connection", "keep-alive")
  req.Header.Set("Accept", "text/html, */*; q=0.01")
  req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
  req.Header.Set("X-Requested-With", "XMLHttpRequest")
  req.Header.Set("Referer", refer)
  req.Header.Set("Accept-Language", "en,fa;q=0.9")
  req.Header.Set("Cookie", "_ga=GA1.2.1444158647.1604418871; ASP.NET_SessionId=dulsflnucit53l0cunbuld11; _gid=GA1.2.1069471190.1609606129; ASP.NET_SessionId=pu2sgzryj5er4jvrndv1nlws")

  resp, err := client.Do(req)
  if err != nil {
    log.Fatalln(err)
    return "", false
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
    return "", false
  }
  sb := string(body)
  // fmt.Println(sb)
  return sb, true
}

func main() {
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/stock")
  defer db.Close()
  if err != nil {
    log.Fatal(err)
  }

  var version string
  err2 := db.QueryRow("SELECT VERSION()").Scan(&version)
  if err2 != nil {
    log.Fatal(err2)
  }

  fmt.Println(version)

  res, err3 := get(
    "http://www.tsetmc.com/tsev2/data/MarketWatchInit.aspx?h=0&r=0",
    "http://www.tsetmc.com/Loader.aspx?ParTree=15131F")
  fmt.Println(res)
  fmt.Println(err3)

  items := strings.Split(res, "@")
  // fmt.Println(items)
  rows := strings.Split(items[2], ";")
  // fmt.Println(rows)
  for _, row := range rows {
    cols := strings.Split(row, ",")
    fmt.Println(row)
    fmt.Println(cols)
  }
}


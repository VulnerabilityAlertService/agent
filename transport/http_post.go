package transport

import (
  "fmt"
  "log"
  "os"
  "net/http"
  "strings"
  "encoding/json"
  "bytes"
  "crypto/tls"
)

const HOST = "https://vas.lepra.jp"

type Postdata struct {
  Token    string `json:"token"`
  Alerts []Item `json:"alerts"`
}

type Item struct {
  Package string `json:"package"`
  Version string `json:"version"`  
}

func Dataset(data []map[string]string, key string) Postdata {
  values := Postdata{}
  for _, record := range data {
    item := Item{}
    item.Package = record["package"]
    item.Version = record["version"]
    values.Alerts = append(values.Alerts, item)
  }

  values.Token = key
  return values
}

func Post(data Postdata) {
  http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

  var buf bytes.Buffer
  b, _ := json.Marshal(data)
  buf.Write(b)

  req, err := http.NewRequest("POST", fmt.Sprint(HOST, "/api/check_pkg"), strings.NewReader(buf.String()))
  if err != nil {
    log.Fatalln("err : ", err)
  }

  // Content-Type 設定
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    log.Fatalln("err : ", err)
  }
  defer resp.Body.Close()

  switch resp.StatusCode {
  case 201:
    return
  case 401:
    log.Print("Unauthorized...")
    os.Exit(1)
  default:
    log.Print("Server error!!")
    os.Exit(1)
  }
}

package main

import (
  "os"
  "log"
  "net/http"

  "github.com/bs-online-judge/api/api"
  "github.com/bs-online-judge/api/models"
)

func main() {
  err := models.CreateTables()
  if err != nil {
    log.Fatal(err)
  }
  log.Fatal(http.ListenAndServe(":" + os.Getenv("API_PORT"), api.Router()))
}

package main

import (
	"github.com/labstack/echo"
	"github.com/youngvform/go-turorial/job-scraper/scraper"
	"os"
	"strings"
)

const homeFileName = "home.html"
const csvFileName = "jobs.csv"

func handleHome (c echo.Context) error {
	return c.File(homeFileName)
}
func handleScrape(c echo.Context) error {
	defer os.Remove(csvFileName)
	query := strings.ToLower(scraper.RemoveSpace(c.FormValue("query")))
	scraper.Scrape(query)
	return c.Attachment(csvFileName, csvFileName)
}

func main()  {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
	
}
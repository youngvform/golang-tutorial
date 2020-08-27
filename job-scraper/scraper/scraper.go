package scraper

import (
"encoding/csv"
"fmt"
"github.com/PuerkitoBio/goquery"
"log"
"net/http"
"os"
"strconv"
"strings"
"time"
)

type extractedJob struct {
	id string
	title string
	location string
	salary string
	summary string
}


func Scrape(searchText string)  {
	var baseUrl = "https://kr.indeed.com/jobs?q=" + searchText + "&limit=50"
	startTime := time.Now()
	var jobs []extractedJob
	requestChannel := make(chan []extractedJob)
	done := make(chan bool)
	pageNumber := getPageNumber(baseUrl)
	for i := 0; i < pageNumber; i++ {
		go getPage(baseUrl, i, requestChannel)
		//extractedJobs := <- c
		//jobs = append(jobs, extractedJobs...)
	}
	for i := 0; i < pageNumber; i++ {
		extractedJobs := <- requestChannel
		jobs = append(jobs, extractedJobs...)
	}
	//writeJobs(jobs, done)
	//fmt.Println(len(jobs), " extracted Done!")
	go writeJobs(jobs, done)
	if <-done == true {
		fmt.Println(len(jobs), " extracted Done!")
	} else {
		fmt.Println("writing failed.")
	}

	endTime := time.Now()
	fmt.Println("Operating time: ", endTime.Sub(startTime))
}

func getPage(baseUrl string, page int, mainChannel chan <- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageUrl := baseUrl + "&start=" + strconv.Itoa(page * 50)
	doc := requestUrl(pageUrl)

	cards := doc.Find(".jobsearch-SerpJobCard")
	cards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
		//job := <- c
		//jobs = append(jobs, job)
	})

	for i := 0; i < cards.Length(); i++ {
		job := <- c
		jobs = append(jobs, job)
	}

	mainChannel <- jobs
}

func writeJobs(jobs []extractedJob, mainChannel chan <- bool)  {
	var names [][]string
	c := make(chan []string)
	file, err := os.Create("jobs.csv")
	checkError(err)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}

	checkError(writer.Write(headers))

	for _, job := range jobs{
		go writeFile(job, c)
	}

	for i := 0; i < len(jobs); i++ {
		name := <-c
		names = append(names, name)
	}
	writer.WriteAll(names)

	mainChannel <- true
}

func writeFile(job extractedJob, c chan <- []string)  {
	slicedJob := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
	c <- slicedJob
}

func extractJob(card *goquery.Selection, c chan <- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := RemoveSpace(card.Find(".title > a").Text())
	location := RemoveSpace(card.Find(".accessible-contrast-color-location").Text())
	salary := RemoveSpace(card.Find(".salaryText").Text())
	summary := RemoveSpace(card.Find(".summary").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func RemoveSpace(str string) string  {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPageNumber(baseUrl string) int {
	pages := 0
	doc := requestUrl(baseUrl)
	pages = doc.Find(".pagination").Find("a").Length()
	return pages
}

func requestUrl(url string) *goquery.Document {
	fmt.Println("Request to " + url);
	res, err := http.Get(url)
	checkError(err)
	checkStatusCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)
	return doc
}

func checkError(err error)  {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatusCode(res *http.Response)  {
	if res.StatusCode != 200 {
		log.Fatalln("Failed status code", res.StatusCode)
	}
}

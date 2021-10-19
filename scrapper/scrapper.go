package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id string
	title string
	location string
	salary string
	summary string
}

// Scrape Indeed by a term
func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)

	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJob := <-c
		jobs = append(jobs, extractedJob...)
	}

	writeJobs(jobs, baseURL)
	fmt.Println("Done, extracted", len(jobs))
}

func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := url + "&start=" + strconv.Itoa(page * 50)
	fmt.Println("Requesting", pageURL);
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".sponTapItem")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})
	
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".jobTitle>span").Text())
	location := CleanString(card.Find(".companyLocation").Text())
	summary := CleanString(card.Find(".job-snippet").Text())
	c <- extractedJob{
		id: id, 
		title: title, 
		location: location, 
		summary: summary}
}

// CleanString cleans a string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	// 메모리 새어나가는 것을 막을 수 있다.
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		//title := s.Find("a").Text()
		//fmt.Printf("Review %d: %s\n", i, title)
		pages = s.Find("a").Length()
	})
	
	return pages
}

// csv 파일 생성
func writeJobs(jobs []extractedJob, url string) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	// 정보들이 파일에 다 Write 되고 실행된다.
	defer w.Flush()

	headers := []string{"ID", "Title", "Location", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{url + "&vjk=" + job.id, job.title, job.location, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}


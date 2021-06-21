package hotelaah

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

type StringPair struct {
	First  string
	Second string
}

func (sp *StringPair) String() string {
	return "[" + sp.First + "," + sp.Second + "]"
}

// implement AahData interface
func (sp *StringPair) Value() []byte {
	return []byte(sp.String())
}

// GET the content from a given url
type Fetcher struct {
	url string
	doc *goquery.Document
}

func NewFetcher(s string) *Fetcher {
	return &Fetcher{
		url: s,
	}
}

// filter \t \n without package strings
// func specialCharFilter(s string) []string {
// 	result := []string
// 	for
//
// }

// the encoding of the given url is gb2312
func (f *Fetcher) GetAndExtract() ([]StringPair, error) {
	resp, err := http.Get(f.url)
	if err != nil {
		log.Printf("[http] GET of %s failed: %v", f.url, err)
		return []StringPair{}, err
	}
	defer resp.Body.Close()
	utfBody, err := iconv.NewReader(resp.Body, "gb2312", "utf-8")
	if err != nil {
		return []StringPair{}, err
	}
	f.doc, err = goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		log.Printf("[goquery] %s's resp failed: %v", f.url, err)
		return []StringPair{}, err
	}
	var sp []StringPair
	// selection := "table[border]='1' tbody"
	selection := "table[border='1'] tbody tr"
	f.doc.Find(selection).Each(func(i int, s *goquery.Selection) {
		var p StringPair
		s.Find("td").Each(func(j int, ss *goquery.Selection) {
			switch j {
			case 1:
				p.First = strings.Trim(ss.Text(), "\n \t")
			case 2:
				p.Second = strings.Trim(ss.Text(), "\n \t")
			}
		})
		sp = append(sp, p)
		// tupleTd := strings.Split(strings.Trim(content, "\t\n"), " ")
		log.Printf("[Fetcher] %d - %s\n", i, p.String())
	})
	return sp, nil
}

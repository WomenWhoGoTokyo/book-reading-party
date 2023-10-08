package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Entry struct {
	AuthorID string
	Author   string
	TitleID  string
	Title    string
	SiteURL  string
	ZipURL   string
}

func findEntries(siteURL string) ([]Entry, error) {
	// Deprecated:
	// doc, err := goquery.NewDocument(siteURL)
	res, err := http.Get(siteURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	// リスト9.3
	// doc.Find("ol li a").Each(func(n int, elem *goquery.Selection) {
	// 	fmt.Println(elem.Text(), elem.AttrOr("href", ""))
	// })

	// リスト9.4
	// pat := regexp.MustCompile(`.*/cards/([0-9]+)/card([0-9]+).html$`)
	// doc.Find("ol li a").Each(func(n int, elem *goquery.Selection) {
	// 	token := pat.FindStringSubmatch(elem.AttrOr("href", ""))
	//	if len(token) != 3 {
	// 		return
	// 	}
	// 	pageURL := fmt.Sprintf("https://www.aozora.gr.jp/cards/%s/card%s.html", token[1], token[2])
	// 	fmt.Println(pageURL)
	// })

	// リスト9.5 - 9.9
	entries := []Entry{}
	pat := regexp.MustCompile(`.*/cards/([0-9]+)/card([0-9]+).html$`)
	doc.Find("ol li a").Each(func(n int, elem *goquery.Selection) {
		token := pat.FindStringSubmatch(elem.AttrOr("href", ""))
		if len(token) != 3 {
			return
		}
		title := elem.Text()
		pageURL := fmt.Sprintf("https://www.aozora.gr.jp/cards/%s/card%s.html", token[1], token[2])
		author, zipURL := findAuthorAndZIP(pageURL)
		if zipURL != "" {
			entries = append(entries, Entry{
				AuthorID: token[1],
				Author:   author,
				TitleID:  token[2],
				Title:    title,
				SiteURL:  siteURL,
				ZipURL:   zipURL,
			})
		}
	})

	return entries, nil
}

// リスト9.6 - 9.7
func findAuthorAndZIP(siteURL string) (string, string) {
	// log.Println("query", siteURL)

	// Deprecated:
	// doc, err := goquery.NewDocument(siteURL)
	res, err := http.Get(siteURL)
	if err != nil {
		return "", ""
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", ""
	}

	// 正誤表 P210 により以下のように修正
	// https://gihyo.jp/book/2023/978-4-297-13419-8/support
	author := doc.Find("table[summary=作者データ] tr:nth-child(2) td:nth-child(2)").First().Text()

	zipURL := ""
	doc.Find("table.download a").Each(func(n int, elem *goquery.Selection) {
		href := elem.AttrOr("href", "")
		if strings.HasSuffix(href, ".zip") {
			zipURL = href
		}
	})

	if zipURL == "" {
		return author, zipURL
	}

	if strings.HasPrefix(zipURL, "http://") || strings.HasPrefix(zipURL, "https://") {
		return author, zipURL
	}

	u, err := url.Parse(siteURL)
	if err != nil {
		return author, ""
	}

	u.Path = path.Join(path.Dir(u.Path), zipURL)
	return author, u.String()
}

func main() {
	listURL := "https://www.aozora.gr.jp/index_pages/person879.html"

	entries, err := findEntries(listURL)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		content, err := extractText(entry.ZipURL)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(entry.SiteURL)
		fmt.Println(content)
	}
}

func extractText(zipURL string) (string, error) {
	resp, err := http.Get(zipURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return "", err
	}

	for _, file := range r.File {
		if path.Ext(file.Name) == ".txt" {
			f, err := file.Open()
			if err != nil {
				return "", err
			}
			b, err := ioutil.ReadAll(f)
			f.Close()
			if err != nil {
				return "", err
			}
			return string(b), nil
		}
	}
	return "", nil
}

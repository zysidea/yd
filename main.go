package main

import (
	"os"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"unicode"
)

var URL = "http://dict.youdao.com/w/"
var isChinese bool=false

func main() {
	args := os.Args
	if len(args) != 2 {
		usage()
		return
	}
	word := args[1]
	if isChineseChar(word){
		isChinese=true
	}
	url := URL + word
	getResult(url)

}

func getResult(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	if isChinese{
		doc.Find("#phrsListTab .trans-container ul p .contentTitle").Each(func(i int, selection *goquery.Selection)  {
			text := selection.Find("a").Text()
			fmt.Println(text)
		})

	}else {
		doc.Find("#phrsListTab .trans-container ul li").Each(func(i int, selection *goquery.Selection)  {
			text := selection.Text()
			fmt.Println(text)
		})
	}
}

func isChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func usage() {
	 fmt.Printf("%s\n", "Usage: yd xxx")
}

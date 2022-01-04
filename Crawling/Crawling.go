package crawling

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func Basic1() {
	// get page
	doc, docErr := goquery.NewDocument("https://modu-print.tistory.com/449?category=863372")
	if docErr != nil {
		fmt.Println(docErr.Error())
		return
	}
	// get id
	codeArea := doc.Find("#code_1606220695117")
	// if it is
	if codeArea != nil {
		fmt.Println(codeArea.Text())
		fmt.Println(codeArea.Html())
	}
}

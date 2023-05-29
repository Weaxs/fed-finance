package fed_finance

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
)

func requestHtml(url string) (doc *goquery.Document, err error) {
	response, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("request failed: %w", err)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	htmlBytes, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("read response body error: %w", err)
		return
	}

	doc, err = goquery.NewDocumentFromReader(strings.NewReader(string(htmlBytes)))
	if err != nil {
		err = fmt.Errorf("parse body to html error: %w", err)
		return
	}
	return
}

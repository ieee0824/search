package main

import (
	"flag"
	"fmt"
	"log"
	u "net/url"
	"strings"

	"github.com/pkg/browser"
)

const baseURL = "https://www.google.com/search"

func main() {
	lang := flag.String("l", "", "")
	flag.Parse()

	if *lang == "" {
		*lang = "ja"
	}
	q := u.Values{}
	q.Add("hl", *lang)
	q.Add("lr", fmt.Sprintf("lang_%s", *lang))
	q.Add("q", strings.Join(flag.Args(), " "))

	if len(flag.Args()) == 0 {
		log.Fatalln("empty search word error")
	}

	url, err := u.Parse(baseURL)
	if err != nil {
		log.Fatalln(err)
	}

	url.RawQuery = q.Encode()

	if err := browser.OpenURL(url.String()); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	u "net/url"
	"os/exec"
	"runtime"
	"strings"
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

	switch runtime.GOOS {
	case "darwin":
		if err := exec.Command("open", url.String()).Run(); err != nil {
			log.Fatalln(err)
		}
	default:
		log.Fatalln("unknown os")
	}
}

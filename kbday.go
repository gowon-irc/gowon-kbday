package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
)

func colourList(in []string) (out []string) {
	colours := []string{"green", "red", "blue", "orange", "magenta", "cyan", "yellow"}
	cl := len(colours)

	for n, i := range in {
		c := colours[n%cl]
		o := fmt.Sprintf("{%s}%s{clear}", c, i)
		out = append(out, o)
	}

	return out
}

func kbday() (string, error) {
	t := time.Now()
	tf := fmt.Sprintf("%s_%d", t.Format("January"), t.Day())

	log.Println(fmt.Sprintf("Fetching birthday info for %s", tf))

	url := fmt.Sprintf("https://kpop.fandom.com/wiki/Template:%s", tf)

	resp, err := soup.Get(url)
	if err != nil {
		return "", err
	}

	doc := soup.HTMLParse(resp)

	bdays := []string{}
	for _, li := range doc.Find("table").FindAll("li") {
		bdays = append(bdays, li.FullText())
	}

	cl := colourList(bdays)

	return fmt.Sprintf("Happy birthday to: %s", strings.Join(cl, ", ")), nil
}

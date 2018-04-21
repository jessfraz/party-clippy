package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

const (
	clippy = `
 _________________________________
/ It looks like you're building a \
\ microservice.                   /
 ---------------------------------
 \
  \
     __
    /  \
    |  |
    @  @
    |  |
    || |/
    || ||
    |\_/|
    \___/
`
)

var (
	colorOptions = []*color.Color{
		color.New(color.FgHiRed),
		color.New(color.FgHiGreen),
		color.New(color.FgHiYellow),
		color.New(color.FgHiBlue),
		color.New(color.FgHiMagenta),
		color.New(color.FgHiCyan),
	}
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		i := 0
		for i < len(colorOptions) {
			// Clear the terminal.
			fmt.Fprintf(w, "\033[2J")
			// Print clippy with a color.
			colorOptions[i].Fprintf(w, clippy)
			if i == len(colorOptions)-1 {
				i = 0
				continue
			}
			i++
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}

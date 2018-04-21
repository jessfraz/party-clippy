package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

const (
	clippy = `
 __
/  \        _____________
|  |       /             \
@  @       | Hello       |
|| ||      | //Build     |
|| ||   <--| Containers  |
|\_/|      | rule!       |
\___/      \_____________/
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
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		quote := strings.Trim(req.URL.Path, "/")
		if quote == "" {
			quote = "Hello //Build!"
		}

		userAgent := req.Header.Get("user-agent")
		if strings.Contains(userAgent, "curl") {
			// we have a request from curl
		}

		i := 0
		for i < len(colorOptions) {
			fmt.Fprintf(w, "\033[H")
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

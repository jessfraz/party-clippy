package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/jessfraz/party-clippy/version"
	"github.com/sirupsen/logrus"
)

const (
	// BANNER is what is printed for help/info output.
	BANNER = `party-clippy

  version: %s
`
)

var (
	port int

	debug bool
	vrsn  bool

	colorOptions = []*color.Color{
		color.New(color.FgRed),
		color.New(color.FgGreen),
		color.New(color.FgYellow),
		color.New(color.FgBlue),
		color.New(color.FgMagenta),
		color.New(color.FgCyan),
	}
)

func init() {
	// parse flags
	flag.IntVar(&port, "port", 8080, "port for the server to listen on")
	flag.IntVar(&port, "p", 8080, "port for the server to listen on (shorthand)")

	flag.BoolVar(&vrsn, "version", false, "print version and exit")
	flag.BoolVar(&vrsn, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, version.VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if vrsn {
		fmt.Printf("party-clippy version %s, build %s", version.VERSION, version.GITCOMMIT)
		os.Exit(0)
	}

	if flag.NArg() >= 1 {
		// parse the arg
		arg := flag.Args()[0]

		if arg == "help" {
			usageAndExit("", 0)
		}

		if arg == "version" {
			fmt.Printf("party-clippy version %s, build %s", version.VERSION, version.GITCOMMIT)
			os.Exit(0)
		}
	}

	// set log level
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func main() {
	// On ^C, or SIGTERM handle exit.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for sig := range c {
			logrus.Infof("Received %s, exiting.", sig.String())
			os.Exit(0)
		}
	}()

	mux := http.NewServeMux()

	// Define bburns easter egg handler.
	mux.HandleFunc("/bburns", func(w http.ResponseWriter, req *http.Request) {
		i := 0
		for i < len(colorOptions) {
			// Clear the terminal.
			fmt.Fprint(w, "\033c")
			// Print clippy with a color.
			colorOptions[i].Fprint(w, bburns)
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			time.Sleep(time.Second / 6)
			if i == len(colorOptions)-1 {
				i = 0
				continue
			}
			i++
		}
	})

	// Define octocat easter egg handler.
	mux.HandleFunc("/octocat", func(w http.ResponseWriter, req *http.Request) {
		i := 0
		for i < len(colorOptions) {
			// Clear the terminal.
			fmt.Fprint(w, "\033c")
			// Print clippy with a color.
			colorOptions[i].Fprint(w, octocat)
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			time.Sleep(time.Second / 6)
			if i == len(colorOptions)-1 {
				i = 0
				continue
			}
			i++
		}
	})

	// Define party easter egg handler.
	mux.HandleFunc("/party", partyHandler)

	// Define wildcard/root handler.
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if !strings.HasPrefix(req.UserAgent(), "curl") {
			fmt.Fprint(w, clippyBlurb+clippyBody0)
			return
		}

		// Do the party handler if it's from curl.
		partyHandler(w, req)
	})

	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

func partyHandler(w http.ResponseWriter, req *http.Request) {
	i := 0
	for i < len(colorOptions) {
		// Clear the terminal.
		fmt.Fprint(w, "\033c")
		// Print clippy with a color.
		colorOptions[i].Fprint(w, clippyBlurb+wiggle(i))
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		time.Sleep(time.Second / 6)
		if i == len(colorOptions)-1 {
			i = 0
			continue
		}
		i++
	}
}

func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}

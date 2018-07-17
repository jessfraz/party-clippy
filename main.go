package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/genuinetools/pkg/cli"
	"github.com/jessfraz/party-clippy/version"
	"github.com/sirupsen/logrus"
)

var (
	port int

	debug bool

	colorOptions = []*color.Color{
		color.New(color.FgRed),
		color.New(color.FgGreen),
		color.New(color.FgYellow),
		color.New(color.FgBlue),
		color.New(color.FgMagenta),
		color.New(color.FgCyan),
	}
)

func main() {
	// Create a new cli program.
	p := cli.NewProgram()
	p.Name = "party-clippy"
	p.Description = "It looks like you are trying to make a webserver"

	// Set the GitCommit and Version.
	p.GitCommit = version.GITCOMMIT
	p.Version = version.VERSION

	// Setup the global flags.
	p.FlagSet = flag.NewFlagSet("global", flag.ExitOnError)
	p.FlagSet.IntVar(&port, "port", 8080, "port for the server to listen on")
	p.FlagSet.IntVar(&port, "p", 8080, "port for the server to listen on (shorthand)")

	p.FlagSet.BoolVar(&debug, "d", false, "enable debug logging")
	p.FlagSet.BoolVar(&debug, "debug", false, "enable debug logging")

	// Set the before function.
	p.Before = func(ctx context.Context) error {
		// Set the log level.
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		return nil
	}

	// Set the main program action.
	p.Action = func(ctx context.Context, args []string) error {
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

		// Define party easter egg handler.
		mux.HandleFunc("/party", func(w http.ResponseWriter, req *http.Request) {
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
		})

		// Define wildcard/root handler.
		mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprint(w, clippyBlurb+clippyBody0)
		})

		logrus.Infof("Starting server on port %d", port)
		logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))

		return nil
	}

	// Run our program.
	p.Run()
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

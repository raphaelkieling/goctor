package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dimiro1/banner"
	"github.com/gookit/color"
)

var flagFile = flag.String("f", "", "Configuration file")

func main() {
	flag.Parse()

	var conf Configuration

	result, err := conf.readFile(*flagFile)

	if err != nil {
		panic(err)
	}

	anyError := false

	templ := fmt.Sprintf("{{.Title \"%v\" \"\" 1}}\n", result.Banner.Title)
	banner.InitString(os.Stdout, true, true, templ)

	for _, exam := range result.Exams {
		exitCode, goctorcodes, err := exam.execute()

		if err != nil {
			anyError = true
			color.Yellow.Println("🔴 " + exam.Name + " - " + exam.Description)
			for _, possibility := range exam.Possibilities {
				if possibility.matchError(exitCode, goctorcodes) {
					color.Red.Println("\t " + possibility.Message)
				}
			}
		} else {
			color.Green.Println("🟢 " + exam.Name + " - " + exam.Description)
		}
	}

	if anyError {
		os.Exit(1)
	}

	os.Exit(0)
}

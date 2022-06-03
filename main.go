package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/dimiro1/banner"
	"github.com/gookit/color"
)

var flagFile = flag.String("f", "", "Configuration file")

func init() {
	flag.Parse()
}

func main() {
	var conf Configuration

	result, err := conf.readFile(*flagFile)

	if err != nil {
		panic(err)
	}

	anyError := false

	templ := fmt.Sprintf("{{.Title \"%v\" \"\" 1}}\n", result.Banner.Title)
	banner.InitString(os.Stdout, true, true, templ)

	for _, exam := range result.Exams {
		outBytes, err := exec.Command("/bin/bash", "-c", exam.Run).Output()
		if err != nil {
			panic(err)
		}
		out := string(outBytes)

		r, _ := regexp.Compile("(goctor::)(.+)")
		foundCodes := r.FindAllString(out, 99)
		hasError := len(foundCodes) > 0

		if hasError {
			anyError = true
			color.Yellow.Println("🔴 " + exam.Name + " - " + exam.Description)
			for _, possibility := range exam.Possibilites {
				for _, code := range foundCodes {
					if code == "goctor::"+possibility.Code {
						color.Red.Println("\t " + possibility.Message)
					}
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

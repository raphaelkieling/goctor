package main

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"regexp"

	"gopkg.in/yaml.v2"
)

type Possiblity struct {
	Goctorcode string `yaml:"goctorcode"`
	Code       int    `yaml:"code"`
	Message    string `yaml:"message"`
}

func (p *Possiblity) matchError(exitCode int, goctorCodes []string) bool {
	if exitCode == p.Code {
		return true
	}

	for _, code := range goctorCodes {
		if code == "goctor::"+p.Goctorcode {
			return true
		}
	}

	return false
}

type Step struct {
	Name          string       `yaml:"name"`
	Description   string       `yaml:"description"`
	Run           string       `yaml:"run"`
	Possibilities []Possiblity `yaml:"possibilities"`
}

func (s *Step) execute() (int, []string, error) {
	cmd := exec.Command("/bin/bash", "-c", s.Run)

	outBytes, err := cmd.Output()
	exitStatus := 0

	if err != nil {
		exitStatus = cmd.ProcessState.ExitCode()
	}

	out := string(outBytes)

	r, err := regexp.Compile("(goctor::)(.+)")

	if err != nil {
		return 0, nil, err
	}

	foundCodes := r.FindAllString(out, 99)

	if len(foundCodes) > 0 || exitStatus != 0 {
		return exitStatus, foundCodes, errors.New("code was found")
	}

	return 0, nil, nil
}

func (s *Step) getName() string {
	if len(s.Description) <= 0 {
		return s.Name
	}
	return s.Name + " - " + s.Description
}

type Exam struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Steps       []Step `yaml:"steps"`
}

func (e *Exam) getName() string {
	if len(e.Description) <= 0 {
		return e.Name
	}
	return e.Name + " - " + e.Description
}

type Configuration struct {
	Version string `yaml:"version"`
	Banner  struct {
		Title string `yaml:"title"`
	} `yaml:"banner"`
	Exams []Exam `yaml:"exams"`
}

func (c *Configuration) readFile(filepath string) (*Configuration, error) {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return c, err
	}

	return c, nil

}

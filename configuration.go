package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Exam struct {
	Name         string `yaml:"name"`
	Description  string `yaml:"description"`
	Run          string `yaml:"run"`
	Possibilites []struct {
		Code    string `yaml:"code"`
		Message string `yaml:"message"`
	} `yaml:"possibilites"`
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

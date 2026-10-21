package main

import (
	"bufio"
	"embed"
	"encoding/json"
	"strings"
)

type License struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Reference string
}

//go:embed licenses.jsonl texts/*
var fs embed.FS

func listLicenses() {
	file, err := fs.Open("licenses.jsonl")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var licenses []License

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var license License

		err := json.Unmarshal(scanner.Bytes(), &license)
		if err != nil {
			continue
		}
		licenses = append(licenses, license)
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	for _, lis := range licenses {
		println(lis.Name)
	}
}

func searchLicense(id string) License {
	file, err := fs.Open("licenses.jsonl")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var license License
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		json.Unmarshal(scanner.Bytes(), &license)

		if license.ID == strings.ToUpper(id) {
			return license
		}
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
	return License{}
}

func getFile(id string) ([]byte, error) {
	return fs.ReadFile("texts/" + id + ".txt")
}

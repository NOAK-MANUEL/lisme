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
	Reference string `json:"reference"`
}

//go:embed licenses.jsonl texts/*
var fs embed.FS

func listLicenses() {
	file, err := fs.Open("licenses.jsonl")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var license License

		err := json.Unmarshal(scanner.Bytes(), &license)
		if err != nil {
			continue
		}
		println(license.Name)
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

}

func searchLicense(id string) License {
	file, err := fs.Open("licenses.jsonl")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var license License

		if err := json.Unmarshal(scanner.Bytes(), &license); err != nil {
			return License{}
		}

		if strings.EqualFold(license.ID, id) {
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

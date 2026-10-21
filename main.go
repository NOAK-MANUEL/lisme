package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	fl := pflag.NewFlagSet("lisme", pflag.ExitOnError)
	fl.Usage = func() {

		fl.PrintDefaults()
	}
	list := fl.Bool("list", false, "List all licenses")
	force := fl.Bool("force", false, "Force a license creation")
	help := fl.Bool("help", false, "help")
	search := fl.String("search", "", "List a license")

	err := fl.Parse(os.Args[1:])

	if err != nil {
		println("Error parsing flags:", err)
		fl.Usage()
		return
	}

	if *help {
		fl.Usage()
		return
	}
	if *list {
		listLicenses()
		return
	}
	if *search != "" && len(*search) >= 3 {
		lis := searchLicense(*search)
		if lis.ID == "" {
			println("License not found")
			return
		}
		fmt.Printf(
			`
			Name: %s,
			Reference: %s,
		


			`, lis.Name, lis.Reference)
		return
	}

	name := fl.Arg(0)
	lis := searchLicense(name)
	if lis.ID == "" {
		println("License not found")
		return
	}

	data, err := getFile(lis.ID)
	if err != nil {
		panic(err)
	}
	_, err = os.Stat("LICENSE")

	if errors.Is(err, os.ErrNotExist) || *force {
		mainData := []byte(lis.Name + "\n")
		mainData = append(mainData, data...)
		if err = os.WriteFile("LICENSE", mainData, 0644); err != nil {
			println("Couldn't write file:", err)
		} else {
			println("Licensed Successfully")
		}

	} else if err != nil {
		panic(err)
	} else {
		println("A LICENSE already exist. Use --force")

	}

}

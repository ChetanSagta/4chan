package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func readOutputFile(filename string) string {
	file_content, err := os.ReadFile(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	return string(file_content)
}

func get_filename_and_format(url string) (string, string) {
	println(url)
	regex_str := "^http:\\/\\/.*?([a-z0-9]+)\\.(\\w+)$"
	regex := regexp.MustCompile(regex_str)
	matches := regex.FindStringSubmatch(url)
	return matches[1], matches[2]
}

func saveToFile(url string) {
	filename, format := get_filename_and_format(url)
	data, err := download(url)
	if err != nil {
		fmt.Println("Error while downloading file: ", err)
		return
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		println("Couldn't find home directory")
		return
	}
	homedir = homedir + "/4chanImages/"
	if _, err := os.Stat(homedir); os.IsNotExist(err) {
		// Create the directory with read/write permissions for the current user
		err := os.MkdirAll(homedir, 0700)
		if err != nil {
			println("Error while creating directory")
			return
		}
	}
	file, err := os.Create(homedir + filename + "." + format)
	if err != nil {
		println("Error while creating file")
		file.Close()
	}
	println("Writing to file")
	file.Write(data)
}

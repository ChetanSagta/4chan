package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func readOutputFile(filename string) string{
	file_content, err := os.ReadFile(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	return string(file_content)
}

func get_filename_and_format(url string)(string, string){
	println(url)
	regex_str:= "^http:\\/\\/.*?([a-z0-9]+)\\.(\\w+)$"
	regex:= regexp.MustCompile(regex_str)
	matches:= regex.FindStringSubmatch(url)
	return matches[1], matches[2]
}

func saveToFile(url string){
	filename, format := get_filename_and_format(url)
	data, err := download(url)
	if(err != nil){
		fmt.Println("Error while downloading file: ", err);
		return;
	}

	file , err:= os.Create(filename+format)
	if err != nil {
		println("Error while creating file")
		file.Close()
	}
	file.Write(data)
}

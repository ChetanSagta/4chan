package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getResponseBody(url string) string {

	conn, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: {}", err)
	}
	body, err := ioutil.ReadAll(conn.Body)
	fmt.Println("Status: ", conn.Status)
	fmt.Println("Status Code: ", conn.StatusCode)
	return string(body)
}

func readOutputFile() string{
	file_content, err := os.ReadFile("output.html") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	return string(file_content)
}


func main() {
	file_content := readOutputFile()
	parse_html(file_content)
}

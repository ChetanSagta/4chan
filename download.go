package main

import (
	"fmt"
	"io"
	"net/http"
)

func download(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`)
	resp, err := client.Do(req)
	fmt.Println("Status: ", resp.Status)
	fmt.Println("Status Code: ", resp.StatusCode)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	status_code := resp.StatusCode
	if status_code != 200 {
		fmt.Println("Url: ", url)
		fmt.Println("StatusCode is not 200. Its:", status_code)
		return nil, err
	}
	body := resp.Body
	fmt.Println(body)

	bytes, error := io.ReadAll(body)
	if error != nil {
		fmt.Println("Error Reading Body:", error)
	}
	return bytes, nil
}

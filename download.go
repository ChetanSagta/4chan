package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

func download(url string) ([]byte, error) {
	resp , err := http.Get(url)
	if(err != nil){
		fmt.Println("Error:" , err)
		return nil,err;
	}

	status_code := resp.StatusCode
	if(status_code != 200){
		fmt.Println("Error:" , status_code)
		return nil, err
	}
	body := resp.Body
	bytes, error:= io.ReadAll(body)
	if(error != nil){
		fmt.Println("Error:" , error)
	}
	return bytes, nil
	
}

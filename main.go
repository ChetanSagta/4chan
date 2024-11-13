package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	nethtml "golang.org/x/net/html"
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

func readOutputFile(filename string) string{
	file_content, err := os.ReadFile(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	return string(file_content)
}

func get_file_url(root *nethtml.Node) string{
	file_tag := get_node_based_on_attr(root, "class", "file")
	fileText_tag := get_node_based_on_attr(file_tag, "class", "fileText")
	href := get_child_tag(fileText_tag, "a")
	url := "http:"+ get_attribute_value(href, "href")
	return url
}

func parseBoard(body *nethtml.Node){

	form := get_node_based_on_attr(body,"name","delform")

	threads := get_all_nodes_based_on_attr(form,"class", "thread");

	for thread := threads.Front(); thread != nil; thread = thread.Next(){
		temp_thread := thread.Value.(*nethtml.Node)
		original := get_node_based_on_attr(temp_thread, "class", "postContainer opContainer")
		url:= get_file_url(original)
		println(url)
		replies := get_all_nodes_based_on_attr(temp_thread, "class", "postContainer replyContainer")

		for reply:= replies.Front(); reply!= nil; reply= reply.Next(){
			temp_reply:= thread.Value.(*nethtml.Node)
			url= get_file_url(temp_reply)
			println(url)
		}
	}
}

func main() {
	file_content := readOutputFile("anime.html")
	html, err := get_html_from_file(file_content)
	if(err != nil){
		println(err)
		return;
	}

	body := get_child_tag(html,"body")
	parseBoard(body)
	

}

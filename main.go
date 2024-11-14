package main

import (
	nethtml "golang.org/x/net/html"
)


func parseBoard(body *nethtml.Node){

	form := get_node_based_on_attr(body,"name","delform")

	threads := get_all_nodes_based_on_attr(form,"class", "thread");

	for thread := threads.Front(); thread != nil; thread = thread.Next(){
		temp_thread := thread.Value.(*nethtml.Node)
		original := get_node_based_on_attr(temp_thread, "class", "postContainer opContainer")
		url:= get_file_url(original)
		saveToFile(url)
		replies := get_all_nodes_based_on_attr(temp_thread, "class", "postContainer replyContainer")

		for reply:= replies.Front(); reply!= nil; reply= reply.Next(){
			temp_reply:= thread.Value.(*nethtml.Node)
			url= get_file_url(temp_reply)
			saveToFile(url)
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

package main

import (
	"os"

	nethtml "golang.org/x/net/html"
)

var URL = ""

func parseBoard(body *nethtml.Node){

	form := get_node_based_on_attr(body,"name","delform")

	threads := get_all_nodes_based_on_attr(form,"class", "thread");

	for thread := threads.Front(); thread != nil; thread = thread.Next(){
		temp_thread := thread.Value.(*nethtml.Node)
		original := get_node_based_on_attr(temp_thread, "class", "postContainer opContainer")
		url:= get_file_url(original)
		saveToFile(url)
		replies := get_all_nodes_based_on_attr(temp_thread, "class", "postContainer replyContainer")
		summary_desktop := get_node_based_on_attr(temp_thread, "class","summary desktop")
		if(summary_desktop != nil){
			a_tag := get_child_tag(summary_desktop,"a")
			href := get_attribute_value(a_tag,"href")
			println("HREF: ", href)
			href = URL+href
			parsePage(href)
		}

		for reply:= replies.Front(); reply!= nil; reply= reply.Next(){
			temp_reply:= reply.Value.(*nethtml.Node)
			reply_url:= get_file_url(temp_reply)
			if reply_url != ""{
			 saveToFile(reply_url)
			}
		}
	}
}


func usage(){
	println("Usage: 4chan <url>")
}

func parsePage(url string){
	bytes, err := download(url)
	if(err != nil){
		println(err)
		return;
	}
	content := string(bytes)

	html, err := get_html_from_string(content)
	if(err != nil){
		println(err)
		return;
	}

	body := get_child_tag(html,"body")
	parseBoard(body)
}

func main() {

	commandlineargs := os.Args
	if len(commandlineargs) == 1 {
		usage()
		return;
	}

	URL = commandlineargs[1]
	
	parsePage(URL)

	println("Download Images Completed")

}

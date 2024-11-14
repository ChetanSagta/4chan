package main

import (
	"container/list"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func get_html_from_file(file_content string) (*html.Node, error){

	doc, err := html.Parse(strings.NewReader(file_content))
	if err != nil {
		fmt.Println("Error while parsing html document :", doc)
		return nil,err
	}
	return get_first_child_element(doc),nil
}

func get_all_boards(root *html.Node) *list.List{

	body := get_child_tag(root,"body")
	if body == nil {
		return nil 
	}

	boards:= get_node_based_on_id(body, "boards")
	board_list := get_all_nodes_based_on_attr(boards, "class", "boardlink")
	return board_list
}

func get_child_tag(root *html.Node,tagname string) *html.Node {

	firstChild := root.FirstChild
	for temp := firstChild; temp != nil; temp = temp.NextSibling {
		if temp.Type == html.ElementNode && temp.Data == tagname {
			return temp
		}

	}
	println("Couldn't find \"",tagname,"\"in the tree")
	return nil
}

func get_first_child_element(root *html.Node) *html.Node {

	firstChild := root.FirstChild
	for temp := firstChild; temp != nil; temp = temp.NextSibling {
		if temp.Type == html.ElementNode {
			return temp
		}
	}
	println("Couldn't find a child for the element")
	return nil
}

func get_all_child_element(root *html.Node) *list.List{

	firstChild := root.FirstChild
	elements :=  list.New()
	for temp := firstChild; temp != nil; temp = temp.NextSibling {
		if temp.Type == html.ElementNode {
			elements.PushBack(temp)
		}
	}
	return elements
}

func get_all_sibling_element(root *html.Node) *list.List{

	firstSibling:= root.NextSibling
	elements :=  list.New()
	for temp := firstSibling; temp != nil; temp = temp.NextSibling {
		if temp.Type == html.ElementNode {
			elements.PushBack(temp)
		}
	}
	return elements
}

func get_attribute_value(node *html.Node, attrib_name string) string{
		for _,attr := range node.Attr{
			if(attr.Key == attrib_name){
				return attr.Val 
			}
		}
	return ""
}

func get_node_based_on_attr(root *html.Node,attr_name string, attr_value string) *html.Node{
 
	if root == nil{
		return nil 
	}
	children := get_all_child_element(root)
	for child := children.Front(); child != nil; child= child.Next(){
		node := child.Value.(*html.Node)
		for _,attr := range node.Attr{
			if(attr.Key == attr_name && attr.Val == attr_value){
				return node
			}
		}
		node=get_node_based_on_attr(node, attr_name, attr_value);
		if(node != nil){
			return node
		}
	}
	return nil
}

func get_all_nodes_based_on_attr(root *html.Node,attr_name string, attr_value string) *list.List{

	all_nodes:= list.New()
	if root == nil{
		return all_nodes
	}

	children := get_all_child_element(root)
	for child := children.Front(); child != nil; child= child.Next(){
		node := child.Value.(*html.Node)
		for _,attr := range node.Attr{
			if(attr.Key == attr_name && attr.Val == attr_value){
				all_nodes.PushBack(node)
			}
		}
		all_nodes.PushBackList(get_all_nodes_based_on_attr(node, attr_name, attr_value));
	}
	return all_nodes 
}

func get_node_based_on_id(node *html.Node, value string) *html.Node{
	return get_node_based_on_attr(node, "id", value);
}

func print_attr(node *html.Node){
	attributes := node.Attr
	for index, attr := range attributes{
		fmt.Println(index, attr)
	}

}

func get_file_url(root *html.Node) string{
	file_tag := get_node_based_on_attr(root, "class", "file")
	fileText_tag := get_node_based_on_attr(file_tag, "class", "fileText")
	href := get_child_tag(fileText_tag, "a")
	url := "http:"+ get_attribute_value(href, "href")
	return url
}

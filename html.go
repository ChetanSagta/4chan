package main

import (
	"container/list"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func parse_html(file_content string) error {

	doc, err := html.Parse(strings.NewReader(file_content))
	if err != nil {
		fmt.Println("Error while parsing html document :", doc)
		return err
	}
	_html := get_first_child_element(doc)
	body := get_child_tag("body", _html)
	if body == nil {
		return err
	}
	// node := get_node_based_on_id(body, "opts-btn");
	node := get_node_based_on_attr(body, "class", "stat-cell")
	println("Node",node.Attr[0].Val)
	// firstchild := body.FirstChild
	// for temp := firstchild; temp != nil; temp = temp.NextSibling {
	// 	if temp.Type == html.ElementNode{
	// 		node:= get_node_based_on_id(temp,"hd");
	// 		if node != nil {
	// 			println(node.Data);
	// 		} else{
	// 			println("Element not found");
	// 		}
	// 	}
	// }
	return nil
}

func get_child_tag(tagname string, root *html.Node) *html.Node {

	firstChild := root.FirstChild
	for temp := firstChild; temp != nil; temp = temp.NextSibling {
		if temp.Type == html.ElementNode && temp.Data == tagname {
			return temp
		}
	}
	println("Couldn't find ", tagname, "in the tree")
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

func get_node_based_on_id(node *html.Node, value string) *html.Node{

	return get_node_based_on_attr(node, "id", value);
}

func print_attr(node *html.Node){
	attributes := node.Attr
	for index, attr := range attributes{
		fmt.Println(index, attr)
	}

}

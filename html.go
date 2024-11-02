package main

import (
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
	_html := get_child_element(doc)
	body := get_child_tag("body", _html)
	if body == nil {
		return err
	}
	firstchild := body.FirstChild
	for temp := firstchild; temp != nil; temp = temp.NextSibling {
		if temp.Type == html.ElementNode{
		fmt.Println("y: ", temp.Data)
		}
	}
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

func get_child_element(root *html.Node) *html.Node {

	firstChild := root.FirstChild
	for temp := firstChild; temp != nil; temp = temp.NextSibling {
		if temp.Type == html.ElementNode {
			return temp
		}
	}
	println("Couldn't find a child for the element")
	return nil
}

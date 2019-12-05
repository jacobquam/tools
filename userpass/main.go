package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/tomnomnom/rawhttp"
	"golang.org/x/net/html"
)

func scanLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func main() {

	// TODO: get URL from command line
	getResp, err := http.Get("<loginURL>")
	if err != nil {
		fmt.Println("Problem with URL")
	}
	defer getResp.Body.Close()

	doc, err := html.Parse(getResp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var csrf_token string
	var f func(*html.Node)
	f = func(n *html.Node) {
		// TODO: make none specific to form data structure
		if n.Type == html.ElementNode && n.Data == "input" {
			attrMap := map[string]string{}
			for _, a := range n.Attr {
				attrMap[a.Key] = a.Val
			}
			// attrMap["name"] needs to be changed to the correct name of the input
			// attribute. TODO Find a way to implement this automatically make dynamic
			if attrMap["type"] != "hidden" || attrMap["name"] != "<name of token>" {
				return
			}
			csrf_token = attrMap["value"]
			//at := n.Attr
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

	}
	f(doc)
	fmt.Println("Final Token Found: " + csrf_token)

	// TODO: get files from command line or master file
	// Will be username file
	usernames, err := scanLines("path/to/username/file")
	if err != nil {
		fmt.Println(err)
	}

	// will be password file
	passwords, err := scanLines("path/to/password/file")
	if err != nil {
		fmt.Println(err)
	}

	// TODO: create gorutines and concurent threads
	for _, user := range usernames {
		for _, pass := range passwords {
			req, err := rawhttp.FromURL("POST", "<loginURL>")
			if err != nil {
				fmt.Println(err)
			}
			// TODO: make more general to form POST data
			req.Body = "useralias=" + user + "&password=" + pass + "&submitLogin=Connect&name_of_token=" + csrf_token
			fmt.Printf("%s\n", req.Body)
		}
	}

}

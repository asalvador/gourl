package main

import (
	"fmt"
	"os"

	"github.com/asalvador/gourl"
)

func main() {
	var u string
	if len(os.Args) > 1 {
		u = os.Args[1]
	} else {
		fmt.Println("url is required")
		return
	}

	url, err := gourl.Parse(u)

	if err != nil {
		fmt.Println("error enountered while parsing", u)
		return
	}

	fmt.Println("URL: ", url.URL)
	fmt.Println("Normalized URL: ", url.String())
	fmt.Println("Scheme: ", url.Scheme)
	fmt.Println("User: ", url.User)
	fmt.Println("Password: ", url.Password)
	fmt.Println("Subdomain: ", url.Subdomain)
	fmt.Println("Domain: ", url.Domain)
	fmt.Println("Hostname: ", url.Hostname)
	fmt.Println("Port: ", url.Port)
	fmt.Println("Path: ", url.Path)
	fmt.Println("Query: ", url.Query)
	fmt.Println("Fragment: ", url.Fragment)

	return
}

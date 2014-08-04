package main

import (
	"fmt"
	"os"

	"github.com/asalvador/gourl/"
)

func main() {
	var u string
	if len(os.Args) > 1 {
		u = os.Args[1]
	} else {
		fmt.Println("url is required")
		return
	}

	gourl.SetTLDFile("../tlds")
	url := gourl.NewURL(u)

	fmt.Println("URL: ", url.URL)
	fmt.Println("Scheme: ", url.Scheme())
	fmt.Println("User: ", url.User())
	fmt.Println("Password: ", url.Password())
	fmt.Println("Subdomain: ", url.Subdomain())
	fmt.Println("Domain: ", url.Domain())
	fmt.Println("Hostname: ", url.Hostname())
	fmt.Println("Path: ", url.Path())
	fmt.Println("Query: ", url.Query())
	fmt.Println("Fragment: ", url.Fragment())

	return
}

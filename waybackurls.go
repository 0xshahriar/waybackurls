package main

import (
"fmt"
"net/http"
"io/ioutil"
"strings"
"os"
)

func main() {
var domain string
fmt.Print("Enter the domain name : ")
fmt.Scan(&domain)

file, err := os.Create(domain + ".txt")
if err != nil {
	fmt.Println(err)
	return
}
defer file.Close()

// construct the url for the Wayback Machine API
url := "http://web.archive.org/cdx/search/cdx?url=*." + domain + "&output=txt&fl=original&collapse=urlkey"

// make the request to the Wayback Machine API
response, err := http.Get(url)
if err != nil {
	fmt.Println(err)
	return
}

// read the response body
body, err := ioutil.ReadAll(response.Body)
if err != nil {
	fmt.Println(err)
	return
}

// convert the response body to a string
bodyString := string(body)

// split the response body into lines
lines := strings.Split(bodyString, "\n")

// write the URLs to the file
for _, line := range lines {
	if line != "" {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

fmt.Println("URLs for " + domain + " have been written to " + domain + ".txt")
}

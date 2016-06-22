// Fetch prints the content found at a URL
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func printerror(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}

func print(something interface{}) {
	fmt.Fprintf(os.Stdout, "%v\n", something)
}

func dealWith(err interface{}, format string, args ...interface{}) {
	if (err != nil) {
		printerror(format, args...)
		os.Exit(1)
	}
}

func main() {
	for _, url := range os.Args[1:2] {

		// fetch the page first
		resp, err := http.Get(url)
		dealWith(err, "fetch: %s, %v\n", url, err)

		// pull out the contents
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		dealWith(err, "fetch: reading %s: %v\n", url, err)

		fmt.Printf("%s", b)
	}
}
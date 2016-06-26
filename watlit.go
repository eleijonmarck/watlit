package main

import (
	"fmt"
	"os"
	// to debug the program
	"io/ioutil"
	"net/http"
)

func main() {
	//sampleYoutubeVideo := "https://www.youtube.com/watch?v=OFnQmD_CP0w"

	youtubeHTML := os.Args[1]
	fmt.Println("Downloading video...")
	fmt.Println(youtubeHTML)
	resp, err := http.Get(youtubeHTML)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("Response from website..")
	fmt.Println(body)
}

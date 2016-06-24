package main

import (
	"fmt"
	"os"
	// to debug the program
)

func main() {
	//sampleYoutubeVideo := "https://www.youtube.com/watch?v=OFnQmD_CP0w"

	youtubeHTML := os.Args[1]
	fmt.Println("Downloading video...")
	fmt.Println(youtubeHTML)
}

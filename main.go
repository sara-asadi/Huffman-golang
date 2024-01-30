package main

import (
	"os"
)

func main() {

	fileName := os.Args[1]

	var z Zipper
	z.Zip(fileName)
	z.Unzip("zipped-" + fileName)

	// fmt.Println(zipper.FreqDic)

	// // convert bytes to string
	// str := string(text)

	// // show file data
	// fmt.Println(str)
}

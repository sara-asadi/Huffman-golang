package main

import "os"

func main() {

	fileName := os.Args[1]

	var z Zipper
	z.Zip(fileName)
	z.Unzip("zipped-" + fileName)
}

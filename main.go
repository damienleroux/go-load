package main

import (
	"flag"
	"fmt"
	"go-load/cmd/loadstream"
	"go-load/cmd/loadwhole"
)

var testFilesFolderPath = flag.String("testfile", "", "path to mock file")

func getApply(count *int) func(a ...interface{}) (n int, err error) {
	return func(a ...interface{}) (n int, err error) {
		(*count)++
		return fmt.Println(">>", a)
	}
}

// test with "go run .\main.go -testfile C:/Users/damien/Documents/workspace/go-load/mock/sample2K.log"
func main() {
	flag.Parse()
	fmt.Printf("read file \"%v\"\n", string(*testFilesFolderPath))
	numberOfPrintForLoadWholeFile := 0
	numberOfPrintForLoadStreamFile := 0

	loadwhole.LoadWholeFile(*testFilesFolderPath, getApply(&numberOfPrintForLoadWholeFile))
	loadstream.LoadStreamFile(*testFilesFolderPath, getApply(&numberOfPrintForLoadStreamFile))

	fmt.Printf("LoadWholeFile treated %v file\n", numberOfPrintForLoadWholeFile)
	fmt.Printf("LoadStreamFile treated %v lines\n", numberOfPrintForLoadStreamFile)
}

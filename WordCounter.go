package main

// Run go mod tidy to handle imports
import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func wordCounter() {

}

func reducer()  {

}

func main() {

	// Check If Input File Exists in same dir
	if _ , err := os.Stat("./test.txt") ; os.IsNotExist( err ) {
		fmt.Println("test.txt doesn't exist")
		os.Exit(100)
	}

	// DataIn Reading
	data , err := ioutil.ReadFile("./test.txt")
	check(err)

	dataIn := string( data )
	fmt.Print( strings.ToLower(dataIn) )
}
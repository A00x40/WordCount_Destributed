package main

// Run go mod tidy to handle imports
import (
	"fmt"
	"io/ioutil"
	"os"	
	"strings"
	"sort"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func wordCounter( n int , words []string ) (map[string]int){
	m := make( map[string]int )

	for _ , v := range(words) { m[v]++ }
	// for k , v := range(m) { fmt.Println(k , ":" ,  v) }
	return m
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
	
	dataIn := strings.ToLower( string( data ) )
	var words []string
	
	var lines []string = strings.Split(dataIn, "\n")
	for _, line := range(lines) {
		line = strings.Trim(line, "\n\r")
		
		words = append(words, strings.Split(line, " ")...)
	}
	
	
	m := wordCounter(len(words), words)

	
	// Create output file
	f, err := os.Create("WordOutput.txt")
	defer f.Close()
	
	// Struct of pair (key, value)
	type kv struct {
		Key   string
        Value int
    }

	// Create slice of key-value pairs of map items to sort it
    var ss []kv
    for k, v := range m {
		ss = append(ss, kv{k, v})
    }
    sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value || ss[i].Value == ss[j].Value && ss[i].Key < ss[j].Key
    })

	// Write to file
	for _ , v := range(ss) { fmt.Fprintln(f, v.Key, ":" ,  v.Value) }
}
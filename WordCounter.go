package main

// Run go mod tidy to handle imports
import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sort"
	"sync"
)

type sharedMap struct {
	mu sync.Mutex
	ma map[string]int
	wg sync.WaitGroup
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func wordCounter(s *sharedMap, words []string, i int ) {
	defer s.wg.Done()
	for _ , v := range(words) { 
		s.mu.Lock()
		s.ma[v]++ 
		s.mu.Unlock()
	}
}

func reducer(s *sharedMap)  {
	// wait for goroutines to finish
	s.wg.Wait()

	// Create output file
	f, err := os.Create("WordCountOutput.txt")
	check(err)
	defer f.Close()
	
	// Struct of pair (key, value)
	type kv struct {
		Key   string
        Value int
    }

	// Create slice of key-value pairs of map items to sort it
    var ss []kv
    for k, v := range s.ma {
		ss = append(ss, kv{k, v})
    }

    sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value || ss[i].Value == ss[j].Value && ss[i].Key < ss[j].Key
    })

	// Write to file
	for _ , v := range(ss) { fmt.Fprintln(f, v.Key, ":" ,  v.Value) }
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
	
	// Split the string to lines
	var lines []string = strings.Split(dataIn, "\n")
	
	// Splite the lines to words
	for _, line := range(lines) {
		line = strings.Trim(line, "\n\r")
		words = append(words, strings.Split(line, " ")...)
	}
	
	s := sharedMap {ma: make(map[string]int)}
	
	// wait group to wait for all goroutines
	s.wg.Add(5)

	mult := len(words) / 5
	go wordCounter(&s, words[:mult],1)
	go wordCounter(&s, words[mult:2*mult],2)
	go wordCounter(&s, words[2*mult:3*mult],3)
	go wordCounter(&s, words[3*mult:4*mult],4)
	go wordCounter(&s, words[4*mult:],5)

	reducer(&s)
}
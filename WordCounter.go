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

func wordCounter( words []string ) {
	m := make( map[string]int )

	// TODO Add Channel & make it send the map to reducer after counting
	for _ , w := range(words) { m[w]++ }
}

func reducer()  {

	// TODO Receive from wordCounter and merge maps to sharedMap
	// Sort SharedMap & Write to OutFile

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

	var i  , j , n = 0 , 0 , len(dataIn)
	for ; i < n ; {
		switch dataIn[i] {
			case ' ' :
				// First Char Whitespace
				if i == 0 { 
					i++
					j++ 

				// Word then Whitespace
				} else if dataIn[i-1] != ' ' { 
					words = append( words , dataIn[j:i] )
					i++
					j = i

				// Whitespace then Whitespace
				} else if dataIn[i-1] == ' ' { i++ }
		
			// Current Not a Whitespace
			default  :
				if i == n-1 { 
					words = append( words , dataIn[j:i+1] ) 
					i++
				} else if dataIn[i] == '\n' {
					words = append( words , dataIn[j:i] )
					i++
					j = i
				} else { i++ }
		}
	}
		
	// Array of Channels to be Used
	var ch [5] chan map[string]int
	for i := 0 ; i < 5 ; i++ { ch[i] = make(chan map[string]int) }

	// Dividing work & Counting Freqs

	// Note Main Function Completes before go routine
	// Use Sync WaitGroup to wait for go routine 

	// You won't be able to see prints inside (go wordCounter) so remove go for testing
	go wordCounter( words[ 0 : (len(words) / 5) ]    )
	go wordCounter( words[ (len(words) / 5)     : 2 * (len(words) / 5) ] )
	go wordCounter( words[ 2 * (len(words) / 5) : 3 * (len(words) / 5) ] )
	go wordCounter( words[ 3 * (len(words) / 5) : 4 * (len(words) / 5) ] )
	go wordCounter( words[ 4 * (len(words) / 5) : ]  )	
}
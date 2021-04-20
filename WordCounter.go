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

func wordCounter( n int , words []string ) {
	m := make( map[string]int )

	for _ , v := range(words) { m[v]++ }
	for k , v := range(m) { fmt.Println(k , ":" ,  v) }

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
				if i == n-1 { words = append( words , dataIn[j:i+1] ) }
				i++
		}
	}

	// Dividing work & Counting Freqs
	for i := 0 ; i < 4 ; i++ { 
		wordCounter( len(words) / 5 , words[ i * (len(words) / 5) : (i+1) * (len(words) / 5) ] )
	}
	
	wordCounter( len(words) / 5 + len(words) % 5 , words[ 4 * (len(words) / 5) : ] )
}
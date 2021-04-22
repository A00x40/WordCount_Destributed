package main

// Run go mod tidy to handle imports
import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"sort"
)

type sharedMap struct {
	mp map[string] int
	mu sync.Mutex
	wg sync.WaitGroup
}

func ( smap *sharedMap ) Inc( key string , value int ) {
	smap.mu.Lock()
	defer smap.mu.Unlock()
	smap.mp[key] += value
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func wordCounter( words []string , c chan map[string]int ) {
	mp := make(map[string]int)
	for _ , w := range(words) { mp[w]++ }
	c <- mp
	close(c)
}

func reducer( smap *sharedMap , in [5] chan map[string]int )  {

	f , err := os.OpenFile("./WordCountOutput.txt" , os.O_RDWR | os.O_TRUNC , 0755 )
	check( err )
	defer f.Close()

	for i := 0 ; i < 5 ; i++ {
		smap.wg.Add(1)
		go func(in chan map[string]int) {
			defer smap.wg.Done()
			for mp := range(in)  {
				for k , v := range(mp) { smap.Inc(k ,v) }
			}

		} (in[i])
	}
	smap.wg.Wait()
	
	// Struct of pair (key, value)
	type kv struct {
		Key   string
        Value int
    }

	// Create slice of key-value pairs of map items to sort it
    var ss []kv
    for k, v := range smap.mp {
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
		
	var ch [5] chan map[string]int
	for i := 0 ; i < 5 ; i++ { ch[i] = make(chan map[string]int) }

	// Dividing work & Counting Freqs
	go wordCounter( words[ 0 : (len(words) / 5) ] , ch[0] )
	go wordCounter( words[ (len(words) / 5)     : 2 * (len(words) / 5) ] , ch[1] )
	go wordCounter( words[ 2 * (len(words) / 5) : 3 * (len(words) / 5) ] , ch[2] )
	go wordCounter( words[ 3 * (len(words) / 5) : 4 * (len(words) / 5) ] , ch[3] )
	go wordCounter( words[ 4 * (len(words) / 5) : ] , ch[4] )

	smap := sharedMap{ mp : map[string]int{} }

	// Check If Output File exists and create it if not
	if _ , err := os.Stat("./WordCountOutput.txt"); os.IsNotExist(err) {
		_ , err := os.Create("./WordCountOutput.txt")
		check( err )	
	} 

	reducer( &smap , ch )
}
package search

import (
	"fmt"
	"log"
)

// Result contains the result of a search
type Result struct{
Field string
Content string
}

// Matchers define the behavior required by the typee
// that want to implement a new search type.
type Matcher interface{
Search(feed *Feed, searchTearm string)([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan <- *Result){
// Perform the search against the specified matcher
searchResults, err := matcher.Search(feed, searchTerm)
    if err != nil{
log.Println(err)
return
    }

    // Write the results to the channel
for _, result := range searchResults{
results <- result
    }
}

// Display writesresults to the terminal window as they
// are received by individual goroutines.
func Display(results chan *Result){
// The channel blocks until a  message is written to the
// channel. Once the channel is closed, the for loop terminates
for result := range results{
fmt.Printf("%s:%s\n\n", result.Field,result.Content)
    }
}

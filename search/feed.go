package search

import(
    "encoding/json"
    "os"
)

const dataFile = "data/data.json"

// Feed contains information we need to process a feed
type Feed struct{
    Name string `json:"site"`
    URI string `json:"link"`
    Type string `json:"type"`
}

// RetrieveFeeds reads and unmarshals the feed data file
func RetrieveFeeds()([]*Feed, error){
// Open the file
file, err := os.Open(dataFile)
if err != nil {
return nil, err
    }
// Schedule the file to be closed once the function
    //returns
    defer file.Close()
// Decode the file to a slice of pointers to the Feed
    // value
    var feeds []*Feed
    err = json.NewDecoder(file).Decode(&feeds)
    // no need to check for err at this level, the caller can do that
return feeds, err
}

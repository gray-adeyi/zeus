package main

import (
    "log"
    "os"
    _ "zeus/matchers"
    "zeus/search"
)

func init(){
    // Change the the device for logging to the stdout
    log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main(){
    // Performs the search for the specified term.  
    search.Run("president")
}

package main

import (
    "code.google.com/p/go-tour/wc"
    //"fmt"
)
import "strings"

func WordCount(s string) map[string]int {
    data := make(map[string]int);
    for _,word := range strings.Fields(s) {
        //fmt.Println(word);
        data[ word ] += 1;
    }
    return data
}

func main() {
    wc.Test(WordCount)
}

package main

import (
    . "fmt"
    Str "strings"
)


func main() {
    Println(Str.Join([]string { "Go", "lang", 
                                "programming" }, 
                     "-"));
    //prints "Go-lang-programming"

    Println(Str.ReplaceAll("Hello Me!", "Me", 
                           "world"));
    //prints "Hello world!"

    Println(Str.HasPrefix("This is a sentence.", 
                          "This"));
    //prints "true"

    Println(Str.HasSuffix("This is a sentence.", 
                          "sentence."));
    //prints "true"

    Println(Str.ToUpper("This is a sentence."))
    //prints "THIS IS A SENTENCE."

    Println(Str.ToLower("AN ALL CAPITALS"+
                        " SENTENCE."));
    //prints "an all capitals sentence."

    Println(Str.Split("Go - lang - programming", 
                      " - "));
    //prints "[Go lang programming]"
}
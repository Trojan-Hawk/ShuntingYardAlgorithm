package main

import (
    "fmt"
)

func intopost(infix string) string {
    // creating a map of special characters and assigning them a value
    specials := map[rune]int{'*': 10, '.': 9, '|': 8}

    // a rune is a character as it's displayed on the screen (utf8)
    postfix := []rune{} // initialise an array of runes
    stack := []rune{}
    
    

    return string(postfix)
}

func main() {
    // Answer: ab.c*.
    fmt.Println("Infix:   ", "a.b.c*")
    fmt.Println("Postfix: ", inpost("a.b.c*"))

    // Answer: abd|.*
    fmt.Println("Infix:   ", "(a.(b|d))*")
    fmt.Println("Postfix: ", inpost("(a.(b|d))*"))

    // Answer: abd|.c*
    fmt.Println("Infix:   ", "(a.(b|d)).c*")
    fmt.Println("Postfix: ", inpost("(a.(b|d)).c*")
    
    // Answer: abb.+.c.
    fmt.Println("Infix:   ", "a.(b.b)+.c")
    fmt.Println("Postfix: ", inpost("a.(b.b)+.c"))
}
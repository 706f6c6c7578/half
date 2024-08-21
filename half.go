package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Read the string from stdin
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = input[:len(input)-1] // Remove the newline character

    // Calculate the length and the middle of the string
    length := len(input)
    mid := length / 2

    // Split the string into two halves
    firstHalf := input[:mid]
    secondHalf := input[mid:]

    // Print the two halves
    fmt.Println("lh:", firstHalf)
    fmt.Println("rh:", secondHalf)
}


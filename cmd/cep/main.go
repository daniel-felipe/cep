
package main

import (
    "fmt"
    "os"
    "strings"
    "bufio"

    "daniel/cep"
)

func getCep() string {
    if (len(os.Args) >= 2) {
        return os.Args[1]
    }

    scanner := bufio.NewReader(os.Stdin)
    cep, err := scanner.ReadString('\n')
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v", err) 
        os.Exit(1)
    }
    return strings.Trim(cep, "\n")
}

func main() {
    address := cep.NewAddress(getCep())
    address.Show()
}


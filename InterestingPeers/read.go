package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "os"
)

//Name is Struct with names in it
type Name struct {
    fName string
    lName string
}

func main() {

    fmt.Printf("Enter a File Name:")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    fname := scanner.Text()

    file, err := os.Open(fname)

    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    r := csv.NewReader(file)
    r.Comma = ' ' // Space-delimited

    nslice := []*Name{}
    for {
        n, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("Error:", err)
            break
        }

        nam := new(Name)
        nam.fName = n[0]
        nam.lName = n[1]
        nslice = append(nslice, nam)
    }
    for i := range nslice {
        name := nslice[i]
        fmt.Println(name)
    }
}

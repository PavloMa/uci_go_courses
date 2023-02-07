package User1

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    type name struct {
        fname string
        lname string
    }

    var temp string
    names := make([]name, 0)

    fmt.Println("Enter the name of a text file:")
    fmt.Scan(&temp)

    file, _ := os.Open(temp)
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        s := strings.Split(scanner.Text(), " ")

		var temp name

		temp.fname, temp.lname = s[0], s[1]
		names = append(names, temp)
    }

    file.Close()

    for _, v := range names {
        fmt.Println(v.fname, v.lname)
    }
}

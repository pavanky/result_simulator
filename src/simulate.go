package main

import (
    "fmt"
    "os"
    "io"
    "encoding/csv"
)

func readContents(dir string, name string) [][]string {
    file, err := os.Open(dir + "/" + name)
    defer file.Close()

    if err != nil {
        panic(err)
    }

    reader := csv.NewReader(file)
    line, err := reader.Read()
    contents := [][]string {line}

    for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }
        contents = append(contents, line)
    }

    return contents
}

func showContents(contents [][]string) {
    for _, line := range contents {
        fmt.Println(line)
    }
}

func main() {
    dir := os.Args[1]

    matches := readContents(dir, "matches.csv")
    showContents(matches)
}

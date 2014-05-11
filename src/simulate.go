package main

import (
    "fmt"
    "os"
    "io"
    "strconv"
    "encoding/csv"
)

type Standing struct {
    name string
    matches int
    points float64
}


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

func currentStandings(matches [][]string) map[string]Standing {

    var m map[string]Standing
    m = make(map[string]Standing)

    for _, match := range matches {
        for i := 0; i < 2; i++ {
            team := match[i]
            pts := match[i + 2]

            if (pts == "X") {
                break
            }

            points, err := strconv.ParseFloat(pts,32)
            if err != nil {
                panic(err)
            }

            c := m[team]
            c.matches = c.matches + 1
            c.points = c.points + points
            c.name = team
            m[team] = c
        }
    }

    return m
}

func showStandings(standings map[string]Standing) {

    fmt.Printf("Team Name\t Matches\t Points\t\n")
    for key, val := range standings {
        fmt.Printf("%s\t\t %2d\t\t %.0f\t\n", key, val.matches, val.points)
    }
}

func main() {
    dir := os.Args[1]

    matches := readContents(dir, "matches.csv")
    standings := currentStandings(matches[1:])
    showStandings(standings)
}

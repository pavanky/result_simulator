package main

import (
    "fmt"
    "os"
    "io"
    "strconv"
    "encoding/csv"
)

type Team struct {
    name string
    division string
    matches int
    points float64
    priority float64
    id int
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

func currentStandings(m map[string]Team, matches [][]string) map[string]Team {
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
            m[team] = c
        }
    }
    return m
}

func showStandings(standings map[string]Team) {
    fmt.Printf("Team Name\t Matches\t Points\t\n")
    for _, val := range standings {
        fmt.Printf("%s\t\t %2d\t\t %.0f\t\n", val.name, val.matches, val.points)
    }
}

func parseTeams(teams [][]string) map[string]Team {
    var m map[string]Team
    m = make(map[string]Team)
    for idx, entry := range teams {
        team := entry[0]
        prty, err := strconv.ParseFloat(entry[2],32)
        if err != nil {
            panic(err)
        }

        c := m[team]
        c.name = team
        c.division = entry[1]
        c.id = idx - 1
        c.priority = prty
        m[team] = c
    }
    return m
}

func main() {
    dir := os.Args[1]
    teams := readContents(dir, "teams.csv")
    matches := readContents(dir, "matches.csv")

    tmap := parseTeams(teams[1:])
    standings := currentStandings(tmap, matches[1:])
    showStandings(standings)
}

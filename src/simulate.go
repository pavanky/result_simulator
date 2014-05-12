package main

import (
    "fmt"
    "os"
    "io"
    "sort"
    "time"
    "strconv"
    "math/rand"
    "encoding/csv"
)

type Team struct {
    name string
    division string
    matches int
    points float64
    priority float64
    id int
    final []float64
}

type Points struct {
    win, lose, draw, other float64
}

type TeamArray []Team
type TeamMap map[string]Team

func (a TeamArray) Len() int           { return len(a) }
func (a TeamArray) Swap(i, j int)      { a[i],a[j] = a[j],a[i] }

func (a TeamArray) Less(i, j int) bool {
    if a[i].points == a[j].points {
        return a[i].priority > a[j].priority
    }
    return a[i].points > a[j].points
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

func parseTeams(teams [][]string) TeamMap {
    var m TeamMap
    m = make(TeamMap)
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

func showStandings(standings TeamMap) {
    teams := []Team{}
    for _,team := range standings {
        teams = append(teams, team)
    }

    sort.Sort(TeamArray(teams))

    fmt.Printf("Team Name\t Matches\t Points\t\n")
    for _, val := range teams {
        fmt.Printf("%s\t\t %2d\t\t %.0f\t\n", val.name, val.matches, val.points)
    }
}

func getWinProb(a Team, b Team) float64 {
    // FIXME: Use an actual table
    return 0.5
}

func getPoints(a Team, b Team, pts Points) (float64, float64) {
    prob := getWinProb(a, b)
    val := rand.Float64()

    if (val < prob) {
        return pts.win, pts.lose
    } else {
        return pts.lose, pts.win
    }

}

func getStandings(m TeamMap, matches [][]string, pts Points, final bool) TeamMap {
    for _, match := range matches {
        a := match[0]
        apts := match[2]
        ateam := m[a]

        b := match[1]
        bpts := match[3]
        bteam := m[b]

        if (apts == "X") || (bpts == "X") {
            if final == false {
                continue
            }
            apoints, bpoints := getPoints(ateam, bteam, pts)
            ateam.points = ateam.points + apoints
            bteam.points = bteam.points + bpoints
        } else {
            apoints, aerr := strconv.ParseFloat(apts,32)
            if aerr != nil {
                panic(aerr)
            }

            bpoints, berr := strconv.ParseFloat(bpts,32)
            if berr != nil {
                panic(berr)
            }
            ateam.points = ateam.points + apoints
            bteam.points = bteam.points + bpoints
        }

        ateam.matches = ateam.matches + 1
        m[a] = ateam

        bteam.matches = bteam.matches + 1
        m[b] = bteam
    }

    return m
}

func main() {
    dir := os.Args[1]
    rand.Seed(time.Now().UTC().UnixNano())

    teams := readContents(dir, "teams.csv")
    matches := readContents(dir, "matches.csv")

    // FIXME: Read this from a file
    pts := Points{2, 0, 0, 0}

    fmt.Println("Current Standings")
    tmap := parseTeams(teams[1:])
    tmap  = getStandings(tmap, matches[1:], pts, false)
    showStandings(tmap)

    fmt.Println("Final Standings")
    tmap = parseTeams(teams[1:])
    tmap = getStandings(tmap, matches[1:], pts, true)
    showStandings(tmap)
}

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
type TeamMap map[string]int

var tm TeamMap

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

func parseTeams(tdata [][]string) TeamArray {
    tm = make(TeamMap)
    nteams := len(tdata)
    teams := make(TeamArray, nteams)

    for idx, entry := range tdata {
        name := entry[0]
        prty, err := strconv.ParseFloat(entry[2],32)
        if err != nil {
            panic(err)
        }

        tm[name] = idx

        var t Team
        t.name = name
        t.division = entry[1]
        t.id = idx
        t.priority = prty
        t.points = 0
        t.final = make([]float64, nteams, nteams)

        teams[idx] = t
    }
    return teams
}

func getWinProb(a Team, b Team) float64 {
    // FIXME: Use an actual table
    return 0.5
}

func predictPoints(a Team, b Team, pts Points) (float64, float64) {
    prob := getWinProb(a, b)
    val := rand.Float64()

    if (val < prob) {
        return pts.win, pts.lose
    } else {
        return pts.lose, pts.win
    }

}

func getPoints(teams TeamArray, matches [][]string, pts Points, final bool) TeamArray {
    for _, match := range matches {
        a := match[0]
        apts := match[2]
        ateam := teams[tm[a]]

        b := match[1]
        bpts := match[3]
        bteam := teams[tm[b]]

        if (apts == "X") || (bpts == "X") {
            if final == false {
                continue
            }
            apoints, bpoints := predictPoints(ateam, bteam, pts)
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
        teams[tm[a]] = ateam

        bteam.matches = bteam.matches + 1
        teams[tm[b]] = bteam
    }

    return teams
}

func showStandings(teams TeamArray) {

    sort.Sort(TeamArray(teams))
    fmt.Printf("Team Name\t Matches\t Points\t\n")
    for _, val := range teams {
        fmt.Printf("%s\t\t %2d\t\t %.0f\t\n", val.name, val.matches, val.points)
    }
    fmt.Printf("\n")
}

func getRanks(teams TeamArray) TeamArray {
    sort.Sort(TeamArray(teams))

    for idx, team := range teams {
        tm[team.name] = idx
        team.final[idx] = team.final[idx] + 1
        team.points = 0
        teams[idx] = team
    }

    return teams
}

func showProbDist(teams TeamArray, num int) {

    fmt.Printf("Team\t\t")
    for i := 0; i < len(teams); i++ {
        fmt.Printf("Pos %d\t\t", i + 1)
    }
    fmt.Printf("\n")

    for _, team := range teams {
        fmt.Printf("%s\t\t", team.name)
        for i := 0; i < len(team.final); i++ {
            fmt.Printf("%2.2f\t\t", 100*team.final[i] / float64(num))
        }
        fmt.Printf("\n")
    }
}

func main() {
    dir := os.Args[1]
    rand.Seed(time.Now().UTC().UnixNano())

    tdata := readContents(dir, "teams.csv")
    mdata := readContents(dir, "matches.csv")

    // FIXME: Read this from a file
    pts := Points{2, 0, 0, 0}

    fmt.Println("Current Standings")
    teams := parseTeams(tdata[1:])
    teams  = getPoints(teams, mdata[1:], pts, false)
    showStandings(teams)

    teams = parseTeams(tdata[1:])

    nruns := 100000
    for i := 0; i < nruns; i++ {
        teams = getPoints(teams, mdata[1:], pts, true)
        teams = getRanks(teams)
    }

    showProbDist(teams, nruns)
}

## Torunament result simulator using go

I have created this package to test out Go. It will naively predict the results of a tournament using apriori probabilities. I am trying to make it general enough that it can be used for any torunament / league.

### Data format

The required needs to be in a csv format. All necessary files need to be in the same directory.

**Contents and format**

- teams.csv: team_name,division,priority
    - if all teams are in same division, use a place holder
    - priority values are used for sorting when teams are tied
    - An example for priority is the net run rate from cricket
    - You can generate this file manually or writing your own program

- matches.csv: team_A,team_B,points_A,points_B
    - use 'X' for points if the match hasn't been played
    - if a match has been played fill in the points earned by the each team

- probability.csv:
    - Each row is probability of winning against the other teams
    - Team order should be the same as the one in teams.csv

- playoffs.csv: Match_ID, Division_standing, Division_standing
    - Match_ID is a unique token used for the match
    - Since teams making the play offs are unknown, provide Division name and their standing

**Build and run**

I developed this on Linux. It should work without a problem on other OS as well.

- Build
    - run `make` to build
    - alternatively `go build src/simulate.go` should work also

- Run
   - `./simulate path/to/data/directory`

**Output**

Work in progress. Currently only reading the match information and showing current standing.

Here is a sample output for IPL 2014 after 35 matches:

```
Team Name        Matches         Points
KXIP              9              14
CSK               9              14
RR                9              12
KKR               9              8
SRH               8              8
RCB               9              6
MI                8              4
DD                9              4
```


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

Work in progress. Currently displays the predicted results after 100000 runs.

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

Team    Pos 1   Pos 2   Pos 3   Pos 4   Pos 5   Pos 6   Pos 7   Pos 8
KXIP    54.60   30.55   11.69   2.95    0.21    0.00    0.00    0.00
CSK     34.92   45.99   14.86   3.77    0.46    0.00    0.00    0.00
RR      9.74    19.71   54.61   13.33   2.21    0.38    0.01    0.00
SRH     0.68    2.96    11.02   34.22   33.74   11.90   4.67    0.81
RCB     0.00    0.02    0.79    5.22    19.29   40.19   24.62   9.87
KKR     0.06    0.74    6.30    37.64   31.80   14.96   6.97    1.54
DD      0.00    0.00    0.01    0.38    3.30    11.58   23.31   61.41
MI      0.00    0.02    0.72    2.49    9.00    20.99   40.42   26.36
```

// Package challenge19 is located https://adventofcode.com/2020/day/19
package challenge19

import (
	"AdventOfCode/file_utils"
	"fmt"
	"strings"
)

var aKey = ""
var bKey = ""

var regexMatcher = regexp.MustCompile(`[ab\|\(\)]*`)


func Run() int {
    rules, _ := getRulesAndPotentialMatches()
    _ = buildRegexFromRules(rules["0"], rules)
    return 0
}

func getRulesAndPotentialMatches() (map[string]string, []string) {
    rows := file_utils.OpenFileIntoSlice("./challenge19/test_input.txt")
    rules := make(map[string]string)
    potentialMatches := make([]string, 0)
    inRules := true
    for _, r := range rows {
        if r == "" {
                inRules = false
                continue
        }
        if inRules {
            splitted := strings.Split(r, ": ")
            key := splitted[0]
            val := strings.Trim(splitted[1], "\"")
            rules[key] = val
            fmt.Println(val)
            if val == "a" {
                aKey = key
            } else if val == "b" {
                bKey = key
            }
        } else {
            potentialMatches = append(potentialMatches, r)
        }
    }
    return rules, potentialMatches
}

func buildRegexFromRules(rule string, rules map[string]string) string {
    if
}

package challenge4

import (
	"AdventOfCode/file_utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	rows := buildPassportRows()
	passports := make([]Passport, 0)
	blankPassport := Passport{}
	for _, row := range rows {
		p := newPassport(row)
		if p != blankPassport {
			passports = append(passports, p)
		}
	}
	valid := len(passports)
	fmt.Println(valid)
	fmt.Println(partTwoValidPassportCount(passports))
}

func buildPassportRows() []string {
	rows := file_utils.OpenFileIntoSlice("./challenge4/input.txt")
	formattedRows := make([]string, 0)
	inRow := false
	acc := ""
	for _, row := range rows {
		if len(row) < 1 && len(acc) > 0 {
			formattedRows = append(formattedRows, acc)
			acc = ""
			inRow = false
		} else if inRow {
			acc += (" " + row)
		} else if len(row) > 0 {
			acc += row
			inRow = true

		}
	}
	return formattedRows
}

func partTwoValidPassportCount(passports []Passport) int {
	validCount := 0
	for _, passport := range passports {
		if passport.RequiredFieldsValid() {
			validCount++
		}

	}
	return validCount
}

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p Passport) RequiredFieldsSet() bool {
	valid := true
	for _, field := range []string{p.byr, p.iyr, p.eyr, p.hgt, p.hcl, p.ecl, p.pid} {
		if len(field) < 1 {
			valid = false
			break
		}
	}
	return valid
}

type validatorFn func() bool

func (p Passport) RequiredFieldsValid() bool {
	valid := true
	for _, fieldValid := range p.validatorFns() {
		if !fieldValid() {
			valid = false
			break
		}
	}
	return valid
}

func (p Passport) validatorFns() []validatorFn {
	return []validatorFn{
		p.validByr,
		p.validIyr,
		p.validEyr,
		p.validHgt,
		p.validHcl,
		p.validEcl,
		p.validPid,
	}
}

func (p Passport) validByr() bool {
	return validYearRange(p.byr, 1920, 2002)
}

func (p Passport) validIyr() bool {
	return validYearRange(p.iyr, 2010, 2020)
}

func (p Passport) validEyr() bool {
	return validYearRange(p.eyr, 2020, 2030)

}

func validYearRange(field string, lowerBound, upperBound int) bool {
	if len(field) != 4 {
		return false
	}
	intField, err := strconv.Atoi(field)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to convert byr to int %s", field))
	}
	return intField >= lowerBound && intField <= upperBound
}

func (p Passport) validHgt() bool {
	var (
		height int
		unit   string
		valid  bool
	)
	valid = false
	if !(strings.Contains(p.hgt, "cm") || strings.Contains(p.hgt, "in")) {
		return valid
	}
	_, err := fmt.Sscanf(p.hgt, "%d%s", &height, &unit)
	if err != nil {
		log.Fatal(err)
	}
	if unit == "in" {
		valid = height >= 59 && height <= 76
	} else if unit == "cm" {
		valid = height >= 150 && height <= 193
	} else {
		valid = false
	}

	return valid
}

func (p Passport) validHcl() bool {
	if len(p.hcl) != 7 {
		return false
	}
	re := regexp.MustCompile("^#(\\d|[a-f]){6}$")
	return re.Match([]byte(p.hcl))
}

func (p Passport) validEcl() bool {
	valid := false
	for _, color := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if p.ecl == color {
			valid = true
			break
		}
	}
	return valid
}

func (p Passport) validPid() bool {
	if len(p.pid) != 9 {
		return false
	}
	re := regexp.MustCompile("^\\d{9}$")

	return re.Match([]byte(p.pid))
}

func newPassport(row string) Passport {
	fields := strings.Split(row, " ")
	p := Passport{}
	for _, field := range fields {
		kv := strings.Split(field, ":")
		key, value := kv[0], kv[1]
		switch key {
		case "byr":
			p.byr = value
		case "iyr":
			p.iyr = value
		case "eyr":
			p.eyr = value
		case "hgt":
			p.hgt = value
		case "hcl":
			p.hcl = value
		case "ecl":
			p.ecl = value
		case "pid":
			p.pid = value
		case "cid":
			p.cid = value
		default:
			log.Fatal(fmt.Sprintf("Field %s with value %s not expected!\n", key, value))
		}
	}

	if p.RequiredFieldsSet() {
		return p
	}
	return Passport{}
}

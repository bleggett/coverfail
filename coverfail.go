package main

import (
	"bytes"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"os"
	"os/exec"
	"regexp"
)

const usageMessage = "" +
	`Usage:	coverfail -threshold X.X`

var (
	threshold    float64
	coverprofile string
)

func init() {
	flag.Float64Var(&threshold, "threshold", 0, "Sets the threshold the actual coverage number will be compared against")
	flag.StringVar(&coverprofile, "coverprofile", "coverage.out", "Write a coverage profile to the file after all tests have passed")
}

func usage() {
	fmt.Fprintln(os.Stderr, usageMessage)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
	os.Exit(2)
}

type ExitError struct {
	Msg  string
	Code int
}

func (e *ExitError) Error() string {
	return e.Msg
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if err := run(coverprofile, threshold); err != nil {
		code := 1
		if err, ok := err.(*ExitError); ok {
			code = err.Code
		}
		if err.Error() != "" {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(code)
	}
}

func run(coverprofile string, threshold float64) error {
	optionalArgs := buildOptionalTestArgs(coverprofile)
	err := coverage(optionalArgs, threshold)
	if err != nil {
		return err
	}
	return nil
}

func buildOptionalTestArgs(coverprofile string) []string {
	args := []string{}
	if coverprofile != "" {
		args = append(args, "-coverprofile=", coverprofile)
	}
	return args
}

func coverage(optArgs []string, threshold float64) error {
	args := append([]string{"test", "-cover", "-coverpkg=./...", "./..."}, optArgs...)
	cmd := exec.Command("go", args...)
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprint(os.Stdout, stdout.String())
		fmt.Fprint(os.Stderr, stderr.String())
		return &ExitError{Code: 1, Msg: "'go test' exited with an error, no coverage results available"}
	}
	totalpct := parsePackagePercentages(stdout)
	fmt.Println("Overall code coverage percentage is: ", totalpct)
	fmt.Println("Threshold is: ", threshold)
	if totalpct < threshold {
		return &ExitError{Code: 1, Msg: "Overall coverage is lower than provided threshold number, bailing with nonzero exit code..."}
	}
	return nil
}


func parsePackagePercentages(output *bytes.Buffer) float64 {
	var total float64

	pctMatch := regexp.MustCompile(`([\d*\.?\d*]+)(%)`)
	outputStr := output.String()
	percents := pctMatch.FindAllString(outputStr, -1)
	for _, pct:= range percents {
		coveragePct, err := strconv.ParseFloat(strings.Trim(pct, "%"), 64); if err != nil {
			panic(fmt.Sprintf("Could not parse code coverage output, error was: %s", err))
		}
		total += coveragePct
	}
	return total
}

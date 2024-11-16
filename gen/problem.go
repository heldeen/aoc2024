package gen

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"time"
)

const (
	glueTemplate = `package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
	
	"github.com/heldeen/aoc2024/challenge"
	"github.com/heldeen/aoc2024/challenge/day{{ .N }}"
)

func init() {
	const inputFlag = "input"
	const inputFlagShort = "i"
	const inputFlagUsage = "path of the input file to use"
	
	var inputFlagValue string
	
    day := &cobra.Command{
		Use:   "{{ .N }}",
		Short: "Problems for Day {{ .N }}",
	}

	a := &cobra.Command{
		Use:   "a",
		Short: "Day {{ .N }}, Problem A",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Flag("input")
			input, err := challenge.FromFileP(flag.Value.String())
			cobra.CheckErr(err)
			fmt.Printf("Day {{ .N }}, Part A - Answer: %v\n", day{{ .N }}.PartA(input))
		},
	}

	a.Flags().StringVarP(&inputFlagValue, inputFlag, inputFlagShort, "./challenge/day1/input.txt", inputFlagUsage)

	day.AddCommand(a)

	b := &cobra.Command{
		Use:   "b",
		Short: "Day {{ .N }}, Part B",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Flag("input")
            input, err := challenge.FromFileP(flag.Value.String())
			cobra.CheckErr(err)
			fmt.Printf("Day {{ .N }}, Part B - Answer: %v\n", day{{ .N }}.PartB(input))
		},
	}

	b.Flags().StringVarP(&inputFlagValue, inputFlag, inputFlagShort, "./challenge/day1/input.txt", inputFlagUsage)

	day.AddCommand(b)

	rootCmd.AddCommand(day)
}
`
	problemTemplate = `package day{{ .N }}

import (
    "github.com/heldeen/aoc2024/challenge"
)

func Part{{ .AB }}(challenge *challenge.Input) int {
    return 0
}
`

	testTemplate = `package day{{ .N }}
		
import (
	"testing"

    "github.com/heldeen/aoc2024/challenge"
)

{{if eq .AB "A"}}` + "const sample = ``" + `
{{end}} 
func Test{{ .AB }}(t *testing.T) {
	want := 42 

    input := challenge.FromLiteral(sample)

	result := Part{{ .AB }}(input)

	if result != want {
	  t.Errorf("Day[{{ .N }}] Part[{{ .AB }}] - wanted [%d] but got [%d]", want, result)
    }
}
`
	ideaRunConfigTemplate = `<component name="ProjectRunConfigurationManager">
  <configuration default="false" name="aoc2024 Day {{ .N }} Part {{ .AB }}" type="GoApplicationRunConfiguration" factoryName="Go Application" nameIsGenerated="true">
    <module name="aoc2024" />
    <working_directory value="$PROJECT_DIR$" />
    <parameters value="{{ .N }} {{toLower .AB }}" />
    <kind value="PACKAGE" />
    <package value="github.com/heldeen/aoc2024" />
    <directory value="$PROJECT_DIR$" />
    <filePath value="$PROJECT_DIR$/main.go" />
    <method v="2" />
  </configuration>
</component>
`

	ideaTestRunConfigTemplate = `<component name="ProjectRunConfigurationManager">
  <configuration default="false" name="Test{{ .AB }} in aoc2024/challenge/day{{ .N }}" type="GoTestRunConfiguration" factoryName="Go Test" nameIsGenerated="true">
    <module name="aoc2024" />
    <working_directory value="$PROJECT_DIR$/challenge/day{{ .N }}" />
    <root_directory value="$PROJECT_DIR$" />
    <kind value="PACKAGE" />
    <package value="github.com/heldeen/aoc2024/challenge/day{{ .N }}" />
    <directory value="$PROJECT_DIR$" />
    <filePath value="$PROJECT_DIR$" />
    <framework value="gotest" />
    <pattern value="^\QTest{{ .AB }}\E$" />
    <method v="2" />
  </configuration>
</component>`
)

type metadata struct {
	N  int
	AB string
}

func GenerateDay(day int) error {

	cmdPath, err := cmdPkgPath()
	if err != nil {
		return err
	}

	probPath, err := pkgPath(day)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(probPath, 0744); err != nil {
		return err
	}

	funcs := template.FuncMap{
		"toLower": strings.ToLower,
	}

	gluePath := filepath.Join(cmdPath, fmt.Sprintf("importDay%d.go", day))
	if _, stat := os.Stat(gluePath); stat != nil && os.IsNotExist(stat) {
		genFile(gluePath, glueTemplate, funcs, metadata{N: day, AB: "A"})
	}

	projP, err := projectPath()
	if err != nil {
		return err
	}

	for _, ab := range []string{"A", "B"} {
		m := metadata{
			N:  day,
			AB: ab,
		}
		genFile(filepath.Join(probPath, fmt.Sprintf("%s.go", strings.ToLower(ab))), problemTemplate, funcs, m)
		genFile(filepath.Join(probPath, fmt.Sprintf("%s_test.go", strings.ToLower(ab))), testTemplate, funcs, m)
		genFile(filepath.Join(projP, ".idea/runConfigurations/", fmt.Sprintf("Run_aoc2024_challenge_day%d_part%s.xml", day, ab)), ideaRunConfigTemplate, funcs, m)
		genFile(filepath.Join(projP, ".idea/runConfigurations/", fmt.Sprintf("Test%s_in_aoc2024_challenge_day%d.xml", ab, day)), ideaTestRunConfigTemplate, funcs, m)
	}
	//
	//goimports := exec.Command("goimports", "-w", probPath)
	//if err := goimports.Run(); err != nil {
	//	return err
	//}

	inputOutputPath := filepath.Join(probPath, "input.txt")
	if _, stat := os.Stat(inputOutputPath); os.IsNotExist(stat) {
		log.Printf("fetching input for day...%d\n", day)
		problemInput, err := getInput(day)
		if err != nil {
			return err
		}

		if err := os.WriteFile(inputOutputPath, problemInput, 0644); err != nil {
			return err
		}

	} else {
		log.Print("input already downloaded, skipping...")
	}

	log.Printf("Generated problems for day %d.", day)

	return nil
}

func projectPath() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to generate package path")
	}

	return filepath.Dir(filepath.Dir(filename)), nil
}

func pkgPath(day int) (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to generate package path")
	}

	return filepath.Join(filepath.Dir(filepath.Dir(filename)), "challenge", fmt.Sprintf("day%d", day)), nil
}

func cmdPkgPath() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to generate package path")
	}

	return filepath.Join(filepath.Dir(filepath.Dir(filename)), "cmd"), nil
}

func genFile(path, t string, funcs template.FuncMap, m metadata) {
	log.Println("creating", path)
	if _, stat := os.Stat(path); os.IsNotExist(stat) {
		t := template.Must(template.New(path).Funcs(funcs).Parse(t))
		cf, err := os.Create(path)
		if err != nil {
			log.Fatalf("creating path %v", err)
		}

		defer mustClose(cf)
		if err := t.Execute(cf, m); err != nil {
			log.Fatalf("exectute template %v", err)
		}
	} else {
		log.Println(path, "already exists, skipping...")
	}
}

func getInput(day int) ([]byte, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(homeDir, ".tokenfile")
	if _, stat := os.Stat(path); os.IsNotExist(stat) {
		log.Println(path, "not found, skipping...")
		return nil, nil
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	exp, err := time.Parse(time.RFC3339, "2025-12-21T21:17:39.211Z")
	if err != nil {
		return nil, err
	}

	ck := http.Cookie{
		Name:       "session",
		Value:      string(b),
		Path:       "/",
		Domain:     ".adventofcode.com",
		Expires:    exp,
		RawExpires: "",
		MaxAge:     0,
		Secure:     true,
		HttpOnly:   true,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day), nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&ck)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer mustClose(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code %s: %s", resp.Status, body)
	}

	body, err := io.ReadAll(resp.Body)
	return body, err
}

func mustClose(c io.Closer) {
	if c == nil {
		return
	}

	if err := c.Close(); err != nil {
		log.Fatalf("error closing io.Closer: %v", err)
	}
}

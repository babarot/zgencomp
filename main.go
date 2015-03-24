package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"
	"time"

	flags "github.com/jessevdk/go-flags"
)

const (
	JSON = "data/templates/sample.json"
	TPL  = "data/templates/sample.tpl"
)

// type JsonData struct {{{1
type JsonData struct {
	Command    string     `json:"command"`
	Properties Properties `json:"properties"`
	Options    Options    `json:"options"`
	Arguments  Arguments  `json:"arguments"`
}

type Properties struct {
	Author  string  `json:"author"`
	License string  `json:"license"`
	Help    Help    `json:"help"`
	Version Version `json:"version"`
}

type Help struct {
	Option      []string `json:"option"`
	Description string   `json:"description"`
}

type Version struct {
	Option      []string `json:"option"`
	Description string   `json:"description"`
}

type Options struct {
	Switch []Switch `json:"switch"`
	Flag   []Flag   `json:"flag"`
}

type Switch struct {
	Option      []string `json:"option"`
	Description string   `json:"description"`
	Exclusion   []string `json:"exclusion"`
}

type Flag struct {
	Option      []string `json:"option"`
	Description string   `json:"description"`
	Exclusion   []string `json:"exclusion"`
	Argument    Argument `json:"argument"`
}

type Argument struct {
	Group string              `json:"group"`
	Type  interface{}         `json:"type"`
	Style map[string][]string `json:"style"`
}

type Arguments struct {
	Always    bool        `json:"always"`
	After_arg bool        `json:"after_arg"`
	Type      interface{} `json:"type"`
}

//}}}

// Library {{{1
func isExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//}}}

// getter function {{{1
func readJson(f string) (jd JsonData, err error) {
	// Read json file if exists
	file, err := os.Open(f)
	if err != nil {
		return
	}
	defer file.Close()

	var contents string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if m, _ := regexp.MatchString("^ *\\/\\/", scanner.Text()); !m {
			if m, _ := regexp.MatchString("^ *\\/\\*", scanner.Text()); !m {
				if m, _ := regexp.MatchString("^ *\\*\\/", scanner.Text()); !m {
					contents += scanner.Text() + "\n"
				}
			}
		}
	}
	err = scanner.Err()
	if err != nil {
		return
	}

	// Deal with json file
	//decoder := json.NewDecoder(file)
	decoder := json.NewDecoder(strings.NewReader(contents))
	err = decoder.Decode(&jd)
	if err != nil {
		err = fmt.Errorf("[json file broken]: %s", err)
		return
	}

	// In case of json file broken
	if jd.Command == "" {
		err = errors.New("[json file broken]: set command name")
	}

	return jd, err
}

func (jd JsonData) jsonOutput(stream io.Writer) (err error) {
	funcMap := template.FuncMap{
		"dealWithOption":            dealWithOption,
		"dealWithSwitchOption":      dealWithSwitchOption,
		"dealWithFlagOption":        dealWithFlagOption,
		"dealWithExclusion":         dealWithExclusion,
		"dealWithSwitchExclusion":   dealWithSwitchExclusion,
		"dealWithFlagExclusion":     dealWithFlagExclusion,
		"dealWithDescription":       dealWithDescription,
		"dealWithFlagArgumentStyle": dealWithFlagArgumentStyle,
		"whetherOptionIsEnabled":    whetherOptionIsEnabled,
		"whetherTypeIsFunc":         whetherTypeIsFunc,
		"setFlagMessage":            setFlagMessage,
		"setAction":                 setAction,
		"helperTrimArrowInType":     helperTrimArrowInType,
		"dateYear":                  func() int { now := time.Now(); return now.Year() },
	}

	// Replace ioutil with Asset
	//contents, err := ioutil.ReadFile("./sample.tpl")
	contents, err := Asset(TPL)
	if err != nil {
		//log.Fatal(err)
		return
	}

	tpl := template.Must(template.New("scan").Funcs(funcMap).Parse(string(contents)))
	data := JsonData{
		jd.Command,
		jd.Properties,
		jd.Options,
		jd.Arguments,
	}

	err = tpl.Execute(stream, data)
	if err != nil {
		//log.Fatal(err)
		return
	}

	return
}

func generateSampleJson(f string) (err error) {
	data, err := Asset(JSON)
	if err != nil {
		return
	}

	if isExists(f) {
		fmt.Fprintf(os.Stderr, "%s is already exists, overwrite it? [y/N]: ", f)
		var ans string
		_, err = fmt.Scanf("%s", &ans)
		if err != nil {
			return
		}

		if strings.ToLower(ans) == "y" {
			err = os.RemoveAll(f)
			if err != nil {
				return
			}
		} else {
			os.Exit(0)
		}
	}

	ioutil.WriteFile(f, data, os.ModePerm)
	return
}

//}}}

type CLI struct {
	Help     bool   `short:"h" long:"help" description:"Show this help and exit"`
	Version  bool   `long:"version" description:"Show version"`
	File     bool   `short:"f" long:"file" description:"Write to file instead of stdout"`
	Generate string `short:"g" long:"genrate" description:"Generate" optional:"yes" optional-value:"sample.json"`
}

const version = "0.2"

var cli CLI

func main() {
	parser := flags.NewParser(&cli, flags.Default)
	parser.Name = "zgencomp"
	parser.Usage = "[OPTION]"

	args, _ := parser.Parse()

	// --version
	// deal with version
	if cli.Version {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	// -g, --genrate
	// deal with generating
	if cli.Generate != "" {
		if err := generateSampleJson(cli.Generate); err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}

	// Reading json file
	f := path.Base(JSON)
	if len(args) != 0 {
		f = args[0]
	}
	jd, err := readJson(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// deal with output json
	out := os.Stdout
	if cli.File {
		f, err := os.Create("_" + jd.Command)
		if err != nil {
			os.Exit(1)
		}
		out = f
	}

	// Analyzing json and output
	err = jd.jsonOutput(out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// dealWithOption {{{
func dealWithOption(s []string) (ret string) {
	if len(s) == 1 {
		ret = s[0]
	} else {
		ret = "'{" + strings.Join(s, ",") + "}'"
	}

	return
}

//}}}

// dealWithSwitchOption {{{
func dealWithSwitchOption(s Switch) (ret string) {
	ret = dealWithOption(s.Option)
	return
}

//}}}

// dealWithFlagOption {{{
func dealWithFlagOption(s Flag) (ret string) {
	ret = dealWithOption(dealWithFlagArgumentStyle(s))
	return
}

//}}}

// dealWithExclusion {{{
func dealWithExclusion(s []string) (ret string) {
	return "(" + strings.Join(s, " ") + ")"
}

//}}}

// dealWithSwitchExclusion {{{
func dealWithSwitchExclusion(s Switch) (ret string) {
	if len(s.Exclusion) == 0 {
		ret = strings.Join(s.Option, " ")
		ret = "(" + ret + ")"
		return
	}

	var exclusion string
	for _, e := range s.Option {
		if stringInSlice(e, s.Exclusion) {
			continue
		}
		exclusion = exclusion + " " + e
	}
	for _, e := range s.Exclusion {
		if stringInSlice(e, s.Option) {
			continue
		}
		exclusion = exclusion + " " + e
	}
	exclusion = strings.TrimSpace(exclusion)
	ret = "(" + exclusion + ")"

	return
}

//}}}

// dealWithFlagExclusion {{{
func dealWithFlagExclusion(s Flag) (ret string) {
	if len(s.Exclusion) == 0 {
		ret = strings.Join(s.Option, " ")
		ret = "(" + ret + ")"
		return
	}

	var exclusion string
	for _, e := range s.Option {
		if stringInSlice(e, s.Exclusion) {
			continue
		}
		exclusion = exclusion + " " + e
	}
	for _, e := range s.Exclusion {
		if stringInSlice(e, s.Option) {
			continue
		}
		exclusion = exclusion + " " + e
	}
	exclusion = strings.TrimSpace(exclusion)
	ret = "(" + exclusion + ")"

	return
}

//}}}

// dealWithDescription {{{
func dealWithDescription(s string) (ret string) {
	ret = strings.Replace(s, "'", "'\\''", -1)
	return
}

//}}}

// helperTrimArrowInType {{{
func helperTrimArrowInType(s string) string {
	return strings.TrimLeft(s, "->")
}

//}}}

// whetherOptionIsEnabled {{{
// check whether there is one or more options
func whetherOptionIsEnabled(s []string) (ret bool) {
	ret = false
	if len(s) != 0 {
		ret = true
	}

	return
}

//}}}

// whetherTypeIsFunc {{{
// check whether Options.Flag.Argument.Type or Arguments.Type is "func"
func whetherTypeIsFunc(s interface{}) (ret bool) {
	switch s.(type) {
	case Flag:
		s = s.(Flag).Argument.Type
	}
	switch s.(type) {
	case string:
		ret = s.(string) == "func"
	default:
		ret = false
	}
	return
}

//}}}

// setFlagMessage {{{
func setFlagMessage(s Flag) (ret string) {
	ret = s.Argument.Group
	if ret == "" {
		ret = " "
	}

	return
}

//}}}

// dealWithFlagArgumentStyle {{{
func dealWithFlagArgumentStyle(s Flag) (ret []string) {
	//retSlice := s.Option
	retMap := make(map[string][]string)

	for k, v := range s.Argument.Style {
		// Skip blank value
		// map[equal:[--after-context] equalable:[] standard:[-A] touch:[-a] touchable:[]]
		// ==> map[standard:[-A] touch:[-a] equal:[--after-context]]
		if len(v) == 0 {
			continue
		}

		// Skip invalid value
		// map[standard:[-A] touch:[-a] equal:[--after-context]]
		// ==> map[standard:[-A] equal:[--after-context]]
		for _, e := range v {
			if stringInSlice(e, s.Option) {
				retMap[k] = v
			}
		}
	}

	var retSlice []string
	for _, e := range s.Option {
		retSlice = append(retSlice, helperAddFlagArgumentStyle(retMap, e))
	}
	return retSlice
}

//}}}

// helperAddFlagArgumentStyle {{{
func helperAddFlagArgumentStyle(m map[string][]string, s string) (ret string) {
	for k, v := range m {
		switch k {
		case "standard":
			for _, e := range v {
				if e == s {
					ret = strings.Replace(e, e, e, -1)
					return
				}
			}
		case "touch":
			for _, e := range v {
				if e == s {
					ret = strings.Replace(e, e, e+"-", -1)
					return
				}
			}
		case "touchable":
			for _, e := range v {
				if e == s {
					ret = strings.Replace(e, e, e+"+", -1)
					return
				}
			}
		case "equal":
			for _, e := range v {
				if e == s {
					ret = strings.Replace(e, e, e+"=-", -1)
					return
				}
			}
		case "equalable":
			for _, e := range v {
				if e == s {
					ret = strings.Replace(e, e, e+"=", -1)
					return
				}
			}
		default:
		}
	}
	ret = s
	return
}

//}}}

// setAction {{{
// set action of completion
func setAction(s interface{}) (ret string) {
	backup := s
	isFlag := false

	switch s.(type) {
	case Flag:
		s = s.(Flag).Argument.Type
		isFlag = true
	}

	switch s.(type) {
	case string:
		// assume "func" and so on
		ret = s.(string)
		switch ret {
		case "func":
			var opt string
			if isFlag {
				opt = backup.(Flag).Option[0]
				re, _ := regexp.Compile("^(--?|\\+)")
				opt = re.ReplaceAllString(opt, "")
				ret = opt + "_func"
			} else {
				ret = "args"
			}
			ret = "->" + ret
		case "file":
			ret = "_files"
		case "dir", "directory":
			ret = "_files -/"
		default:
			ret = " "
		}

	case []string:
		// assume "(word1 word2...)"
		ret = strings.Join(s.([]string), " ")
		ret = "(" + ret + ")"
	case []interface{}:
		// assume "(word1 word2...)"
		for _, v := range s.([]interface{}) {
			ret = ret + " " + v.(string)
		}
		ret = "(" + strings.TrimSpace(ret) + ")"

	case map[string]string:
		// assume "((word1\:desc1 word2\:desc2...))"
		for k, v := range s.(map[string]string) {
			ret = ret + k + "\\:" + "\"" + v + "\"" + " "
		}
		ret = "((" + strings.TrimSpace(ret) + "))"

	case map[string]interface{}:
		for k, v := range s.(map[string]interface{}) {
			ret = ret + k + "\\:" + "\"" + v.(string) + "\"" + " "
		}
		ret = "((" + strings.TrimSpace(ret) + "))"

	default:
		ret = "D A N"
	}

	return
}

//}}}

//}}}

// vim: fdm=marker fdc=3

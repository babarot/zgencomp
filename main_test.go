package main

import (
	"testing"
)

// readJson
func TestReadJson(t *testing.T) {
}

// jsonOutput
func TestJsonOutput(t *testing.T) {
}

// generateSampleJson
func TestGenerateSampleJson(t *testing.T) {
}

// main
func TestMain(t *testing.T) {
}

//

// dealWithOption {{{
func TestDealWithOptionIfArgument(t *testing.T) {
	actual := dealWithOption([]string{"-a"})
	expected := "-a"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestDealWithOptionIfArguments(t *testing.T) {
	actual := dealWithOption([]string{"-a", "--all"})
	expected := "'{-a,--all}'"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//}}}

// dealWithSwitchOption {{{
func TestDealWithSwitchOptionIfArgument(t *testing.T) {
	var s Switch = Switch{
		Option:      []string{"-a"},
		Description: "",
		Exclusion:   []string{},
	}

	actual := dealWithSwitchOption(s)
	expected := "-a"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestDealWithSwitchOptionIfArguments(t *testing.T) {
	var s Switch = Switch{
		Option:      []string{"-a", "--all"},
		Description: "",
		Exclusion:   []string{},
	}

	actual := dealWithSwitchOption(s)
	expected := "'{-a,--all}'"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//}}}

// dealWithFlagOption {{{
func TestDealWithFlagOption(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{"-a", "--all"},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  nil,
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{},
				"equalable": []string{},
			},
		},
	}

	actual := dealWithFlagOption(f)
	expected := "'{-a,--all}'"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestDealWithFlagOptionEqual(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{"-a", "--all"},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  nil,
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{"--all"},
				"equalable": []string{},
			},
		},
	}

	actual := dealWithFlagOption(f)
	expected := "'{-a,--all=-}'"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//}}}

// dealWithExclusion {{{
func TestDealWithExclusion(t *testing.T) {
	actual := dealWithExclusion([]string{"-a", "--all"})
	expected := "(-a --all)"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//}}}

// dealWithSwitchExclusion {{{
func TestDealWithSwitchExclusionNoExclusion(t *testing.T) {
	var s Switch = Switch{
		Option:      []string{"-opt"},
		Description: "",
		Exclusion:   []string{},
	}

	actual := dealWithSwitchExclusion(s)
	expected := "(-opt)"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestDealWithSwitchExclusionExclusion(t *testing.T) {
	var s Switch = Switch{
		Option:      []string{"-opt"},
		Description: "",
		Exclusion:   []string{"--other"},
	}

	actual := dealWithSwitchExclusion(s)
	expected := "(-opt --other)"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestDealWithSwitchExclusionToggle(t *testing.T) {
	var s Switch = Switch{
		Option:      []string{"-opt"},
		Description: "",
		Exclusion:   []string{"--other", "-opt"},
	}

	actual := dealWithSwitchExclusion(s)
	expected := "(--other)"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//}}}

// dealWithFlagExclusion {{{
func TestDealWithFlagExclusionNoExclusion(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{"-opt"},
		Description: "",
		Exclusion:   []string{},
	}

	actual := dealWithFlagExclusion(f)
	expected := "(-opt)"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestDealWithFlagExclusionExclusion(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{"-opt"},
		Description: "",
		Exclusion:   []string{"--other"},
	}

	actual := dealWithFlagExclusion(f)
	expected := "(-opt --other)"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestDealWithFlagExclusionToggle(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{"-opt"},
		Description: "",
		Exclusion:   []string{"--other", "-opt"},
	}

	actual := dealWithFlagExclusion(f)
	expected := "(--other)"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//}}}

// dealWithDescription {{{
func TestDealWithDescription(t *testing.T) {
	actual := dealWithDescription("description")
	expected := "description"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestDealWithDescriptionIncludeSingleQuotation(t *testing.T) {
	actual := dealWithDescription("I'm lovin' it")
	expected := "I'\\''m lovin'\\'' it"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//}}}

// dealWithFlagArgumentStyle {{{
func helperEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestDealWithFlagArgumentStyle(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{"-a", "--all"},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  nil,
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{"--all"},
				"equalable": []string{},
			},
		},
	}

	actual := dealWithFlagArgumentStyle(f)
	expected := []string{"-a", "--all=-"}

	//if !reflect.DeepEqual(actual, expected) {
	if !helperEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestDealWithFlagArgumentStyle2(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{"-a", "--all"},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  nil,
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{"-a"},
				"equal":     []string{"--all"},
				"equalable": []string{},
			},
		},
	}

	actual := dealWithFlagArgumentStyle(f)
	expected := []string{"-a+", "--all=-"}

	//if !reflect.DeepEqual(actual, expected) {
	if !helperEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//}}}

// whetherOptionIsEnabled {{{
func TestWhetherOptionIsEnabledFalse(t *testing.T) {
	actual := whetherOptionIsEnabled([]string{})
	expected := false

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestWhetherOptionIsEnabledTrue(t *testing.T) {
	actual := whetherOptionIsEnabled([]string{"-a"})
	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//}}}

// whetherTypeIsFunc {{{
func TestWhetherTypeIsFunc(t *testing.T) {
	actual := whetherTypeIsFunc("func")
	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestWhetherTypeIsFuncFlag(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  "func",
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{},
				"equalable": []string{},
			},
		},
	}

	actual := whetherTypeIsFunc(f)
	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestWhetherTypeIsNotFunc(t *testing.T) {
	actual := whetherTypeIsFunc("file")
	expected := false

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestWhetherTypeIsNotString(t *testing.T) {
	actual := whetherTypeIsFunc([]string{})
	expected := false

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//}}}

// setFlagMessage {{{
func TestSetFlagMessage(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "akb",
			Type:  "",
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{},
				"equalable": []string{},
			},
		},
	}

	actual := setFlagMessage(f)
	expected := "akb"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSetFlagMessageBlank(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  "",
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{},
				"equalable": []string{},
			},
		},
	}

	actual := setFlagMessage(f)
	expected := " "

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//}}}

// setArgument {{{
func TestSetArgumentIfStringFile(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  "file",
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{},
				"equalable": []string{},
			},
		},
	}

	actual := setAction(f)
	expected := "_files"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSetArgumentIfStringDir(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  "dir",
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{},
				"equalable": []string{},
			},
		},
	}

	actual := setAction(f)
	expected := "_files -/"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSetArgumentIfStringFuncIfOptionIsBlank(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  "func",
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{},
				"equalable": []string{},
			},
		},
	}

	actual := setAction(f)
	expected := "func"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSetArgumentIfStringFuncIfFlag(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{"--all"},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  "func",
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{},
				"equalable": []string{},
			},
		},
	}

	actual := setAction(f)
	expected := "->all_func"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSetArgumentIfStringFuncIfNonFlag(t *testing.T) {
	var f Arguments = Arguments{
		Always: false,
		Type:   "func",
	}

	actual := setAction(f.Type)
	expected := "->args"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSetArgumentIfStringIsBlank(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  "",
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{},
				"equalable": []string{},
			},
		},
	}

	actual := setAction(f)
	expected := " "

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSetArgumentIfSlice(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  []string{"word1", "word2"},
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{},
				"equalable": []string{},
			},
		},
	}

	actual := setAction(f)
	expected := "(word1 word2)"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSetArgumentIfSliceInterface(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  []interface{}{"word1", "word2"},
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{},
				"equalable": []string{},
			},
		},
	}

	actual := setAction(f)
	expected := "(word1 word2)"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSetArgumentIfMap(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  map[string]string{"word1": "desc1", "word2": "desc2"},
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{},
				"equalable": []string{},
			},
		},
	}

	actual := setAction(f)
	expected1 := "((word1\\:\"desc1\" word2\\:\"desc2\"))"
	expected2 := "((word2\\:\"desc2\" word1\\:\"desc1\"))"

	if actual != expected1 && actual != expected2 {
		t.Errorf("got %v\nwant %v or %v", actual, expected1, expected2)
	}
}

func TestSetArgumentIfMapInterface(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  map[string]interface{}{"word1": "desc1", "word2": "desc2"},
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{},
				"equalable": []string{},
			},
		},
	}

	actual := setAction(f)
	expected1 := "((word1\\:\"desc1\" word2\\:\"desc2\"))"
	expected2 := "((word2\\:\"desc2\" word1\\:\"desc1\"))"

	if actual != expected1 && actual != expected2 {
		t.Errorf("got %v\nwant %v or %v", actual, expected1, expected2)
	}
}

func TestSetArgumentIfInvalid(t *testing.T) {
	var f Flag = Flag{
		Option:      []string{},
		Description: "",
		Exclusion:   []string{},
		Argument: Argument{
			Group: "",
			Type:  1,
			Style: map[string][]string{
				"standard":  []string{},
				"touch":     []string{},
				"touchable": []string{},
				"equal":     []string{},
				"equalable": []string{},
			},
		},
	}

	actual := setAction(f)
	expected := "[[Parse Error]]"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//}}}

// helperTrimArrowInType {{{
func TestHelperTrimArrowInType(t *testing.T) {
	actual := helperTrimArrowInType("->hello")
	expected := "hello"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestHelperTrimArrowInTypeReturn(t *testing.T) {
	actual := helperTrimArrowInType("he->llo")
	expected := "he->llo"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//}}}

// helperAddFlagArgumentStyle {{{
func TestHelperAddFlagArgumentStyleIfStandard(t *testing.T) {
	m := map[string][]string{
		"standard":  []string{"-opt"},
		"touch":     []string{},
		"touchable": []string{},
		"equal":     []string{},
		"equalable": []string{},
	}

	actual := helperAddFlagArgumentStyle(m, "-opt")
	expected := "-opt"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestHelperAddFlagArgumentStyleIfTouch(t *testing.T) {
	m := map[string][]string{
		"standard":  []string{},
		"touch":     []string{"-opt"},
		"touchable": []string{},
		"equal":     []string{},
		"equalable": []string{},
	}

	actual := helperAddFlagArgumentStyle(m, "-opt")
	expected := "-opt-"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestHelperAddFlagArgumentStyleIfTouchable(t *testing.T) {
	m := map[string][]string{
		"standard":  []string{},
		"touch":     []string{},
		"touchable": []string{"-opt"},
		"equal":     []string{},
		"equalable": []string{},
	}

	actual := helperAddFlagArgumentStyle(m, "-opt")
	expected := "-opt+"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestHelperAddFlagArgumentStyleIfEqual(t *testing.T) {
	m := map[string][]string{
		"standard":  []string{},
		"touch":     []string{},
		"touchable": []string{},
		"equal":     []string{"-opt"},
		"equalable": []string{},
	}

	actual := helperAddFlagArgumentStyle(m, "-opt")
	expected := "-opt=-"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestHelperAddFlagArgumentStyleIfEqualable(t *testing.T) {
	m := map[string][]string{
		"standard":  []string{},
		"touch":     []string{},
		"touchable": []string{},
		"equal":     []string{},
		"equalable": []string{"-opt"},
	}

	actual := helperAddFlagArgumentStyle(m, "-opt")
	expected := "-opt="

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestHelperAddFlagArgumentStyleReturnAsItIs(t *testing.T) {
	m := map[string][]string{
		"false": []string{""},
	}

	actual := helperAddFlagArgumentStyle(m, "--true")
	expected := "--true"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//func TestHelperAddFlagArgumentStyleParseErrorReturnBlank(t *testing.T) {
//	m := map[string][]string{
//		"standard": []string{""},
//	}
//
//	actual := helperAddFlagArgumentStyle(m, "-false")
//	expected := ""
//
//	if actual != expected {
//		t.Errorf("got %v\nwant %v", actual, expected)
//	}
//}

func TestHelperAddFlagArgumentStyleGetRidOf(t *testing.T) {
	m := map[string][]string{
		"standard":  []string{},
		"touch":     []string{"--true"},
		"touchable": []string{},
		"equal":     []string{"--false"},
		"equalable": []string{},
	}

	actual := helperAddFlagArgumentStyle(m, "--true")
	expected := "--true-"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//}}}

// vim: fdm=marker fdc=3

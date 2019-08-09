package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

var completionsRaw = flag.String("completions", "", "arg1_name:arg1_DESC|||arg2_name:arg2_DESC")
var defaultCompletion = flag.String("default", "", "defaults to first completion")
var title = flag.String("title", "", "")

type completions []prompt.Suggest

func main() {
	flag.Parse()
	if *completionsRaw == "" {
		flag.Usage()
		os.Exit(1)
	}
	var suggests []prompt.Suggest
	for _, c := range strings.Split(*completionsRaw, "|||") {
		r := strings.SplitN(c, ":", 2)
		suggests = append(suggests, prompt.Suggest{Text: r[0], Description: r[1]})
	}

	completer := func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(suggests, d.GetWordBeforeCursor(), true)
	}

	if *defaultCompletion == "" {
		defaultCompletion = &suggests[0].Text
	}

	var choices []string
	for _, s := range suggests {
		choices = append(choices, s.Text)
	}

	if *title == "" {
		fmt.Printf("(%s)\n", strings.Join(choices, "|"))
	} else {
		fmt.Printf("%s (%s)\n", *title, strings.Join(choices, "|"))
	}

	t := prompt.Input("> ", completer, prompt.OptionWriter(prompt.NewStderrWriter()), prompt.OptionShowCompletionAtStart())
	var output string
	if t == "" {
		output = *defaultCompletion
	} else {
		output = t
	}

	fmt.Print(output)
}

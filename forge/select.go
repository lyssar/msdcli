package forge

import (
	"github.com/chzyer/readline"
	forgeVersionApi "github.com/kleister/go-forge/version"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"strings"
)

type noBellStdout struct{}

func (n *noBellStdout) Write(p []byte) (int, error) {
	if len(p) == 1 && p[0] == readline.CharBell {
		return 0, nil
	}
	return readline.Stdout.Write(p)
}

func (n *noBellStdout) Close() error {
	return readline.Stdout.Close()
}

var NoBellStdout = &noBellStdout{}

func RenderSelect(versions forgeVersionApi.Versions) (*forgeVersionApi.Version, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F32E {{ .ID | cyan }} (mc {{ .Minecraft | red }})",
		Inactive: "  {{ .ID | cyan }} (mc {{ .Minecraft | red }})",
		Selected: "\U0001F32E {{ .ID | red | cyan }}",
	}

	searcher := func(input string, index int) bool {
		versions := versions[index]
		name := strings.Replace(strings.ToLower(versions.ID), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Stdout:    NoBellStdout,
		Label:     "Select forge version",
		Items:     versions,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	cobra.CheckErr(err)

	return &versions[i], nil
}
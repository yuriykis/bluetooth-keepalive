package mac

import (
	"fmt"
	"os/exec"
	"strings"
)

// helper functions to run osascript commands

func run(command string) (string, error) {
	cmd := exec.Command("osascript", "-e", command)
	output, err := cmd.CombinedOutput()
	prettyOutput := strings.Replace(string(output), "\n", "", -1)
	if err != nil {
		return "", err
	}
	return prettyOutput, nil
}

func build(params ...string) string {
	var validParams []string
	for _, param := range params {
		if param != "" {
			validParams = append(validParams, param)
		}
	}
	return strings.Join(validParams, " ")
}

func wrapInQuotes(text string) string {
	return "\"" + text + "\""
}

func getCurrentVolume() string {
	return Build(
		"output volume of (get volume settings)",
	)
}

func setVolume(v float32) string {
	return Build(
		fmt.Sprintf("set volume output volume %f", v),
	)
}

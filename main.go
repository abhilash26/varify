package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
)

// Embed the default varify.json content
//go:embed varify.json
var defaultJSON []byte

type Module map[string]string

type Config struct {
	Output   string            `json:"output"`
	Selector string            `json:"selector"`
	Modules  map[string]Module `json:"modules"`
}

func convertToCSSVars(config Config) string {
	var cssContent strings.Builder
	cssContent.WriteString(config.Selector + " {\n")

	for module, values := range config.Modules {
		for key, value := range values {
			cssContent.WriteString(fmt.Sprintf("  --%s-%s: %s;\n", module, key, value))
		}
	}

	cssContent.WriteString("}\n")
	return cssContent.String()
}

func minifyCSS(cssContent string) (string, error) {
	m := minify.New()
	m.Add("text/css", &css.Minifier{})
	return m.String("text/css", cssContent)
}

func createDefaultJSON(filePath string) error {
	return os.WriteFile(filePath, defaultJSON, 0644)
}

func main() {
	inputFileFlag := flag.String("i", "", "Path to the input JSON file (uses embedded default if omitted)")
	outputFileFlag := flag.String("o", "", "Path to the output CSS file (overrides JSON output setting)")
	minifyFlag := flag.Bool("m", false, "Enable minification of the output CSS")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), `Usage:
  varify [options]

Options:
  -i <path>      Path to the input JSON file (default: embedded varify.json)
  -o <path>      Path to the output CSS file (overrides JSON output setting)
  -m             Enable minification of the output CSS
  -h             Show this help message
Example:
  varify -i ./input.json -o ./output.css -m
`)
	}

	flag.Parse()

	inputFile := *inputFileFlag
	overrideOutput := *outputFileFlag
	enableMinify := *minifyFlag

	var fileContent []byte
	var err error

	if inputFile == "" {
		inputFile = "./varify.json"
		if _, err := os.Stat(inputFile); os.IsNotExist(err) {
			if err := createDefaultJSON(inputFile); err != nil {
				fmt.Println("Error creating default varify.json:", err)
				return
			}
		}
		fileContent, err = os.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("Failed to read file %s: %v\n", inputFile, err)
			flag.Usage()
			return
		}
	} else {
		fileContent, err = os.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("Failed to read file %s: %v\n", inputFile, err)
			flag.Usage()
			return
		}
	}

	var config Config
	if err := json.Unmarshal(fileContent, &config); err != nil {
		fmt.Printf("Failed to parse JSON: %v\n", err)
		flag.Usage()
		return
	}

	if overrideOutput != "" {
		config.Output = overrideOutput
	}

	cssContent := convertToCSSVars(config)

	if enableMinify {
		cssContent, err = minifyCSS(cssContent)
		if err != nil {
			fmt.Printf("Error minifying CSS: %v\n", err)
			return
		}
	}

	if err := os.WriteFile(config.Output, []byte(cssContent), 0644); err != nil {
		fmt.Printf("Failed to write to file %s: %v\n", config.Output, err)
		return
	}
	fmt.Printf("CSS variables have been written to %s\n", config.Output)
}

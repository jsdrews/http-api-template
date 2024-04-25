package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	fileContents := make([]string, 0)

	if err := filepath.WalkDir("./api", walk(&fileContents)); err != nil {
		log.Fatal(err)
	}

	result, err := mergeYamlValues(fileContents)
	if err != nil {
		log.Fatal(err)
	}

	if err = os.WriteFile("./api/api.gen.yaml", []byte(result), 0o644); err != nil {
		log.Fatal(err)
	}
}

func walk(contents *[]string) func(path string, d fs.DirEntry, err error) error {
	return func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Exclude oapi codegen config file
		if entry.Name() == "config.yaml" {
			return nil
		}

		fileName := strings.ToLower(path)

		if strings.HasSuffix(fileName, "yaml") || strings.HasSuffix(fileName, "yml") {
			b, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			*contents = append(*contents, string(b))
		}

		return nil
	}
}

func mergeYamlValues(values []string) (string, error) {
	var result map[any]any

	var bs []byte

	for _, value := range values {
		var override map[any]any

		bs = []byte(value)

		if err := yaml.Unmarshal(bs, &override); err != nil {
			return "", err
		}

		// check if is nil. This will only happen for the first value
		if result == nil {
			result = override
		} else {
			result = mergeMaps(result, override)
		}
	}

	bs, err := yaml.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func mergeMaps(a, b map[any]any) map[any]any {
	out := make(map[any]any, len(a))
	for k, v := range a {
		out[k] = v
	}

	for k, v := range b {
		if v, ok := v.(map[any]any); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[any]any); ok {
					out[k] = mergeMaps(bv, v)
					continue
				}
			}
		}

		out[k] = v
	}

	return out
}
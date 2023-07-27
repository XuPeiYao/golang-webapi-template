package common

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Configuration map[string]any

const ConfigurationPathSeparator = ":"

func processPath(path string) []string {
	if len(path) == 0 {
		return []string{}
	}

	return strings.Split(path, ConfigurationPathSeparator)
}

func (this Configuration) GetValue(path ...string) any {
	return getValue(this, path...)
}

func (this Configuration) GetValueOrDefault(path string, defaultValue any) any {
	result := getValue(this, path)
	if result == nil {
		return defaultValue
	}
	return result
}

func getValue(config map[string]any, path ...string) any {
	anyPathHasSplitSymbol := false
	for _, pathItem := range path {
		anyPathHasSplitSymbol = anyPathHasSplitSymbol || strings.Contains(
			pathItem,
			ConfigurationPathSeparator,
		)
		if anyPathHasSplitSymbol {
			break
		}
	}

	if anyPathHasSplitSymbol == false {
		if len(path) == 0 {
			return nil
		}

		value := config[path[0]]

		if len(path) == 1 {
			return value
		}
		if value == nil && len(path) > 1 {
			return nil
		}

		if reflect.TypeOf(value).Implements(reflect.TypeOf((Configuration)(nil)).Elem()) {
			return getValue(value.(map[string]any), path[1:]...)
		} else {
			return nil
		}
	}

	splitedPath := []string{}

	for _, pathItem := range path {
		subSplitedPath := processPath(pathItem)
		for _, subPathItem := range subSplitedPath {
			splitedPath = append(splitedPath, subPathItem)
		}
	}

	return getValue(config, splitedPath...)
}

type ConfigurationFilePath string

func LoadConfigurationFromJsonFile(path ConfigurationFilePath) Configuration {
	var configuration Configuration

	// Read configuration file as bytes
	configurationBytes, err := os.ReadFile(string(path))

	// If can't read file, return empty configuration
	if err != nil {
		fmt.Printf("Not found configuration file: %s\n", path)
		configuration = make(Configuration)
		return configuration
	}

	// Parse configuration file
	err = json.Unmarshal(configurationBytes, &configuration)

	// If can't parse file, return empty configuration
	if err != nil {
		fmt.Printf("InnerError parsing configuration file: %s\n", path)
		configuration = make(Configuration)
	}

	return configuration
}

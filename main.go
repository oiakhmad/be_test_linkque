package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// InputFormat adalah struktur untuk input JSON
type InputFormat struct {
	InputFormat string `json:"input_format"`
}

// MappedResult adalah struktur untuk hasil mapping
type MappedResult struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	// Contoh input JSON
	inputJSON := `{
		"input_format": "Jhon Key aspe s se ee 12 Surabaya 12 Timur"
	}`

	// Parsing input JSON
	var input InputFormat
	err := json.Unmarshal([]byte(inputJSON), &input)
	if err != nil {
		fmt.Println("Error parsing input JSON:", err)
		return
	}

	// Proses mapping
	result, err := mapInputToStruct(input.InputFormat)
	if err != nil {
		fmt.Println("Error processing input format:", err)
		return
	}

	// Output hasil mapping dalam JSON
	outputJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Error generating output JSON:", err)
		return
	}

	fmt.Println(string(outputJSON))
}

func mapInputToStruct(input string) (MappedResult, error) {
	parts := strings.Fields(input)
	nameParts := []string{}
	age := ""
	cityParts := []string{}

	for i, part := range parts {
		if _, err := strconv.Atoi(part); err == nil {
			age = part
			cityParts = parts[i+1:]
			break
		}
		nameParts = append(nameParts, part)
	}

	if age == "" || len(cityParts) == 0 {
		return MappedResult{}, fmt.Errorf("invalid input format")
	}

	name := strings.ToUpper(strings.Join(nameParts, " "))
	city := strings.ToUpper(strings.ToLower(strings.Join(cityParts, " ")))

	ageInt, err := strconv.Atoi(age)
	if err != nil {
		return MappedResult{}, fmt.Errorf("invalid age format")
	}

	return MappedResult{
		Name: name,
		Age:  ageInt,
		City: city,
	}, nil
}

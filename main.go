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
		"input_format": "Jhon Key aspe s  12tahun Surabaya  Timur"
	}`

	// Parsing input JSON
	var input InputFormat
	err := json.Unmarshal([]byte(inputJSON), &input)
	if err != nil {
		fmt.Println("Error parsing input JSON:", err)
		return
	}

	// Proses mapping
	result, err := mapInputToStruct1(input.InputFormat)
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

func mapInputToStruct1(input string) (MappedResult, error) {
	parts := strings.Fields(input)
	nameParts := []string{}
	age := ""
	cityParts := []string{}

	for i, part := range parts {
		// Check if the part represents an age in various formats (e.g., "24", "24thun", "24Th")
		if numericPart, err := strconv.Atoi(part); err == nil || (len(part) > 2 && parseAgeSuffix(part, &numericPart) == nil) {
			age = strconv.Itoa(numericPart)
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

func parseAgeSuffix(input string, result *int) error {
	// Extract numeric prefix
	var numericPart strings.Builder
	for _, ch := range input {
		if ch >= '0' && ch <= '9' {
			numericPart.WriteRune(ch)
		} else {
			break
		}
	}

	if numericPart.Len() == 0 {
		return fmt.Errorf("no numeric part found")
	}

	parsed, err := strconv.Atoi(numericPart.String())
	if err != nil {
		return err
	}

	// Check for valid suffix
	suffix := strings.ToLower(input[numericPart.Len():])
	validSuffixes := []string{"tahun", "thun", "tahn", "th"}
	for _, valid := range validSuffixes {
		if suffix == valid {
			*result = parsed
			return nil
		}
	}

	return fmt.Errorf("invalid age suffix")
}

// Example structure

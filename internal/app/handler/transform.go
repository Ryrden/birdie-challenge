package handler

import (
	"encoding/json"
	"io"
	"os"
)

func TransformJSON(args []string) error {
	jsonFile, err := os.Open("input.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	file, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var data map[string]string
	if err := json.Unmarshal(file, &data); err != nil {
		return err
	}

	// Modify data
	for _, op := range args {
		_ = op
	}

	modifiedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Write data on output.json
	if err := os.WriteFile("output.json", modifiedData, 0644); err != nil {
		return err
	}

	return nil
}

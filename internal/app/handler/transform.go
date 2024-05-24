package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func HandleJSON(args []string) error {
	data, err := readFile("input.json")
	if err != nil {
		return err
	}

	if err := TransformJSON(data, args); err != nil {
		return err
	}

	return saveFile("output.json", data)
}

func readFile(filename string) (map[string]string, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	file, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var data map[string]string
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func saveFile(filename string, data map[string]string) error {
	modifiedDataBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filename, modifiedDataBytes, 0644); err != nil {
		return err
	}

	return nil
}

func TransformJSON(data map[string]string, operations []string) (error) {
	if len(operations) == 0 {
		log.Println("No operations to perform")
		return nil
	}

	for _, op := range operations {
		parts := strings.Split(op, ":")
		if len(parts) < 2 {
			return fmt.Errorf("invalid operation: %s", op)
		}
		operation, key := parts[0], parts[1]

		switch operation {
		case "set":
			if len(parts) != 3 {
				log.Printf("invalid operation: %s, ignoring", op)
				continue
			}
			data[key] = parts[2]
		case "rename":
			if len(parts) != 3 {
				log.Printf("invalid operation: %s, ignoring", op)
				continue
			}
			value, ok := data[key];
			if !ok {
				log.Printf("key %s not found for operation %s", key, op)
				continue
			}
			newKey := parts[2]
			delete(data, key)
			data[newKey] = value
		
		case "delete":
			_, ok := data[key]
			if !ok {
				log.Printf("key %s not found for operation %s", key, op)
				continue
			}
			delete(data, key)
		}
	}

	return  nil
}

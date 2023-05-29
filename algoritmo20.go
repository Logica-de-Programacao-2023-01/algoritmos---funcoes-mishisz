package main

import "fmt"

func filterStrings(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, fmt.Errorf("O slice está vazio")
	}

	result := make([]string, 0)

	for _, str := range slice {
		if len(str) > 5 {
			result = append(result, str)
		}
	}

	return result, nil
}

package main

import (
	"errors"
	"fmt"
	"strings"
)

func concatenarComVirgulas(stringsSlice []string) (string, error) {
	if len(stringsSlice) == 0 {
		return "", errors.New("Slice está vazio")
	}

	concatenado := strings.Join(stringsSlice, ",")
	return concatenado, nil
}

package util

import "strconv"

func EsVacio(value string) bool {
	return len(value) == 0
}

func EsNumero(value string) bool {
	if EsVacio(value) {
		return false
	}
	_, err := strconv.Atoi(value)
	return err == nil
}

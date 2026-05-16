package utils

import "strings"

type Field struct {
	Name string
	Type string
}

func ParseFields(input string) []Field {

	fields := []Field{}

	if input == "" {
		return fields
	}

	pairs := strings.Split(input, ",")

	for _, pair := range pairs {

		parts := strings.Split(pair, ":")

		if len(parts) != 2 {
			continue
		}

		fields = append(fields, Field{
			Name: strings.TrimSpace(parts[0]),
			Type: strings.TrimSpace(parts[1]),
		})
	}

	return fields
}
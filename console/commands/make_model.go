package commands

import "github.com/shehabalqudiry/alqudiry_gocore/console/generators"

func MakeModel(
	name string,
	fields string,
) {

	generators.GenerateModel(
		name,
		fields,
	)
}
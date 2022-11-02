package asset

import (
	"go-boilerplate/core"
	"go-boilerplate/store/db"
)

func scanRow(scanner db.Scanner, output *core.Asset) error {
	if scanner.Next() {
		if err := scanner.StructScan(
			output,
		); err != nil {
			return err
		}
	}
	defer scanner.Close()

	return nil
}

func scanRows(scanner db.Scanner, outputs []*core.Asset) error {
	for scanner.Next() {
		output := &core.Asset{}
		err := scanner.StructScan(output)
		if err != nil {
			return err
		}
		outputs = append(outputs, output)
	}
	defer scanner.Close()

	return nil
}

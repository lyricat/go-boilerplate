package snapshot

import (
	"go-boilerplate/core"
	"go-boilerplate/store/db"
)

func scanRow(scanner db.Scanner, output *core.Snapshot) error {
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

func scanRows(scanner db.Scanner, outputs []*core.Snapshot) error {
	for scanner.Next() {
		output := &core.Snapshot{}
		err := scanner.StructScan(output)
		if err != nil {
			return err
		}
		outputs = append(outputs, output)
	}
	defer scanner.Close()

	return nil
}

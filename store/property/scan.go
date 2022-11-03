package property

import (
	"go-boilerplate/core"
	"go-boilerplate/store/db"
)

func scanRow(scanner db.Scanner, pp *core.Property) error {
	defer scanner.Close()

	if scanner.Next() {
		if err := scanner.Scan(
			&pp.Key,
			&pp.Value,
			&pp.UpdatedAt,
		); err != nil {
			return err
		}
	}

	return nil
}

package utils

import "fmt"

type saver interface {
	Save() error // All structs should have same method or data in which we want to use this same interface.
	Print()
	PrintAll() error
}

func SaveData(data saver) error {
	err := data.Save()

	if err != nil {
		return err
	}

	data.PrintAll()
	fmt.Println()

	return nil
}

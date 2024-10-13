package fileoperations

import (
	"fmt"
	"io/fs"
	"os"
)

func WriteToFile(value [] byte, filePath string, permissions fs.FileMode) {
	err := os.WriteFile(filePath, value, permissions)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Value successfully written to file.")

}


func ReadFromFile (filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(data)
}
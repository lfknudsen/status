package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func Err(context string, err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, context+":\n", err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: status <code> or status <title> or status <all>")
		os.Exit(0)
	}
	e, err := os.Executable()
	Err("Failed to get executable path", err)
	filename := filepath.Join(filepath.Dir(e), ".status-codes.csv")
	code := os.Args[1]
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	Err("Failed to read file", err)
	defer closeFile(file)

	reader := csv.NewReader(bufio.NewReader(file))
	records, err := reader.ReadAll()
	Err("Failed to parse csv", err)
	records = records[1:] // skip header row
	if code == "all" {
		for _, record := range records {
			fmt.Printf("%s %s\n%s\n", record[0], record[1], record[2])
		}
	} else {
		for _, record := range records {
			if record[0] == code || record[1] == code {
				fmt.Printf("%s %s\n%s\n", record[0], record[1], record[2])
				os.Exit(0)
			}
		}
		fmt.Printf("Could not find status \"%s\"\n", code)
	}
}

func closeFile(f *os.File) {
	if f != nil {
		_ = f.Close()
	}
}

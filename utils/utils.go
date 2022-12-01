package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCsvFile(day string) [][]string {
	f, err := os.Open("./inputs/day" + day + ".csv")
	if err != nil {
		log.Fatal("Unable to read input file day"+day, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for day"+day, err)
	}

	return records
}

func ReadTxtFile(filePath string) []string {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines
}

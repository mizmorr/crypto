package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// curl https://www.ece.unb.ca/tervo/ece4253/polyprime.shtml | grep -w binary > test.txt

func get_irreducible() {
	readFile, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	file, err := os.Create("irreducible.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	for fileScanner.Scan() {
		array := strings.Split(fileScanner.Text(), " ")
		for _, a := range array {
			if strings.Contains(a, "binary") {
				file.WriteString(strings.Split(strings.Split(strings.Split(a, "?")[1], "=")[1], "\"")[0] + "\n")
			}
		}

	}
	readFile.Close()
}

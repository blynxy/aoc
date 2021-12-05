package main

import (
	"bufio"
	"log"
	"os"
)
func FileByLine(filename string) []string {
	var r []string
	file, err := os.Open(filename)
	if err != nil{
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return r
}




func FileByString(filename string) []string {
	var r []string
	file, err := os.Open(filename)
	if err != nil{
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return r
}
package main

import (
	"geektrust/service"
	"bufio"
	
	"os"
	"log"
)

func main() {
	
	//option1
	file, err := os.Open("testfiles/input1.txt")
	//option 2
    // argsWithoutProg := os.Args[1]
	// file,err:=os.Open(argsWithoutProg)
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		log.Println("Geting lines")
		txtlines = append(txtlines, scanner.Text())
	}
 service.ProcessInput(txtlines)
	defer file.Close()
}

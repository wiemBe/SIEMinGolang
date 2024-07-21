package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	
)

func Filter(fileName string ) {
	file , err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sourceIps []string
	var destinationIps []string
	seenSourceIps := make(map[string]bool)
	seenDestinationIps := make(map[string]bool)

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()

		sourceIpRegex := regexp.MustCompile(`This is Source ip: ([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)`)
		destinationIpRegex := regexp.MustCompile(`This is Destination ip: ([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)`)

		if matches := sourceIpRegex.FindStringSubmatch(line); matches != nil {
			ip := matches[1]
			if !seenSourceIps[ip]{
				sourceIps = append(sourceIps, ip)
				seenSourceIps[ip] = true
				
			}
		}
		if matches := destinationIpRegex.FindStringSubmatch(line); matches != nil {
			ip := matches[1]
			if !seenDestinationIps[ip]{
				destinationIps = append(destinationIps, ip)
				seenDestinationIps[ip] = true
			}
			destinationIps =append(destinationIps, matches[1])
		}
	}

	if err := scanner.Err(); err != nil{
		fmt.Println("error reading file:", err)
	}
	fmt.Println("Source Ips", sourceIps)
	fmt.Println("destination ip " , destinationIps)



	

}
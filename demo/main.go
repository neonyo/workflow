package main

import (
	"context"
	"demo/parse"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	data, err := parseFile("diagram_1.bpmn")
	if err != nil {
		log.Fatalln(err)
	}
	xml := parse.NewXMLParser()
	result, err := xml.Parse(context.Background(), data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result.FlowID, result.FlowVersion, result.FlowVersion)
	for _, v := range result.Nodes {
		fmt.Println(v.CandidateExpressions, v.NodeType, v.NodeID, v.NodeName, v.FormResult, v.Properties, v.Routers, "=")
	}
}

func parseFile(name string) ([]byte, error) {
	fullName, err := filepath.Abs(name)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(fullName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

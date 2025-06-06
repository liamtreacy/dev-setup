package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetKubeNS() ([]string, error) {
	var kubeDefFile = "config/kubenamespaces"

	file, err := os.Open(kubeDefFile)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	defer file.Close()
	return lines, nil
}

func main() {
	// Load in file with names of kubenamespaces of interest
	// That file will not be saved to git
	//
	// Options:
	// 	Show what kubenamespaces can be done
	//  Get directory of caller
	//	Get logs of selected kube (to calling dir)
	dir, err := os.Getwd()
	kube_namespaces, err := GetKubeNS()
	fmt.Println(kube_namespaces)

	if err == nil {
		fmt.Println("hello! you are calling from: ", dir)
	}
}

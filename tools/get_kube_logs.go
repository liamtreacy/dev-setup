package main

import (
	"fmt"
	"os"
	"bufio"
)

func main()  {
	// Load in file with names of kubenamespaces of interest
	// That file will not be saved to git
	//
	// Options:
	// 	Show what kubenamespaces can be done
	//  Get directory of caller
	//	Get logs of selected kube (to calling dir)
	dir, err := os.Getwd()
	var kubeDefFile = "config/kubenamespaces"
	file, err := os.Open(kubeDefFile)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err == nil{
	fmt.Println("hello! you are calling from: ", dir)
	}
}

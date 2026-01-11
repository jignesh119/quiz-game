package main

import (
	"flag"
	"fmt"
	"os"
)

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}

func main(){
	csvFileName:=flag.String("csv","quiz.csv","a csv file in the format of 'problem,solution'")
	_=csvFileName
	flag.Parse()

	file,err:=os.Open(*csvFileName)
	if err!=nil{
		exit(fmt.Sprintf("Failed to open file %v ",*csvFileName))
	}
	_=file
}

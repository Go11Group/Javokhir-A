package main

import (
	"encoding/json"
	"excel/structs"
	"log"
	"os"
//	"reflect"
	"strconv"
	"fmt"
	"github.com/xuri/excelize/v2"
)

// func GetFielsNames(obj interface{}) {
// 	obj = obj.(structs.Person)
// 	t := reflect.TypeOf(obj)

// 	for
// }

func main() {
	people := []*structs.Person{}
	personCredentials := []string{"Id", "FirstName", "LastName", "Email", "Age"}

	file, err := os.Open("jsons/People.json")
	if err != nil {
		log.Fatal("21", err)
	}

	// Decoding data in people.json file to people slice of person struct
	decoder := json.NewDecoder(file)
	decoder.Decode(&people)

	excelFile := excelize.NewFile()
	_, err = excelFile.NewSheet("People")
	if err != nil {
		log.Fatal("31", err)
	}

	for i, val := range personCredentials {
		cell := string(byte(65+i)) + "1"
		excelFile.SetCellValue("People", cell, val)
	}
	row := 2
	col := 0
	for _, person := range people {
		excelFile.SetCellValue("People", string(byte(65+col))+strconv.Itoa(row), person.Id)
		excelFile.SetCellValue("People", string(byte(65+col+1))+strconv.Itoa(row), person.FirstName)
		excelFile.SetCellValue("People", string(byte(65+col+2))+strconv.Itoa(row), person.LastName)
		excelFile.SetCellValue("People", string(byte(65+col+3))+strconv.Itoa(row), person.Email)
		excelFile.SetCellValue("People", string(byte(65+col+4))+strconv.Itoa(row), person.Age)
		row++
	}

	fmt.Println("Hi")

	err = excelFile.SaveAs("People.xlsx")
	if err != nil {
		log.Fatal("40", err)
	}
	
}

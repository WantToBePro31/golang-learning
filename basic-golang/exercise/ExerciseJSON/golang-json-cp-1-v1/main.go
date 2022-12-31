package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Study struct {
	StudyName string `json:"study_name"`
	StudyCredit int `json:"study_credit"`
	Grade string `json:"grade"`
}

type Report struct {
	// TODO: answer here
	Id string `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
	Semester int `json:"Semester"`
	Studies []Study
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	// TODO: answer here
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return Report{}, err
	}
	var data Report
	err = json.Unmarshal(jsonData, &data)
	// var data map[string]string
	// err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return Report{}, err
	}
	return data, nil
}

func GradePoint(report Report) float64 {
	if len(report.Studies) == 0 {
		return 0.0
	}
	grade := map[string]float64 {
		"A": 4.0,
		"AB": 3.5,
		"B": 3.0,
		"BC": 2.5,
		"C": 2.0,
		"CD": 1.5,
		"D": 1.0,
		"DE": 0.5,
		"E": 0.0,
	}
	totalScore, totalCredit := 0.0, 0.0
	for _, val := range report.Studies {
		credit := float64(val.StudyCredit)
		score := grade[val.Grade]
		totalScore += credit * score
		totalCredit += credit
	}
	return totalScore/totalCredit // TODO: replace this
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}

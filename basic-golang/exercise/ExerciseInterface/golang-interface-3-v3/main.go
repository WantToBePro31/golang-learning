package main

import (
	"fmt"
	"strconv"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time any) string {
	switch time.(type) {
	case string:
		timeStr := time.(string)
		if len(timeStr) != 5 {
			return "Invalid input"
		}
		if timeStr[0] == '0' || (timeStr[0] == '1' && (timeStr[1] >= '0' && timeStr[1] < '2')) {
			return fmt.Sprintf("%d%d:%d%d AM", int(timeStr[0])-48, int(timeStr[1])-48, int(timeStr[3])-48, int(timeStr[4])-48)
		} else if timeStr[0] == '1' && timeStr[1] == '2' && timeStr[3] == '0' && timeStr[4] == '0' {
			return "12:00 PM"
		} else {
			hour, _ := strconv.Atoi(timeStr[:2])
			if hour-12 < 10 {
				return fmt.Sprintf("0%d:%d%d PM", hour-12, int(timeStr[3])-48, int(timeStr[4])-48)
			}
			return fmt.Sprintf("%d:%d%d PM", hour-12, int(timeStr[3])-48, int(timeStr[4])-48)
		}
	case []int:
		timeSlice := time.([]int)
		if len(timeSlice) != 2 {
			return "Invalid input"
		}
		var timeMinute string
		if timeSlice[1] < 10 {
			timeMinute = fmt.Sprintf("0%d", timeSlice[1])
		} else {
			timeMinute = strconv.Itoa(timeSlice[1])
		}
		if timeSlice[0] < 10 {
			return fmt.Sprintf("0%d:%s AM", timeSlice[0], timeMinute)
		} else if timeSlice[0] < 12{
			return fmt.Sprintf("%d:%s AM", timeSlice[0], timeMinute)
		} else if timeSlice[0] == 12 && timeSlice[1] == 0 {
			return "12:00 PM"
		} else {
			if timeSlice[0]-12 < 10 {
				return fmt.Sprintf("0%d:%s PM", timeSlice[0]-12, timeMinute)
			}
			return fmt.Sprintf("%d:%s PM", timeSlice[0]-12, timeMinute)
		}
	case map[string]int:
		timeMap := time.(map[string]int)
		_, ok_hour := timeMap["hour"]
		_, ok_minute := timeMap["minute"]
		if !ok_hour || !ok_minute {
			return "Invalid input"
		}
		var timeMinute string
		if timeMap["minute"] < 10 {
			timeMinute = fmt.Sprintf("0%d", timeMap["minute"])
		} else {
			timeMinute = strconv.Itoa(timeMap["minute"])
		}
		if timeMap["hour"] < 10 {
			return fmt.Sprintf("0%d:%s AM", timeMap["hour"], timeMinute)
		} else if timeMap["hour"] < 12 {
			return fmt.Sprintf("%d:%s AM", timeMap["hour"], timeMinute)
		} else if timeMap["hour"] == 12 && timeMap["minute"] == 0 {
			return "12:00 PM"
		} else {
			if timeMap["hour"]-12 < 10 {
				return fmt.Sprintf("0%d:%s PM", timeMap["hour"]-12, timeMinute)
			}
			return fmt.Sprintf("%d:%s PM", timeMap["hour"]-12, timeMinute)
		}
	case Time:
		timeStruct := time.(Time)
		var timeMinute string
		if timeStruct.Minute < 10 {
			timeMinute = fmt.Sprintf("0%d", timeStruct.Minute)
		} else {
			timeMinute = strconv.Itoa(timeStruct.Minute)
		}
		if timeStruct.Hour < 10 {
			return fmt.Sprintf("0%d:%s AM", timeStruct.Hour, timeMinute)
		} else if timeStruct.Hour < 12 {
			return fmt.Sprintf("%d:%s AM", timeStruct.Hour, timeMinute)
		} else if timeStruct.Hour == 12 && timeStruct.Minute == 0 {
			return "12:00 PM"
		} else {
			if timeStruct.Hour-12 < 10 {
				return fmt.Sprintf("0%d:%s PM", timeStruct.Hour-12, timeMinute)
			}
			return fmt.Sprintf("%d:%s PM", timeStruct.Hour-12, timeMinute)
		}
	default:
		return "Invalid input"
	}
	return "Invalid input" // TODO: replace this
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var reader *bufio.Reader

type Parser interface {
	ParseMinutes(string) []int
	ParseHours(string) []int
	ParseDaysOfMonth(string) []int
	ParseMonths(string) []int
	ParseDaysOfWeek(string) []int
	ParseCommand(string) string
}

type CronData struct {
	Minutes     []int
	Hours       []int
	DaysOfMonth []int
	Months      []int
	DaysOfWeek  []int
	Command     string
}

type ArgContainer struct {
	Minutes     string
	Hours       string
	DaysOfMonth string
	Months      string
	DaysOfWeek  string
	Command     string
}

var NewArgContainer ArgContainer
var NewCronData CronData

func (cd *CronData) ParseMinutes(minutes string) []int {
	result := []int{}
	hasAsterist, stringLength := DoesContainAsterisk(minutes)
	if hasAsterist {
		if stringLength > 1 {
			tempMinutes := strings.Split(minutes, "/")
			minuteVal, err := strconv.Atoi(tempMinutes[1])
			if err != nil {
				fmt.Printf("hasAsterisk: Could not convert minute data to int:%v", err)
			}

			for i := 0; i < 60; i += minuteVal { // As the result should not contain 60
				result = append(result, i)
			}

		} else if stringLength == 1 {
			allMinutes := []int{}
			for i := 1; i <= 60; i++ {
				allMinutes = append(allMinutes, i)
			}
			result = allMinutes
		}
	}

	hasComma := DoesContainComma(minutes)
	if hasComma {
		tempMinutes := strings.Split(minutes, ",")
		for _, x := range tempMinutes {
			tempMinutesInt, _ := strconv.Atoi(x)
			result = append(result, tempMinutesInt)
		}
	}

	hasHyphen := DoesContainHyphen(minutes)
	if hasHyphen {
		tempMinutes := strings.Split(minutes, "-")
		minuteStart, err := strconv.Atoi(tempMinutes[0])
		if err != nil {
			fmt.Printf("hasHyphen: Could not convert minutes data to int:%v", err)
		}
		minuteStop, err := strconv.Atoi(tempMinutes[1])
		if err != nil {
			fmt.Printf("hasHyphen: Could not convert hours data to int:%v", err)
		}
		for i := minuteStart; i <= minuteStop; i++ {
			result = append(result, i)
		}
	}

	if !hasAsterist && !hasComma && !hasHyphen {
		tempMinutes, err := strconv.Atoi(minutes)
		if err != nil {
			fmt.Printf("Doesn't Contain any Special character: Could not convert minute data to int:%v", err)
		}

		result = append(result, tempMinutes)
	}

	// fmt.Println(result)
	return result

}

func (cd *CronData) ParseHours(hours string) []int {
	result := []int{}
	hasAsterist, stringLength := DoesContainAsterisk(hours)
	if hasAsterist {
		if stringLength > 1 {
			tempHours := strings.Split(hours, "/")
			hourVal, err := strconv.Atoi(tempHours[1])
			if err != nil {
				fmt.Printf("hasAsterisk: Could not convert hours data to int:%v", err)
			}

			for i := 0; i <= 23; i += hourVal { // As the result should not contain 24 hours
				result = append(result, i)
			}

		} else if stringLength == 1 {
			allHours := []int{}
			for i := 0; i <= 23; i++ {
				allHours = append(allHours, i)
			}
			result = allHours
		}
	}

	hasComma := DoesContainComma(hours)
	if hasComma {
		tempHours := strings.Split(hours, ",")
		for _, x := range tempHours {
			tempHoursInt, _ := strconv.Atoi(x)
			result = append(result, tempHoursInt)
		}
	}

	hasHyphen := DoesContainHyphen(hours)
	if hasHyphen {
		tempHours := strings.Split(hours, "-")
		hourStart, err := strconv.Atoi(tempHours[0])
		if err != nil {
			fmt.Printf("hasHyphen: Could not convert hours data to int:%v", err)
		}
		hourStop, err := strconv.Atoi(tempHours[1])
		if err != nil {
			fmt.Printf("hasHyphen: Could not convert hours data to int:%v", err)
		}
		for i := hourStart; i <= hourStop; i++ {
			result = append(result, i)
		}
	}

	if !hasAsterist && !hasComma && !hasHyphen {
		tempHours, err := strconv.Atoi(hours)
		if err != nil {
			fmt.Printf("Doesn't Contain any Special character: Could not convert hour data to int:%v", err)
		}

		result = append(result, tempHours)
	}
	// fmt.Println(result)
	return result
}

func (cd *CronData) ParseDaysOfMonth(daysOfMonth string) []int {
	result := []int{}
	hasAsterist, stringLength := DoesContainAsterisk(daysOfMonth)
	if hasAsterist {
		if stringLength > 1 {
			tempDaysOfMonth := strings.Split(daysOfMonth, "/")
			daysOfMonthVal, err := strconv.Atoi(tempDaysOfMonth[1])
			if err != nil {
				fmt.Printf("hasAsterisk: Could not convert DaysOfMonth data to int:%v", err)
			}

			for i := 0; i <= 31; i += daysOfMonthVal {

				if i == 0 { // for cronjob, Day of the month: 1-31
					i = 1
					result = append(result, i)
					i = 0
					continue
				}
				result = append(result, i)

			}

		} else if stringLength == 1 {
			allDaysOfMonthVal := []int{}
			for i := 1; i <= 30; i++ {
				allDaysOfMonthVal = append(allDaysOfMonthVal, i)
			}
			result = allDaysOfMonthVal
		}
	}

	hasComma := DoesContainComma(daysOfMonth)
	if hasComma {
		tempDaysOfMonth := strings.Split(daysOfMonth, ",")
		for _, x := range tempDaysOfMonth {
			tempDaysOfMonthInt, _ := strconv.Atoi(x)
			result = append(result, tempDaysOfMonthInt)
		}
	}

	hasHyphen := DoesContainHyphen(daysOfMonth)
	if hasHyphen {
		tempDaysOfMonth := strings.Split(daysOfMonth, "-")
		daysOfMonthStart, err := strconv.Atoi(tempDaysOfMonth[0])
		if err != nil {
			fmt.Printf("hasHyphen: Could not convert daysOfMonth data to int:%v", err)
		}
		daysOfMonthStop, err := strconv.Atoi(tempDaysOfMonth[1])
		if err != nil {
			fmt.Printf("hasHyphen: Could not convert daysOfMonth data to int:%v", err)
		}
		for i := daysOfMonthStart; i <= daysOfMonthStop; i++ {
			result = append(result, i)
		}
	}

	if !hasAsterist && !hasComma && !hasHyphen {
		tempDaysOfMonth, err := strconv.Atoi(daysOfMonth)
		if err != nil {
			fmt.Printf("Doesn't Contain any Special character: Could not convert hour data to int:%v", err)
		}

		result = append(result, tempDaysOfMonth)
	}
	// fmt.Println(result)
	return result
}

func (cd *CronData) ParseMonths(months string) []int {
	result := []int{}
	hasAsterist, stringLength := DoesContainAsterisk(months)
	if hasAsterist {
		if stringLength > 1 {
			tempMonths := strings.Split(months, "/")
			monthsVal, err := strconv.Atoi(tempMonths[1])
			if err != nil {
				fmt.Printf("hasAsterisk: Could not convert month data to int:%v", err)
			}

			for i := 0; i <= 12; i += monthsVal {
				if i == 0 { // for cronjob, Month:	1-12 (or JAN to DEC)
					i = 1
					result = append(result, i)
					i = 0
					continue
				}
				result = append(result, i)
			}

		} else if stringLength == 1 {
			allMonthsVal := []int{}
			for i := 1; i <= 12; i++ {
				allMonthsVal = append(allMonthsVal, i)
			}
			result = allMonthsVal
		}
	}

	hasComma := DoesContainComma(months)
	if hasComma {
		tempMonths := strings.Split(months, ",")
		for _, x := range tempMonths {
			tempMonthsInt, _ := strconv.Atoi(x)
			result = append(result, tempMonthsInt)
		}
	}

	hasHyphen := DoesContainHyphen(months)
	if hasHyphen {
		tempMonths := strings.Split(months, "-")
		monthsStart, err := strconv.Atoi(tempMonths[0])
		if err != nil {
			fmt.Printf("hasHyphen: Could not convert months data to int:%v", err)
		}
		monthsStop, err := strconv.Atoi(tempMonths[1])
		if err != nil {
			fmt.Printf("hasHyphen: Could not convert months data to int:%v", err)
		}
		for i := monthsStart; i <= monthsStop; i++ {
			result = append(result, i)
		}
	}

	if !hasAsterist && !hasComma && !hasHyphen {
		tempMonths, err := strconv.Atoi(months)
		if err != nil {
			fmt.Printf("Doesn't Contain any Special character: Could not convert months data to int:%v", err)
		}

		result = append(result, tempMonths)
	}
	// fmt.Println(result)
	return result
}

func (cd *CronData) ParseDaysOfWeek(daysOfWeek string) []int {
	result := []int{}
	hasAsterist, stringLength := DoesContainAsterisk(daysOfWeek)
	if hasAsterist {
		if stringLength > 1 {
			tempDaysOfWeek := strings.Split(daysOfWeek, "/")
			daysOfWeekVal, err := strconv.Atoi(tempDaysOfWeek[1])
			if err != nil {
				fmt.Printf("hasAsterisk: Could not convert daysOfWeek data to int:%v", err)
			}

			for i := 1; i <= 7; i += daysOfWeekVal {
				result = append(result, i)
			}

		} else if stringLength == 1 {
			allDaysOfWeekVal := []int{}
			for i := 1; i <= 7; i++ {
				allDaysOfWeekVal = append(allDaysOfWeekVal, i)
			}
			result = allDaysOfWeekVal
		}
	}

	hasComma := DoesContainComma(daysOfWeek)
	if hasComma {
		tempDaysOfWeek := strings.Split(daysOfWeek, ",")
		for _, x := range tempDaysOfWeek {
			tempDaysOfWeekInt, _ := strconv.Atoi(x)
			result = append(result, tempDaysOfWeekInt)
		}
	}

	hasHyphen := DoesContainHyphen(daysOfWeek)
	if hasHyphen {
		tempDaysOfWeek := strings.Split(daysOfWeek, "-")
		tempDaysOfWeekStart, err := strconv.Atoi(tempDaysOfWeek[0])
		if err != nil {
			fmt.Printf("hasHyphen: Could not convert daysOfWeek data to int:%v", err)
		}
		tempDaysOfWeekStop, err := strconv.Atoi(tempDaysOfWeek[1])
		if err != nil {
			fmt.Printf("hasHyphen: Could not convert daysOfWeek data to int:%v", err)
		}
		for i := tempDaysOfWeekStart; i <= tempDaysOfWeekStop; i++ {
			result = append(result, i)
		}
	}

	if !hasAsterist && !hasComma && !hasHyphen {
		tempDaysOfWeek, err := strconv.Atoi(daysOfWeek)
		if err != nil {
			fmt.Printf("Doesn't Contain any Special character: Could not convert daysOfWeek data to int:%v", err)
		}

		result = append(result, tempDaysOfWeek)
	}
	// fmt.Println(result)
	return result
}

func (cd *CronData) ParseCommand(command string) string {
	result := ""
	// fmt.Println("got data:", command)
	for i := 0; i <= len(command)-1; i++ {
		result += string(command[i])
	}
	// fmt.Println(result)
	return result
}

func DoesContainAsterisk(newString string) (bool, int) {
	result := strings.Contains(newString, "*")
	return result, len(newString)
}

func DoesContainComma(newString string) bool {
	return strings.Contains(newString, ",")
}

func DoesContainHyphen(newString string) bool {
	return strings.Contains(newString, "-")
}

func CalculateValues(parserObject Parser) {
	NewCronData.Minutes = parserObject.ParseMinutes(NewArgContainer.Minutes)
	NewCronData.Hours = parserObject.ParseHours(NewArgContainer.Hours)
	NewCronData.DaysOfMonth = parserObject.ParseDaysOfMonth(NewArgContainer.DaysOfMonth)
	NewCronData.Months = parserObject.ParseMonths(NewArgContainer.Months)
	NewCronData.DaysOfWeek = parserObject.ParseDaysOfWeek(NewArgContainer.DaysOfWeek)
	NewCronData.Command = parserObject.ParseCommand(NewArgContainer.Command)
	fmt.Printf("minute \t\t\t %v\n", NewCronData.Minutes)
	fmt.Printf("hour \t\t\t %v\n", NewCronData.Hours)
	fmt.Printf("day of month \t\t %v\n", NewCronData.DaysOfMonth)
	fmt.Printf("month \t\t\t %v\n", NewCronData.Months)
	fmt.Printf("day of week \t\t %v\n", NewCronData.DaysOfWeek)
	fmt.Printf("command \t\t %v\n", NewCronData.Command)
}

func main() {

	newArguments := os.Args[1:]

	allArguments := strings.Split(newArguments[0], " ")
	// fmt.Println("allArguments:", allArguments)

	for id, arg := range allArguments {
		switch id {
		case 0:
			NewArgContainer.Minutes = arg
		case 1:
			NewArgContainer.Hours = arg
		case 2:
			NewArgContainer.DaysOfMonth = arg
		case 3:
			NewArgContainer.Months = arg
		case 4:
			NewArgContainer.DaysOfWeek = arg
		case 5:
			NewArgContainer.Command = arg
		}
	}

	var newParser Parser = &NewCronData

	CalculateValues(newParser)
}

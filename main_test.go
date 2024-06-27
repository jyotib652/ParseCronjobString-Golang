package main

import (
	"testing"
)

func TestParseMinutes(t *testing.T) {
	testCases := []struct {
		name     string
		minutes  string
		expected []int
	}{
		{
			name:     "minutes in string form without any special characters",
			minutes:  "10",
			expected: []int{10},
		},
		{
			name:     "minutes with * in string form ",
			minutes:  "*/10",
			expected: []int{0, 10, 20, 30, 40, 50},
		},
		{
			name:    "minutes with only * in string form ",
			minutes: "*",
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
				21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
				31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
				41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
				51, 52, 53, 54, 55, 56, 57, 58, 59},
		},
		{
			name:     "minutes with comma in string form ",
			minutes:  "10,25",
			expected: []int{10, 25},
		},
		{
			name:     "minutes with hyphen in string form ",
			minutes:  "10-17",
			expected: []int{10, 11, 12, 13, 14, 15, 16, 17},
		},
	}

	for _, tc := range testCases {
		var testCronData CronData
		testResult := testCronData.ParseMinutes(tc.minutes)
		if len(testResult) != len(tc.expected) {
			t.Errorf("%s failed!, expected:%v, got:%v", tc.name, len(tc.expected), len(testResult))
		}
	}
}

func TestParseHours(t *testing.T) {
	testCases := []struct {
		name     string
		hours    string
		expected []int
	}{
		{
			name:     "hours in string form without any special characters",
			hours:    "10",
			expected: []int{10},
		},
		{
			name:     "hours with * in string form ",
			hours:    "*/10",
			expected: []int{0, 10, 20},
		},
		{
			name:  "hours with only * in string form ",
			hours: "*",
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
				21, 22, 23},
		},
		{
			name:     "hours with comma in string form ",
			hours:    "10, 23",
			expected: []int{10, 23},
		},
		{
			name:     "hours with hyphen in string form ",
			hours:    "10-17",
			expected: []int{10, 11, 12, 13, 14, 15, 16, 17},
		},
	}

	for _, tc := range testCases {
		var testCronData CronData
		testResult := testCronData.ParseHours(tc.hours)
		if len(testResult) != len(tc.expected) {
			t.Errorf("%s failed!, expected:%v, got:%v", tc.name, len(tc.expected), len(testResult))
		}
	}
}

func TestParseDaysOfMonth(t *testing.T) {
	testCases := []struct {
		name        string
		daysOfMonth string
		expected    []int
	}{
		{
			name:        "daysOfMonth in string form without any special characters",
			daysOfMonth: "10",
			expected:    []int{10},
		},
		{
			name:        "daysOfMonth with * in string form ",
			daysOfMonth: "*/10",
			expected:    []int{1, 10, 20, 30},
		},
		{
			name:        "daysOfMonth with only * in string form ",
			daysOfMonth: "*",
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
				21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		},
		{
			name:        "daysOfMonth with comma in string form ",
			daysOfMonth: "10, 23",
			expected:    []int{10, 23},
		},
		{
			name:        "daysOfMonth with hyphen in string form ",
			daysOfMonth: "10-17",
			expected:    []int{10, 11, 12, 13, 14, 15, 16, 17},
		},
	}

	for _, tc := range testCases {
		var testCronData CronData
		testResult := testCronData.ParseDaysOfMonth(tc.daysOfMonth)
		if len(testResult) != len(tc.expected) {
			t.Errorf("%s failed!, expected:%v, got:%v", tc.name, len(tc.expected), len(testResult))
		}
	}
}

func TestParseMonths(t *testing.T) {
	testCases := []struct {
		name     string
		months   string
		expected []int
	}{
		{
			name:     "months in string form without any special characters",
			months:   "10",
			expected: []int{10},
		},
		{
			name:     "months with * in string form ",
			months:   "*/5",
			expected: []int{1, 5, 10},
		},
		{
			name:   "months with only * in string form ",
			months: "*",
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				11, 12},
		},
		{
			name:     "months with comma in string form ",
			months:   "10, 11",
			expected: []int{10, 11},
		},
		{
			name:     "months with hyphen in string form ",
			months:   "7-12",
			expected: []int{7, 8, 9, 10, 11, 12},
		},
	}

	for _, tc := range testCases {
		var testCronData CronData
		testResult := testCronData.ParseMonths(tc.months)
		if len(testResult) != len(tc.expected) {
			t.Errorf("%s failed!, expected:%v, got:%v", tc.name, len(tc.expected), len(testResult))
		}
	}
}

func TestParseDaysOfWeek(t *testing.T) {
	testCases := []struct {
		name       string
		daysOfWeek string
		expected   []int
	}{
		{
			name:       "daysOfWeek in string form without any special characters",
			daysOfWeek: "7",
			expected:   []int{7},
		},
		{
			name:       "daysOfWeek with * in string form ",
			daysOfWeek: "*/2",
			expected:   []int{1, 3, 5, 7},
		},
		{
			name:       "daysOfWeek with only * in string form ",
			daysOfWeek: "*",
			expected:   []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:       "daysOfWeek with comma in string form ",
			daysOfWeek: "6, 7",
			expected:   []int{10, 11},
		},
		{
			name:       "daysOfWeek with hyphen in string form ",
			daysOfWeek: "4-6",
			expected:   []int{4, 5, 6},
		},
	}

	for _, tc := range testCases {
		var testCronData CronData
		testResult := testCronData.ParseDaysOfWeek(tc.daysOfWeek)
		if len(testResult) != len(tc.expected) {
			t.Errorf("%s failed!, expected:%v, got:%v", tc.name, len(tc.expected), len(testResult))
		}
	}
}

func TestParseCommand(t *testing.T) {
	var testCronData CronData
	testName := "testing ParseCommand"
	testData := "/usr/bin/ls"
	testResult := testCronData.ParseCommand(testData)
	if len(testResult) != len(testData) {
		t.Errorf("%s failed!, expected:%v, got:%v", testName, len(testData), len(testResult))
	}
}

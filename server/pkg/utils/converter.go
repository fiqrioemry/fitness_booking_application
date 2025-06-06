package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/datatypes"
)

func StringToIntWithDefault(s string, defaultValue int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return i
}
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func IntSliceToJSON(data []int) datatypes.JSON {
	bytes, _ := json.Marshal(data)
	return datatypes.JSON(bytes)
}

func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	m := map[int]bool{}
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if !m[v] {
			return false
		}
	}
	return true
}

func JSONToIntSlice(jsonBytes []byte) []int {
	var result []int
	_ = json.Unmarshal(jsonBytes, &result)
	return result
}

func ParseJSONToIntSlice(jsonStr string) []int {
	var result []int
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Printf("Failed to parse JSON string to []int: %v\n", err)
		return []int{}
	}
	return result
}

func ParseDate(dateStr string) (time.Time, error) {
	if t, err := time.Parse("2006-01-02", dateStr); err == nil {
		return t.UTC(), nil
	}
	if t, err := time.Parse(time.RFC3339, dateStr); err == nil {
		return t.UTC(), nil
	}

	return time.Time{}, fmt.Errorf("invalid date, format must be YYYY-MM-DD or ISO 8601")
}

func GetTaxRate() float64 {
	val := os.Getenv("PAYMENT_TAX_RATE")
	if val == "" {
		return 0.05
	}
	rate, err := strconv.ParseFloat(val, 64)
	if err != nil || rate < 0 {
		return 0.05
	}
	return rate
}

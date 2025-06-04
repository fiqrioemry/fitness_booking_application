package utils

import (
	"encoding/json"
	"fmt"
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

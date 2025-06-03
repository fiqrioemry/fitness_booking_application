package utils

import (
	"encoding/json"
	"fmt"
	"strconv"

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

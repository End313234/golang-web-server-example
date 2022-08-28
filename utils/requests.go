package utils

import "encoding/json"

func ReadBody[T any](body []byte, toConvert T) T {
	json.Unmarshal(body, toConvert)
	return toConvert
}

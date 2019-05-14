package reader

import (
	"go-ff/common/service/resources/defines"
	"strconv"
)

// GetFloat64 helper method
func (r *Reader) GetFloat64(word string) float64 {
	if word == "=" {
		return 0.0
	}

	num, _ := strconv.ParseFloat(word, 64)
	return num
}

// GetUInt64 helper method
func (r *Reader) GetUInt64(word string) uint64 {
	if word == "=" {
		return 0
	}

	num, err := strconv.ParseUint(word, 10, 64)
	if err != nil {
		define, ok := defines.List[word]
		if ok {
			return (uint64)(define)
		}
	}

	return num
}

// GetUInt32 helper method
func (r *Reader) GetUInt32(word string) uint32 {
	if word == "=" {
		return 0
	}

	num, err := strconv.ParseUint(word, 10, 32)
	if err != nil {
		define, ok := defines.List[word]
		if ok {
			return (uint32)(define)
		}
	}

	return (uint32)(num)
}

// GetUInt16 helper method
func (r *Reader) GetUInt16(word string) uint16 {
	if word == "=" {
		return 0
	}

	num, err := strconv.ParseUint(word, 10, 16)
	if err != nil {
		define, ok := defines.List[word]
		if ok {
			return (uint16)(define)
		}
	}

	return (uint16)(num)
}

// GetUInt8 helper method
func (r *Reader) GetUInt8(word string) uint8 {
	if word == "=" {
		return 0
	}

	num, err := strconv.ParseUint(word, 10, 8)
	if err != nil {
		define, ok := defines.List[word]
		if ok {
			return (uint8)(define)
		}
	}

	return (uint8)(num)
}

// GetInt64 helper method
func (r *Reader) GetInt64(word string) int64 {
	if word == "=" {
		return 0
	}

	num, err := strconv.ParseInt(word, 10, 64)
	if err != nil {
		define, ok := defines.List[word]
		if ok {
			return define
		}
	}

	return num
}

// GetInt32 helper method
func (r *Reader) GetInt32(word string) int32 {
	if word == "=" {
		return 0
	}

	num, err := strconv.ParseInt(word, 10, 32)
	if err != nil {
		define, ok := defines.List[word]
		if ok {
			return (int32)(define)
		}
	}

	return (int32)(num)
}

// GetInt16 helper method
func (r *Reader) GetInt16(word string) int16 {
	if word == "=" {
		return 0
	}

	num, err := strconv.ParseInt(word, 10, 16)
	if err != nil {
		define, ok := defines.List[word]
		if ok {
			return (int16)(define)
		}
	}

	return (int16)(num)
}

// GetBool helper method
// It also read 'TRUE' as true, 'FALSE' as false
// n <= 0 as false, 0 < n as true
func (r *Reader) GetBool(word string) bool {
	if word == "TRUE" || word == "true" {
		return true
	}
	if word == "FALSE" || word == "false" {
		return false
	}

	num, _ := strconv.ParseInt(word, 10, 16)
	return num > 0
}

package lib

import "strconv"

func StringToUint32(number string) (uint32, error) {
	conv, err := strconv.Atoi(number)
	if err == nil {
		return uint32(conv), nil
	}
	return 0, err
}

func StringToInt32(number string) (int32, error) {
	conv, err := strconv.Atoi(number)
	if err == nil {
		return int32(conv), nil
	}
	return 0, err
}

func StringToUint16(number string) (uint16, error) {
	conv, err := strconv.Atoi(number)
	if err == nil {
		return uint16(conv), nil
	}
	return 0, err
}

func StringToUint8(number string) (uint8, error) {
	conv, err := strconv.Atoi(number)
	if err == nil {
		return uint8(conv), nil
	}
	return 0, err
}

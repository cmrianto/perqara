package helper

import (
	"regexp"
	"strings"
)

func FormatPhoneNumber(phone string) (string, error) {
	notNumber, err := regexp.Compile(`\D`)
	if err != nil {
		return "", err
	}

	zeros, err := regexp.Compile(`^0+`)
	if err != nil {
		return "", err
	}

	dialCodes, err := regexp.Compile(`^\+?62`)
	if err != nil {
		return "", err
	}

	phone = notNumber.ReplaceAllString(phone, "")
	phone = zeros.ReplaceAllString(phone, "")
	phone = dialCodes.ReplaceAllString(phone, "")

	return phone, nil
}

func GetMessageFromFormattedErr(in string) string {
	var msg string
	splittedString := strings.Split(in, ":")

	if len(splittedString) > 0 {
		msg = strings.TrimSpace(splittedString[len(splittedString)-1])
	}

	return msg
}

func RemoveDuplicateSpace(keyword string) (string, error) {
	space, err := regexp.Compile(`\s+`)
	if err != nil {
		return "", err
	}
	keyword = space.ReplaceAllString(keyword, " ")
	return keyword, nil
}

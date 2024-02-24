package helper

import "regexp"

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func RemoveDuplicateInt64(in []int64) []int64 {
	allKeys := make(map[int64]bool)
	list := []int64{}
	for _, item := range in {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func IsHaveSpace(in string) bool {
	return regexp.MustCompile(`\s`).MatchString(in)
}

func IsHaveTab(in string) bool {
	return regexp.MustCompile(`\t`).MatchString(in)
}

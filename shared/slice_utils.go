package shared

import "strings"

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func ToSet(list []string) map[string]bool {
	set := make(map[string]bool, len(list))
	for _, item := range list {
		set[strings.ToLower(item)] = true
	}
	return set
}

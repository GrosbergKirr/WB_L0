package tools

import "project0/models"

func Checker(cache map[string]models.Order, uid string) bool {
	var keys []string
	for i := range cache {
		keys = append(keys, i)
	}

	return IndexFunc(keys, uid)
}
func IndexFunc(s []string, n string) bool {
	for i := range s {
		if s[i] == n {
			return false
		}
	}
	return true
}

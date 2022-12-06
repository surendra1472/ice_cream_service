package helpers

import (
	"ic-service/app/model/request"
)

func GetArrayString(cusArrayString request.CusArrayString, existing []string) []string {
	if cusArrayString.Set {
		if cusArrayString.Value != nil {
			return *cusArrayString.Value
		}
	}
	return existing
}

func GetString(cusString request.CusString, existing string) string {
	if cusString.Set {
		if cusString.Value != nil {
			return *cusString.Value
		}
	}
	return existing
}

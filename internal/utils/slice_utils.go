package utils

import "reflect"

func SearchSliceById[T any](slice *[]T, id string) (*T, bool) {
	for _, item := range *slice {
		if v, ok := getIdFromItem(item); ok && v == id {
			return &item, true
		}
	}
	return nil, false
}

func getIdFromItem(item interface{}) (string, bool) {
	val := reflect.ValueOf(item)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	field := val.FieldByName("Id")
	if field.IsValid() && field.Kind() == reflect.String {
		return field.String(), true
	}

	return "", false
}

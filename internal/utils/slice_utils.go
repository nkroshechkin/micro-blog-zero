package utils

import "reflect"

func SliceSearchById[T any](slice []T, id string) (*T, bool) {
	for i, item := range slice {
		if v, ok := getIdFromItem(item); ok && v == id {
			return &slice[i], true
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

func SliceFilter[T any](slice []T, callback func(item T) bool) []T {
	filteredSlice := make([]T, len(slice), cap(slice))
	for _, item := range slice {
		if callback(item) {
			filteredSlice = append(filteredSlice, item)
		}
	}

	return filteredSlice
}

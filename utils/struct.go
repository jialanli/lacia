package lacia

import (
	"reflect"
)

func CopyStructByFieldGraceful(from, to any) {
	toValue, fromValue := reflect.ValueOf(to), reflect.ValueOf(from)

	if fromValue.Kind() != reflect.Ptr || toValue.Kind() != reflect.Ptr {
		return
	}

	toElem, fromElem := toValue.Elem(), fromValue.Elem()
	for i := 0; i < toElem.NumField(); i++ {
		toField := toElem.Type().Field(i)
		fromField, ok := fromElem.Type().FieldByName(toField.Name)
		if ok && fromField.Type == toField.Type {
			toElem.Field(i).Set(fromElem.FieldByName(toField.Name))
		}
	}

	return
}

func CopyStructByTagGraceful(from, to any) {
	toValue, fromValue := reflect.ValueOf(to), reflect.ValueOf(from)

	if fromValue.Kind() != reflect.Ptr || toValue.Kind() != reflect.Ptr {
		return
	}

	toElem, fromElem := toValue.Elem(), fromValue.Elem()

	for i := 0; i < toElem.NumField(); i++ {
		toTag, toFieldType := toElem.Type().Field(i).Tag, toElem.Type().Field(i).Type
		for j := 0; j < fromElem.NumField(); j++ {
			fromTag, fromFieldType := fromElem.Type().Field(j).Tag, fromElem.Type().Field(j).Type
			if fromTag == toTag && toFieldType == fromFieldType {
				toElem.Field(i).Set(fromElem.Field(j))
				break
			}
		}
	}

	return
}

func CopyStructByCustomTagGraceful(from, to any, customTag string) {
	toValue, fromValue := reflect.ValueOf(to), reflect.ValueOf(from)

	if fromValue.Kind() != reflect.Ptr || toValue.Kind() != reflect.Ptr {
		return
	}

	toElem, fromElem := toValue.Elem(), fromValue.Elem()

	for i := 0; i < toElem.NumField(); i++ {
		toTag, toFieldType := toElem.Type().Field(i).Tag, toElem.Type().Field(i).Type
		if len(toTag) == 0 {
			continue
		}

		for j := 0; j < fromElem.NumField(); j++ {
			fromTag, fromFieldType := fromElem.Type().Field(j).Tag, fromElem.Type().Field(j).Type
			if len(fromTag) == 0 {
				continue
			}

			if fromTag == toTag && toFieldType == fromFieldType {
				toElem.Field(i).Set(fromElem.Field(j))
				break
			}
		}
	}

	return
}

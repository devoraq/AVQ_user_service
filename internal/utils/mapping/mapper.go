package mapping

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"time"
)

const tag = "map"

func MappingStructDAO(dst, src any) error {
	if reflect.TypeOf(src).Kind() != reflect.Ptr {
		return errors.New("src must be a pointer")
	}

	dstVal := reflect.Indirect(reflect.ValueOf(dst))
	srcVal := reflect.ValueOf(src).Elem()

	if dstVal.Kind() != reflect.Struct {
		fmt.Println(dstVal.Kind())
		return errors.New("unsupported type")
	}

	srcType := srcVal.Type()

	for i := 0; i < srcType.NumField(); i++ {
		srcField := srcType.Field(i)
		dstField, ok := dstVal.Type().FieldByName(srcField.Name)
		if !ok {
			continue
		}

		if srcField.Tag.Get(tag) != dstField.Tag.Get(tag) {
			continue
		}

		srcValue := srcVal.Field(i)
		dstValue := dstVal.Field(i)

		if !srcValue.CanSet() || !dstValue.IsValid() {
			continue
		}

		switch srcField.Type {
		case reflect.TypeOf(sql.NullString{}):
			if dstValue.Kind() == reflect.String {
				newVal := toNullString(dstValue.String())
				srcValue.Set(reflect.ValueOf(newVal))
			}

		case reflect.TypeOf(sql.NullTime{}):
			if dstValue.Type() == reflect.TypeOf(time.Time{}) {
				t, ok := dstValue.Interface().(time.Time)
				if ok {
					newVal := toNullTimestamp(t)
					srcValue.Set(reflect.ValueOf(newVal))
				}
			}

		case reflect.TypeOf(sql.NullBool{}):
			if dstValue.Kind() == reflect.Bool {
				newVal := toNullBool(dstValue.Bool())
				srcValue.Set(reflect.ValueOf(newVal))
			}

		case reflect.TypeOf(sql.NullInt16{}):
			if dstValue.Kind() == reflect.Int16 {
				newVal := toNullInt16(dstValue.Interface().(int16))
				srcValue.Set(reflect.ValueOf(newVal))
			}

		case reflect.TypeOf(sql.NullInt32{}):
			if dstValue.Kind() == reflect.Int32 {
				newVal := toNullInt32(dstValue.Interface().(int32))
				srcValue.Set(reflect.ValueOf(newVal))
			}

		case reflect.TypeOf(sql.NullInt64{}):
			if dstValue.Kind() == reflect.Int64 {
				newVal := toNullInt64(dstValue.Interface().(int64))
				srcValue.Set(reflect.ValueOf(newVal))
			}
		}
	}

	return nil
}

func MappingStruct(dst, src any) error {
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src).Elem()

	if dstVal.Kind() != reflect.Struct && srcVal.Kind() != reflect.Struct {
		return errors.New("unsupported type")
	}

	srcType := srcVal.Type()

	for i := 0; i < dstVal.Type().NumField(); i++ {
		srcField := srcType.Field(i)

		dstField, ok := dstVal.Type().FieldByName(srcField.Name)
		if !ok {
			continue
		}

		if srcField.Tag.Get(tag) == dstField.Tag.Get(tag) {
			srcValue := srcVal.Field(i)
			dstValue := dstVal.Field(i)

			if !dstValue.IsValid() && !dstValue.CanSet() {
				continue
			}
			if srcValue.Type() != dstValue.Type() {
				continue
			}

			srcValue.Set(dstValue)
		}
	}

	return nil
}

package mapping

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const tag = "map"

func MapStructDAO(src, dst any) error {
	srcVal, err := validateSrc(src)
	if err != nil {
		return err
	}

	dstVal, err := validateDst(dst)
	if err != nil {
		return err
	}

	srcType := srcVal.Type()

	for i := 0; i < srcType.NumField(); i++ {
		srcField := srcType.Field(i)
		dstField, ok := dstVal.Type().FieldByName(srcField.Name)
		if !ok {
			continue
		}

		srcValue := srcVal.Field(i)
		dstValue := dstVal.FieldByName(srcField.Name)

		srcTag := srcField.Tag.Get(tag)
		dstTag := dstField.Tag.Get(tag)

		if srcTag != dstTag {
			continue
		}
		if !dstValue.CanSet() {
			continue
		}

		switch dstField.Type {
		case reflect.TypeOf(sql.NullString{}):
			val := srcValue.Interface().(string)
			dstValue.Set(reflect.ValueOf(toNullString(val)))
		case reflect.TypeOf(sql.NullTime{}):
			val := srcValue.Interface().(time.Time)
			dstValue.Set(reflect.ValueOf(toNullTimestamp(val)))
		case reflect.TypeOf(sql.NullBool{}):
			val := srcValue.Interface().(bool)
			dstValue.Set(reflect.ValueOf(toNullBool(val)))
		case reflect.TypeOf(sql.NullInt16{}):
			val := srcValue.Interface().(int16)
			dstValue.Set(reflect.ValueOf(toNullInt16(val)))
		case reflect.TypeOf(sql.NullInt32{}):
			val := srcValue.Interface().(int32)
			dstValue.Set(reflect.ValueOf(toNullInt32(val)))
		case reflect.TypeOf(sql.NullInt64{}):
			val := srcValue.Interface().(int64)
			dstValue.Set(reflect.ValueOf(toNullInt64(val)))
		default:
			if dstValue.Type() == srcValue.Type() {
				dstValue.Set(srcValue)
			}
		}
	}

	return nil
}

func MapStruct(src, dst any) error {
	srcVal, err := validateSrc(src)
	if err != nil {
		return err
	}
	srcType := srcVal.Type()

	dstVal, err := validateDst(dst)
	if err != nil {
		return err
	}

	for i := 0; i < srcType.NumField(); i++ {
		srcField := srcType.Field(i)

		srcValue := srcVal.Field(i)
		dstValue := dstVal.FieldByName(srcField.Name)

		if dstValue.CanSet() {
			switch srcValue.Type() {
			case reflect.TypeOf(&timestamppb.Timestamp{}):
				timepb := srcValue.Interface().(*timestamppb.Timestamp)
				dstValue.Set(reflect.ValueOf(timepb.AsTime()))
			default:
				dstValue.Set(srcValue)
			}
		}
	}

	return nil
}

func validateSrc(src any) (reflect.Value, error) {
	srcVal := reflect.Indirect(reflect.ValueOf(src))
	if srcVal.Kind() != reflect.Struct {
		return srcVal, errors.New("src must be a struct or pointer to struct")
	}

	return srcVal, nil
}

func validateDst(dst any) (reflect.Value, error) {
	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Ptr {
		return dstVal, fmt.Errorf("dst must be a pointer, got %s", dstVal.Kind())
	}
	if dstVal.Type().Elem().Kind() != reflect.Struct {
		return dstVal, fmt.Errorf("")
	}
	if dstVal.IsNil() {
		return dstVal, fmt.Errorf("dst cannot be nil")
	}

	return dstVal.Elem(), nil
}

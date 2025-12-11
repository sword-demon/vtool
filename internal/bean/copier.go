package bean

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

// Converter 类型转换函数类型
type Converter func(srcValue reflect.Value, dstType reflect.Type) (reflect.Value, error)

// Options 复制选项
type Options struct {
	Converter    Converter
	IgnoreFields []string
	DeepCopy     bool
	IgnoreEmpty  bool
}

// 默认选项
var defaultOptions = Options{
	DeepCopy:    false,
	IgnoreEmpty: false,
	Converter:   nil,
}

// Copy 复制结构体
// src 源结构体指针，dst 目标结构体指针
func Copy(src, dst interface{}, opts ...Options) error {
	if src == nil || dst == nil {
		return errors.New("source and destination cannot be nil")
	}

	// 合并选项
	options := defaultOptions
	if len(opts) > 0 {
		options = opts[0]
	}

	// 获取反射值
	srcVal, dstVal, err := validateAndGetValues(src, dst)
	if err != nil {
		return err
	}

	// 执行复制
	return copyValue(srcVal, dstVal, options)
}

// CopyWithoutNil 复制结构体，跳过nil指针
func CopyWithoutNil(src, dst interface{}) error {
	return Copy(src, dst, Options{
		DeepCopy:    true,
		IgnoreEmpty: true,
	})
}

// DeepCopy 深度复制结构体
func DeepCopy(src, dst interface{}) error {
	return Copy(src, dst, Options{
		DeepCopy: true,
	})
}

// validateAndGetValues 验证输入并获取反射值
func validateAndGetValues(src, dst interface{}) (reflect.Value, reflect.Value, error) {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	// 检查指针
	if srcVal.Kind() != reflect.Ptr || dstVal.Kind() != reflect.Ptr {
		return reflect.Value{}, reflect.Value{}, errors.New("source and destination must be pointers")
	}

	// 解引用
	for srcVal.Kind() == reflect.Ptr {
		if srcVal.IsNil() {
			return reflect.Value{}, reflect.Value{}, errors.New("source pointer is nil")
		}
		srcVal = srcVal.Elem()
	}

	for dstVal.Kind() == reflect.Ptr {
		if dstVal.IsNil() {
			// 创建目标对象
			dstVal.Set(reflect.New(dstVal.Type().Elem()))
		}
		dstVal = dstVal.Elem()
	}

	// 检查类型
	if srcVal.Kind() != reflect.Struct || dstVal.Kind() != reflect.Struct {
		return reflect.Value{}, reflect.Value{}, errors.New("source and destination must be structs")
	}

	return srcVal, dstVal, nil
}

// copyValue 复制值
func copyValue(srcVal, dstVal reflect.Value, options Options) error {
	srcType := srcVal.Type()

	// 创建字段映射（字段名 -> 字段信息）
	fieldMap := make(map[string]reflect.StructField)
	for i := 0; i < srcType.NumField(); i++ {
		field := srcType.Field(i)
		// 跳过不可导出字段
		if field.PkgPath != "" {
			continue
		}
		// 跳过忽略字段
		if contains(options.IgnoreFields, field.Name) {
			continue
		}
		fieldMap[field.Name] = field
	}

	// 复制字段
	for i := 0; i < dstVal.NumField(); i++ {
		dstField := dstVal.Type().Field(i)
		dstFieldValue := dstVal.Field(i)

		// 跳过不可导出字段
		if dstField.PkgPath != "" {
			continue
		}

		// 查找对应的源字段
		if srcField, ok := fieldMap[dstField.Name]; ok {
			srcFieldValue := srcVal.FieldByIndex(srcField.Index)

			// 检查是否可以设置
			if !dstFieldValue.CanSet() {
				continue
			}

			// 检查是否忽略空值
			if options.IgnoreEmpty && isZeroValue(srcFieldValue) {
				continue
			}

			// 执行复制
			if err := copyField(srcFieldValue, dstFieldValue, options); err != nil {
				return fmt.Errorf("error copying field %s: %w", dstField.Name, err)
			}
		}
	}

	return nil
}

// copyField 复制单个字段
func copyField(srcVal, dstVal reflect.Value, options Options) error {
	// 如果类型相同，直接复制
	if srcVal.Type() == dstVal.Type() {
		// 深度复制
		if options.DeepCopy {
			return deepCopyValue(srcVal, dstVal)
		}
		dstVal.Set(srcVal)
		return nil
	}

	// 使用转换器
	if options.Converter != nil {
		converted, err := options.Converter(srcVal, dstVal.Type())
		if err != nil {
			return err
		}
		dstVal.Set(converted)
		return nil
	}

	// 尝试类型转换
	return tryConvert(srcVal, dstVal)
}

// deepCopyValue 深度复制值
func deepCopyValue(srcVal, dstVal reflect.Value) error {
	switch srcVal.Kind() {
	case reflect.Ptr:
		if srcVal.IsNil() {
			return nil
		}
		if dstVal.IsNil() {
			dstVal.Set(reflect.New(srcVal.Elem().Type()))
		}
		return deepCopyValue(srcVal.Elem(), dstVal.Elem())

	case reflect.Struct:
		for i := 0; i < srcVal.NumField(); i++ {
			if err := deepCopyValue(srcVal.Field(i), dstVal.Field(i)); err != nil {
				return err
			}
		}
		return nil

	case reflect.Slice:
		dstVal.Set(reflect.MakeSlice(srcVal.Type(), srcVal.Len(), srcVal.Len()))
		for i := 0; i < srcVal.Len(); i++ {
			if err := deepCopyValue(srcVal.Index(i), dstVal.Index(i)); err != nil {
				return err
			}
		}
		return nil

	case reflect.Map:
		dstVal.Set(reflect.MakeMap(srcVal.Type()))
		for _, key := range srcVal.MapKeys() {
			newKey := reflect.New(key.Type()).Elem()
			if err := deepCopyValue(key, newKey); err != nil {
				return err
			}
			newValue := reflect.New(srcVal.MapIndex(key).Type()).Elem()
			if err := deepCopyValue(srcVal.MapIndex(key), newValue); err != nil {
				return err
			}
			dstVal.SetMapIndex(newKey, newValue)
		}
		return nil

	default:
		dstVal.Set(srcVal)
		return nil
	}
}

// tryConvert 尝试类型转换
func tryConvert(srcVal, dstVal reflect.Value) error {
	// 数字类型转换
	if isNumeric(srcVal.Type()) && isNumeric(dstVal.Type()) {
		return convertNumeric(srcVal, dstVal)
	}

	// 字符串转换
	if srcVal.Type() == reflect.TypeOf(time.Time{}) && dstVal.Type() == reflect.TypeOf("") {
		t := srcVal.Interface().(time.Time)
		dstVal.SetString(t.Format(time.RFC3339))
		return nil
	}

	if srcVal.Type() == reflect.TypeOf("") && dstVal.Type() == reflect.TypeOf(time.Time{}) {
		str := srcVal.String()
		t, err := time.Parse(time.RFC3339, str)
		if err != nil {
			return err
		}
		dstVal.Set(reflect.ValueOf(t))
		return nil
	}

	return fmt.Errorf("cannot convert from %s to %s", srcVal.Type(), dstVal.Type())
}

// isNumeric 检查是否为数字类型
func isNumeric(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return true
	}
	return false
}

// convertNumeric 转换数字类型
func convertNumeric(srcVal, dstVal reflect.Value) error {
	var dstValue int64
	isFloat := false

	switch srcVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		dstValue = srcVal.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		dstValue = int64(srcVal.Uint())
	case reflect.Float32, reflect.Float64:
		isFloat = true
	}

	switch dstVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if isFloat {
			dstVal.SetInt(int64(srcVal.Float()))
		} else {
			dstVal.SetInt(dstValue)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if isFloat {
			dstVal.SetUint(uint64(srcVal.Float()))
		} else {
			dstVal.SetUint(uint64(dstValue))
		}
	case reflect.Float32, reflect.Float64:
		if isFloat {
			dstVal.SetFloat(srcVal.Float())
		} else {
			dstVal.SetFloat(float64(dstValue))
		}
	}

	return nil
}

// isZeroValue 检查是否为零值
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Complex64, reflect.Complex128:
		return v.Complex() == 0
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isZeroValue(v.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	case reflect.String:
		return v.Len() == 0
	}
	return false
}

// contains 检查切片是否包含元素
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

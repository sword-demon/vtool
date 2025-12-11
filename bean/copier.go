package bean

import "github.com/sword-demon/vtool/internal/bean"

// Converter 类型转换函数类型
type Converter = bean.Converter

// Options 复制选项
type Options = bean.Options

// Copy 复制结构体
// src 源结构体指针，dst 目标结构体指针
func Copy(src, dst interface{}, opts ...Options) error {
	return bean.Copy(src, dst, opts...)
}

// CopyWithoutNil 复制结构体，跳过nil指针
func CopyWithoutNil(src, dst interface{}) error {
	return bean.CopyWithoutNil(src, dst)
}

// DeepCopy 深度复制结构体
func DeepCopy(src, dst interface{}) error {
	return bean.DeepCopy(src, dst)
}

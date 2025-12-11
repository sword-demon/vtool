package bean

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SourceStruct struct {
	Name  string
	Age   int
	Email string
}

type DestStruct struct {
	Name  string
	Age   int
	Email string
}

type PartialStruct struct {
	Name string
	City string
}

type NumericSource struct {
	IntValue    int
	Int64Value  int64
	FloatValue  float64
	StringValue string
}

type NumericDest struct {
	IntValue    int64
	FloatValue  float32
	StringValue string
	IntValue2   int
}

type NestedSource struct {
	ID   int
	Name string
	Org  *Organization
}

type Organization struct {
	Name string
	Code string
}

type NestedDest struct {
	ID   int
	Name string
	Org  *Organization
}

func TestCopy(t *testing.T) {
	t.Run("基本复制", func(t *testing.T) {
		src := &SourceStruct{
			Name:  "John",
			Age:   30,
			Email: "john@example.com",
		}
		dst := &DestStruct{}

		err := Copy(src, dst)
		assert.NoError(t, err)
		assert.Equal(t, "John", dst.Name)
		assert.Equal(t, 30, dst.Age)
		assert.Equal(t, "john@example.com", dst.Email)
	})

	t.Run("部分字段复制", func(t *testing.T) {
		src := &SourceStruct{
			Name:  "John",
			Age:   30,
			Email: "john@example.com",
		}
		dst := &PartialStruct{}

		err := Copy(src, dst)
		assert.NoError(t, err)
		assert.Equal(t, "John", dst.Name)
		assert.Equal(t, "", dst.City)
	})

	t.Run("不同类型转换", func(t *testing.T) {
		src := &NumericSource{
			IntValue:   100,
			Int64Value: 999,
			FloatValue: 3.14,
			StringValue: "test",
		}
		dst := &NumericDest{}

		err := Copy(src, dst)
		assert.NoError(t, err)
		assert.Equal(t, int64(100), dst.IntValue)
		assert.Equal(t, float32(3.14), dst.FloatValue)
		assert.Equal(t, "test", dst.StringValue)
	})

	t.Run("忽略空值", func(t *testing.T) {
		src := &SourceStruct{
			Name:  "John",
			Age:   0, // 零值
			Email: "john@example.com",
		}
		dst := &SourceStruct{
			Name:  "Old",
			Age:   50,
			Email: "old@example.com",
		}

		err := Copy(src, dst, Options{IgnoreEmpty: true})
		assert.NoError(t, err)
		assert.Equal(t, "John", dst.Name)
		assert.Equal(t, 50, dst.Age)     // 保持不变
		assert.Equal(t, "john@example.com", dst.Email)
	})

	t.Run("深度复制", func(t *testing.T) {
		org := &Organization{
			Name: "ACME",
			Code: "123",
		}
		src := &NestedSource{
			ID:   1,
			Name: "John",
			Org:  org,
		}
		dst := &NestedDest{}

		err := DeepCopy(src, dst)
		assert.NoError(t, err)
		assert.Equal(t, 1, dst.ID)
		assert.Equal(t, "John", dst.Name)
		assert.NotNil(t, dst.Org)
		assert.Equal(t, "ACME", dst.Org.Name)
		assert.Equal(t, "123", dst.Org.Code)

		// 修改源对象不应该影响目标对象
		org.Name = "Changed"
		assert.Equal(t, "ACME", dst.Org.Name) // 保持不变
	})

	t.Run("错误情况", func(t *testing.T) {
		// nil源
		err := Copy(nil, &DestStruct{})
		assert.Error(t, err)

		// nil目标
		err = Copy(&SourceStruct{}, nil)
		assert.Error(t, err)

		// 非指针
		err = Copy(SourceStruct{}, &DestStruct{})
		assert.Error(t, err)

		// 非结构体
		var src int
		var dst int
		err = Copy(&src, &dst)
		assert.Error(t, err)
	})

	t.Run("时间转换", func(t *testing.T) {
		type SrcWithTime struct {
			Name string
			Time time.Time
		}
		type DestWithTime struct {
			Name string
			Time string
		}

		now := time.Now()
		src := &SrcWithTime{
			Name: "Test",
			Time: now,
		}
		dst := &DestWithTime{}

		err := Copy(src, dst)
		assert.NoError(t, err)
		assert.Equal(t, "Test", dst.Name)
		assert.NotEmpty(t, dst.Time)
	})

	t.Run("自定义转换器", func(t *testing.T) {
		type Src struct {
			Value string
		}
		type Dest struct {
			Value int
		}

		src := &Src{Value: "123"}
		dst := &Dest{}

		converter := func(srcValue reflect.Value, dstType reflect.Type) (reflect.Value, error) {
			if srcValue.Type() == reflect.TypeOf("") && dstType.Kind() == reflect.Int {
				intVal := 0
				fmt.Sscanf(srcValue.String(), "%d", &intVal)
				return reflect.ValueOf(intVal), nil
			}
			return srcValue, nil
		}

		err := Copy(src, dst, Options{Converter: converter})
		assert.NoError(t, err)
		assert.Equal(t, 123, dst.Value)
	})

	t.Run("忽略字段", func(t *testing.T) {
		src := &SourceStruct{
			Name:  "John",
			Age:   30,
			Email: "john@example.com",
		}
		dst := &SourceStruct{
			Name:  "OldName",
			Age:   50,
			Email: "old@example.com",
		}

		err := Copy(src, dst, Options{IgnoreFields: []string{"Age"}})
		assert.NoError(t, err)
		assert.Equal(t, "John", dst.Name)
		assert.Equal(t, 50, dst.Age)     // 保持不变
		assert.Equal(t, "john@example.com", dst.Email)
	})
}
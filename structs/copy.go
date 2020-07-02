package structs

import (
	"fmt"
	"os"
	"reflect"
)

//把a struct的值，赋值到b struct对应的一模一样的key上

func CopyStruct(sourceStruct interface{}, dstStruct interface{}) {
	defer func() {
		if co := recover(); co != nil {
			str := "SetStruct error:发生panic :" + fmt.Sprint(co)
			fmt.Println(str)
			os.Exit(1)
		}
	}()
	v := reflect.ValueOf(dstStruct)
	sourceV := reflect.ValueOf(sourceStruct)
	if sourceV.Kind() == reflect.Ptr {
		sourceV = sourceV.Elem()
	}
	if sourceV.Kind() != reflect.Struct {
		panic("CopyStruct: sourceStruct must be struct or point of struct")
	}
	sourceT := sourceV.Type()
	switch v.Kind() {
	case reflect.Ptr:
		t := v.Type().Elem()

		for i := 0; i < sourceV.NumField(); i++ {
			sourceFieldName := sourceT.Field(i).Name
			sourceType := sourceT.Field(i).Type
			sourceValue := sourceV.Field(i)

			for i := 0; i < v.Elem().NumField(); i++ {
				fieldName := t.Field(i).Name
				vType := t.Field(i).Type

				if fmt.Sprint(vType) == fmt.Sprint(sourceType) && fieldName == sourceFieldName {
					v.Elem().Field(i).Set(sourceValue)

				}
			}
		}

	default:
		panic("CopyStruct error:要传入dstStruct的是结构体指针")

	}

}

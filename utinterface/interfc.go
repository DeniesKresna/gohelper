package utinterface

import "reflect"

func IsStruct(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.Struct
}

func IsPointer(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.Pointer
}

func IsMapStringInterface(i interface{}) bool {
	_, ok := i.(map[string]interface{})

	return ok
}

func IsSlice(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.Slice
}

func IsSliceOfMapStringInterface(i interface{}) bool {
	_, ok := i.([]map[string]interface{})

	return ok
}

func IsPointerOfStruct(i interface{}) bool {
	if !IsPointer(i) {
		return false
	}

	return reflect.TypeOf(i).Elem().Kind() == reflect.Struct
}

func IsPointerOfSliceOfStruct(i interface{}) bool {
	if !IsPointer(i) {
		return false
	}

	iType := reflect.TypeOf(i).Elem()

	if iType.Kind() != reflect.Slice {
		return false
	}

	return iType.Elem().Kind() == reflect.Struct
}

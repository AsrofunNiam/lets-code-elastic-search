package helper

import "reflect"

func MessageDataFoundOrNot(data interface{}) string {
	s := reflect.ValueOf(data)

	if s.Kind() == reflect.Slice {
		if s.Len() > 0 {
			return "Record found"
		} else {
			return "Record not found"
		}
	}
	if s.Kind() == reflect.Struct {
		return "Record found"
	}
	if data == nil {
		return "Record not found"
	}
	panic("MessageData FoundOrNor() given parameter must be slice or struct")

}

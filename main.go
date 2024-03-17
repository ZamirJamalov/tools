package main

import (
	"errors"
	"fmt"
)

func isInt(i interface{}) bool {
	switch i.(type) {
	case int:
		return true
	default:
		return false
	}
}

func Substr(params ...interface{}) error {
	if len(params) > 3 {
		panic(errors.New("Too many parameters"))
	}
	idx := 0
	var str string
	var strP *string
	var start_pos int
	var start_posP *int
	var start_pos_org int
	var length int
	var lengthP *int
	for _, param := range params {
		idx++
		switch idx {
		case 1:
			{
				if isInt(param) == true {
					return errors.New("First param must be string type")
				} else {
					str = param.(string)
					strP = &str
					//fmt.Println(str)
				}
			}
		case 2:
			{
				if isInt(param) == false {
					return errors.New("Second param must be int type")
				} else {
					start_pos = param.(int)
					start_posP = &start_pos
					start_pos_org = start_pos
					fmt.Println(start_pos)
				}
			}
		case 3:
			{
				if isInt(param) == false {
					return errors.New("Third param must be int type")
				} else {
					length = param.(int)
					lengthP = &length
					if length < 0 {
						panic(errors.New("Length must be greater than 0"))
					}

				}
			}
		}

	}
	if strP == nil {
		panic(errors.New("String Param Required"))
	}
	if start_posP == nil {
		panic(errors.New("Start Position Param Requiered"))
	}
	if lengthP == nil {
		length = len(str)
	}

	if start_pos < 0 {
		start_pos = len(str) - -1*start_pos
		fmt.Println(start_pos)

		if length- -1*start_pos_org > 0 {
			length = len(str)
		} else {
			length = start_pos + length
		}

	} else {
		start_pos = start_pos - 1
		length = start_pos + length
	}
	fmt.Println(length)
	fmt.Println(start_pos)

	arr := []rune(str)
	fmt.Println(string(arr[start_pos:length]))
	return nil
}



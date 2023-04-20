package mypackage

import "strconv"

var Name string = "Vin"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, val := range intArr {
		strArr = append(strArr, strconv.Itoa(val))
	}
	return strArr
}

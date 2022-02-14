package main

import (
	"errors"
	"fmt"
)


func main () {
	fmt.Println(deret(2,4,5))
	fmt.Println(deret(5,8,7))
	arr := []float64{4, -7, -5, 3.3, 3, 9, 0, 10, 0.2}
	fmt.Println(sort(arr,"asc"))
	fmt.Println(sort(arr,"desc"))
	fmt.Println(sort(arr,"disc"))
	
}

// PSEUDOCODE
// BEGIN deret
//	 PASS IN first,second, and len as NUMBER
//	 SET diff TO second MINUS first
//	 STORE arr WITH []
//	 FOR i FROM 0 to len INCREMENT BY 1
//	 	IF i EQUAL 0
//	 		SET arr[i] WITH first
//	 	ELSE 
//	 		SET arr[i] WITH  arr[i MINUS 1] PLUS diff
//	 	END IF
//	 END FOR
//	 PASS OUT res
// END
func deret(first, second, len int) []int {
	diff := second - first
	var res []int 
	for i:= 0 ; i < len ; i++ {
		if i == 0 {
			res = append(res,first) 
		} else {
			res = append(res, res[i-1]+diff)
		}
	}
	return res
}

// PSEUDOCODE
// BEGIN sort
//	 PASS IN data as ARRAY OF NUMBER and action as STRING
//	 STORE arr WITH []
//	 FOR i FROM 0 to LENGTH OF data INCREMENT BY 1
//	 	SET index with i
// 		SET temp with data[i]
// 		FOR j FROM i PLUS 1 to LENGTH OF data INCREMENT BY 1
//		 	IF action EQUAL asc
//		 		IF data[i] > data[j]
// 					SET index = j
// 					SET data[i]= data[j]
// 				END IF
//		 	ELSE IF action EQUAL desc
//		 		IF data[i] < data[j]
// 					SET index = j
// 					SET data[i]= data[j]
// 				END IF
// 			ELSE
// 				PASS OUT nil, ERROR WITH MESSAGE "action must be filled with asc or desc"
//		 	END IF
// 		END FOR
// 		ST data[i] WITH data[index]
// 		SET data[index] WITH temp
//	 END FOR
//	 PASS OUT res , nil
// END
func sort(data []float64, action string) ([]float64, error) {
	for i:= 0 ; i < len(data) ; i++ {
		index := i
		temp := data[i]
		for j:= i+1 ; j < len(data) ; j++ {
			if action == "asc" {
				if data[i] > data[j] {
					index = j
					data[i] = data[j]
				}
			} else if action == "desc" {
				if data[i] < data[j] {
					index = j
				}
			} else {
				return nil, errors.New("action must be filled with asc or desc")
			}
		}
		data[i] = data[index]
		data[index] = temp
	}
	return data, nil
}
package helper

import "fmt"

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
func CheckDBErr(e error) {
	if e != nil {
		fmt.Println("DB Connection Failed")
		panic(e)
	} else {
		fmt.Println("DB Connection Success")
	}
}

func CheckEnvErr(e error) {
	if e != nil {
		fmt.Println("File environment failed")
	} else {
		fmt.Println("File environment successful")
	}
}

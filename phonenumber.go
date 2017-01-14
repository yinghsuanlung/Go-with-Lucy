package main

import "fmt"

func main() {
	phoneNum := map[string]string{"Lucy": "0958083802", "Jerry": "0972858987", "Jessie": "0952259611"}
	jessieNum, yes := phoneNum["Jessie"]
	if yes {
		fmt.Println("Jessie's number is:", jessieNum)
	} else {
		fmt.Println("We dont have Jessie's number")
	}
	delete(phoneNum, "Jessie")
}

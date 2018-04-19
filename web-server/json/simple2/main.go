package main

import (
	"encoding/json"
	"fmt"
)

type statusCode struct {
	Code        int    `json:"Code"`
	Description string `json:"Descrip"`
}

type statusCodes []statusCode

func main() {
	rcvd := `[{"Code":200,"Descrip":"StatusOK"},{"Code":301,"Descrip":"StatusMovedPermantenly"},{"Code":302,"Descrip":"StatusFound"}]`

	var data statusCodes
	if err := json.Unmarshal([]byte(rcvd), &data); err != nil {
		fmt.Println("Error found: ", err)
	}
	fmt.Println(data)

	for _, v := range data {
		fmt.Printf("Code:\t %v, Description:\t %v\n", v.Code, v.Description)
	}

}

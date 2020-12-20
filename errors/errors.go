package errors

import "fmt"

type Error struct {
	Status int `json:"status"`
	Message string `json:"error_message"`
}

func GetFileError(err error){
	if err != nil{
		fmt.Println("File Error:", err)
		return
	}	
}

func GetMultiPartFormError(err error){
	if err != nil{
		fmt.Println("Multipart Form File Error:", err)
		return
	}
}

func GetDirectoryError(err error){
	if err != nil{
		fmt.Println("Directory Error:", err)
		return
	}
}

func GetError(err error){
	if err != nil{
		fmt.Println(err)
		return
	}
}
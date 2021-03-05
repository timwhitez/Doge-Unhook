package main

import (
	"log"
)

func main(){
	err := RefreshPE(`c:\windows\system32\ntdll.dll`)
	if err != nil {
		log.Println("RefreshPE failed:", err)
	}
	return
}
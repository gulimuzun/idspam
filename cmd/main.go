package main

import (
	"fmt"
	"idspam/internal/case_handle"
	"os"
	"path/filepath"
)

func main(){
	if len(os.Args)<2{
		fmt.Println("UsAge:idspam <command>")
		fmt.Println("for more information,use 'idspam help'")
		return 
	}
	cmd := os.Args[1]
	switch cmd{
	case "train":
		if len(os.Args)>5{
			fmt.Println("Wrong command,too much command-line argument")
		}else if len(os.Args)==5{
			err:=case_handle.CaseTrain(os.Args[2],os.Args[3],os.Args[4])
			if err!=nil{
				case_handle.HandleError(err)
			}
		}else if len(os.Args)<=3{
			fmt.Println("Wrong command,lacking command-line argument")
		}else{
			pathWd,err:=os.Getwd()
			if err!=nil {
				case_handle.HandleError(err)
			}
			pathTarget:=filepath.Join(pathWd,"model.gob")
			err=case_handle.CaseTrain(os.Args[2],os.Args[3],pathTarget)
			if err!=nil{
				case_handle.HandleError(err)
			}
		}
	case "identify":
		if len(os.Args)>4{
			fmt.Println("Wrong command,too much command-line argument")
		}else if len(os.Args)<=3{
			fmt.Println("Wrong command,lacking command-line argument")
		}else{
			err :=case_handle.CaseIdentify(os.Args[2],os.Args[3])
			if err!=nil{
				case_handle.HandleError(err)
			}
		}
	case "help":
		err:=case_handle.CaseHelp()
		if err!=nil{
			case_handle.HandleError(err)
		}
	default:
		fmt.Println("Unknown command")
	}
}
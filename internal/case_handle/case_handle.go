package case_handle

import (
	"fmt"
	"idspam/internal/train"
	"os"
)
func CaseTrain(pathTrainHamSet string,pathTrainSpamSet string,pathGob string)error{
	err :=train.OutputTrainFile(train.TrainAndCalc(pathTrainHamSet,pathTrainSpamSet),pathGob)
	if err != nil {
		return err
	}
	fmt.Println("successfully train")
	return nil
}
func CaseIdentify(pathModel string,pathTarget string)error{
	model,err:=train.LoadTrainFile(pathModel)
	if err != nil{
		return err
	}
	if train.PredictIsHam(model,pathTarget){
		fmt.Println("Ham")
	}else{
		fmt.Println("Spam")
	}
	return nil
}
func CaseHelp()error{
	fmt.Print(`idspam - Naive Bayes spam/ham classifier

USAGE:
  idspam <command> [arguments]

COMMANDS:
  train      Train a model from ham/spam datasets and save as a .gob file
  identify   Identify whether a target file is ham or spam using a trained model
  help       Show this help message

TRAIN:
  idspam train <ham_path> <spam_path> [model_gob_path]

  Arguments:
    ham_path         Path to ham training set (directory or file set)
    spam_path        Path to spam training set (directory or file set)
    model_gob_path   Optional. Output path for model .gob file.
                     Default: ./model.gob (current working directory)

  Examples:
    idspam train ./data/ham ./data/spam
    idspam train ./data/ham ./data/spam ./nb_model.gob

IDENTIFY:
  idspam identify <model_gob_path> <target_path>

  Arguments:
    model_gob_path   Path to saved model .gob file
    target_path      Path to file to be identified

  Examples:
    idspam identify ./model.gob ./test/email.txt

NOTES:
  - If you omit model_gob_path in 'train', the model is saved to the current directory by default.
  - Use 'idspam help' to see this message.
`)
	return nil
}
func HandleError(err error)error{
	if err!=nil{
		fmt.Fprintln(os.Stderr,"error:",err)
		os.Exit(1)
	}
	return nil
}
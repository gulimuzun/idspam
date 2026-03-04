package train

import (
	"encoding/gob"
	"idspam/internal/process"
	"idspam/internal/token"
	"math"
	"os"
	"fmt"
)

type NaiveBayesModel struct {
	Vocab          map[string]int //所有词的集合
	WordCountSpam  map[string]int //spam类型的词
	WordCountHam   map[string]int //ham类型的词
	TotalSpamWords int            // spam 中所有词出现次数总和
	TotalHamWords  int            // ham 中所有词出现次数总和
	SpamDocs       int            // spam 邮件数
	HamDocs        int            // ham 邮件数

}

func TrainAndCalc(pathHam string,pathSpam string) *NaiveBayesModel {
	var model NaiveBayesModel
	model.SpamDocs = process.CalcTextSum(pathSpam)
	model.HamDocs = process.CalcTextSum(pathHam)
	model.TotalHamWords = process.CalcWordSum(pathHam)
	model.TotalSpamWords = process.CalcWordSum(pathSpam)
	model.WordCountHam = process.ProcessDataToWordBag(pathHam)
	model.WordCountSpam = process.ProcessDataToWordBag(pathSpam)
	model.Vocab = make(map[string]int)
	for k,v:=range process.ProcessDataToWordBag(pathHam){
		model.Vocab[k]=v
	}
	for k,v:=range process.ProcessDataToWordBag(pathSpam){
		model.Vocab[k]+=v
	}
	return &model
}
func OutputTrainFile(model *NaiveBayesModel,path string) error{
	f,err:=os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	enc:= gob.NewEncoder(f)
	return enc.Encode(model)
}
func LoadTrainFile(path string) (*NaiveBayesModel,error){
	f,err:=os.Open(path)
	if err!=nil {
		return nil,err
	}
	defer f.Close()
	enc:=gob.NewDecoder(f)
	var model NaiveBayesModel
	err= enc.Decode(&model)
	if err != nil {
		return nil,err
	}
	return &model,nil
}
func PredictIsHam(model *NaiveBayesModel, path string) bool {
	vocabSize := len(model.Vocab)
	totalDocSize := model.SpamDocs + model.HamDocs
	pHam := math.Log(float64(model.HamDocs) / float64(totalDocSize))
	pSpam := math.Log(float64(model.SpamDocs) / float64(totalDocSize))
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr,"error:",err)
		os.Exit(1)
	}
	wordList := token.Tokenizer(f)
	for _, word := range wordList {
		hamSum, _ := model.WordCountHam[word]
		spamSum, _ := model.WordCountSpam[word]
		pHam += math.Log(float64(hamSum+1) / float64(model.TotalHamWords+vocabSize))
		pSpam += math.Log(float64(spamSum+1) / float64(model.TotalSpamWords+vocabSize))
	}
	f.Close()
	if pHam > pSpam {
		return true
	} else {
		return false
	}
}
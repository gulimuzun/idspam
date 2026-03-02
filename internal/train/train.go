package train

import (
	"fmt"
	"idspam/internal/process"
	"idspam/internal/token"
	"io/fs"
	"math"
	"os"
	"path/filepath"
)
type NaiveBayesModel struct{
	Vocab map[string]int //所有词的集合
	WordCountSpam map[string]int //spam类型的词
	WordCountHam map[string]int //ham类型的词
	TotalSpamWords int                  // spam 中所有词出现次数总和
    TotalHamWords  int                  // ham 中所有词出现次数总和
    SpamDocs       int                  // spam 邮件数
    HamDocs        int                  // ham 邮件数

}
func TrainAndCalc() *NaiveBayesModel{
	var model NaiveBayesModel
	model.SpamDocs=process.CalcTextSum("trainset/spam")
	model.HamDocs=process.CalcTextSum("trainset/ham")
	model.TotalHamWords=process.CalcWordSum("trainset/ham")
	model.TotalSpamWords=process.CalcWordSum("trainset/spam")
	model.WordCountHam=process.ProcessDataToWordBag("trainset/ham")
	model.WordCountSpam=process.ProcessDataToWordBag("trainset/spam")
	model.Vocab=process.ProcessDataToWordBag("trainset")
	return &model
}
func PredictIsHam(model *NaiveBayesModel,path string) bool{
	vocabSize := len(model.Vocab)
	totalDocSize := model.SpamDocs+model.HamDocs
	pHam := math.Log(float64(model.HamDocs)/float64(totalDocSize))
	pSpam := math.Log(float64(model.SpamDocs)/float64(totalDocSize))
	f,err := os.Open(path)
	if err!= nil {
		panic(err)
	}
	wordList := token.Tokenizer(f)
	for _,word := range wordList{
		hamSum,_:=model.WordCountHam[word]
		spamSum,_:=model.WordCountSpam[word]
		pHam+=math.Log(float64(hamSum+1)/float64(model.TotalHamWords+vocabSize))
		pSpam+=math.Log(float64(spamSum+1)/float64(model.TotalSpamWords+vocabSize))
	}
	f.Close()
	if pHam>pSpam{
		return true
	}else{
		return false
	}
}
func Test() error{
	m:=TrainAndCalc()
	hamIsHam,hamIsSpam,spamIsHam,spamIsSpam:=0,0,0,0
	fmt.Println("Ham测试集的情况")
	err := filepath.WalkDir("testset/ham", func(path string, d fs.DirEntry,err error) error{
		if err!=nil{
			return err
		}
		if d.IsDir()==true{
			return nil
		}
		if filepath.Ext(d.Name()) ==".txt"{
			ok:=PredictIsHam(m,path)
			if ok {
				hamIsHam++
			}else {
				hamIsSpam++
			}
		}
		return nil
	})
	if err!=nil {
		panic(err)
	}
	fmt.Printf("ham:%v\nspam:%v\n",hamIsHam,hamIsSpam)
	fmt.Println("spam测试集的情况")
	err = filepath.WalkDir("testset/spam", func(path string, d fs.DirEntry,err error) error{
		if err!=nil{
			return err
		}
		if d.IsDir()==true{
			return nil
		}
		if filepath.Ext(d.Name()) ==".txt"{
			ok:=PredictIsHam(m,path)
			if ok {
				spamIsHam++
			}else {
				spamIsSpam++
			}
		}
		return nil
	})
	if err!=nil {
		panic(err)
	}
	fmt.Printf("ham:%v\nspam:%v",spamIsHam,spamIsSpam)
	return nil
}
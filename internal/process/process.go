package process

import (
	"idspam/internal/token"
	"io/fs"
	"os"
	"path/filepath"
)
func ProcessDataToWordBag(root string) map[string]int{
	wordBag:=make(map[string]int)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry,err error) error{
		if err!=nil{
			return err
		}
		if d.IsDir()==true{
			return nil
		}
		if filepath.Ext(d.Name()) ==".txt"{
			f,err := os.Open(path)
			if(err != nil){
				return err
			}
			bag:=token.Tokenizer(f)
			for _,word := range bag{
				wordBag[word]++
			}
			f.Close()
		}
		return nil
	})
	if err!=nil {
		panic(err)
	}
	return wordBag
}
func CalcWordSum(root string) int {
	sum := 0
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry,err error) error{
		if err!=nil{
			return err
		}
		if d.IsDir()==true{
			return nil
		}
		if filepath.Ext(d.Name()) ==".txt"{
			f,err := os.Open(path)
			if(err != nil){
				return err
			}
			bag:=token.Tokenizer(f)
			sum+=len(bag)
			f.Close()
		}
		return nil
	})
	if err!=nil {
		panic(err)
	}
	return sum
}
func CalcTextSum(root string) int{
	sum := 0
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry,err error) error{
		if err!=nil{
			return err
		}
		if d.IsDir()==true{
			return nil
		}
		if filepath.Ext(d.Name()) ==".txt"{
			sum++
		}
		return nil
	})
	if err!=nil {
		panic(err)
	}
	return sum
}
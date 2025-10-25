package utils

import (
	"crypto/rand"
	"cypto/rand"
	"math/big"
)

const charset="abcdefghijhklmnopqrstuvwxyz0123456789"

func GenerateShortCode(length int) (string, error){
	result:= make([]byte,length)
	charsetLength:=big.NewInt(int64(len(charset)))

	for i:=0;i<length;i++{
		randdomIndex,err:= rand.Int(rand.Reader,charsetLength)

		if err!=nil{
			return  "",err
		}
		result[i]=charset[randdomIndex.Int64()]
	}
	return string(result),nil
}
func GenerateUniqueId() (string,error){
	return  GenerateShortCode(16)
}
package models

import (
	"errors"
	"fmt"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func CheckPasswordLevel(pass string) error {
	pass=strings.ToLower(pass)	
	if  len(pass)<8 {
		return fmt.Errorf("password is less than 8")
	}
	num:=`[0-9]{1}`
	aToZ:=`[a-z]{1}`

	if b,_:=regexp.MatchString(num,pass);!b {
		return fmt.Errorf("password need numbers")	
	}
	if b,_:=regexp.MatchString(aToZ,pass);!b{
		return fmt.Errorf("password need character")
	}
return nil
}

func PasswordHash (password string)(string,error)  {
	if len(password)==0{
		return"",errors.New("password can not be empty")
	}
	h,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return string(h),err
}
func Checkpasswordsame (original,password string)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(original),[]byte(password))
	return err==nil
}
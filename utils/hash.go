package hash

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

//GeneratePassword 给密码就行加密操作
func GeneratePassword(userPassword string) (string, error) {
	byteStr, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	result := string(byteStr)
	return result, err
}

//ValidatePassword 密码比对
func ValidatePassword(userPassword string, hashed string) (err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return err
	}
	return nil

}

func main() {
	// 生成hash密码样例
	userPassword := "123456"
	passwordStr, err := GeneratePassword(userPassword)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(passwordStr)
	// 验证hash密码样例
	mysqlStr := "$2a$10$F0jm7AUM6wOtywZe3VYRG.Jwl0JuHWUvlMyPuqbOI9RCqd/w5Ew6m"
	err = ValidatePassword(userPassword, mysqlStr)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("验证通过")
	}

}

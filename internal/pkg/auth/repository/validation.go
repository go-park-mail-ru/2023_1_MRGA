package repository

import "fmt"

func (r *AuthRepository) CheckUsername(username string) error {
	return nil
}

func (r *AuthRepository) CheckEmail(email string) error {
	return nil
}

func CheckAge(age int) error {
	if age > 150 || age < 18 {

		return fmt.Errorf("age is not correct")
	}
	return nil
}

package repository

import (
	"fmt"
	"time"
)

func (r *AuthRepository) CheckBirthDay(birthDay string) error {
	birth, err := time.Parse("2006-01-02", birthDay)
	if err != nil {
		return err
	}
	now := time.Now()
	age := now.Year() - birth.Year()
	if now.Month() > birth.Month() {
		age -= 1
	}
	if now.Month() == birth.Month() && now.Day() < birth.Day() {
		age -= 1
	}

	if age > 150 || age < 18 {

		return fmt.Errorf("age is not correct")
	}

	return nil
}

func (r *AuthRepository) CheckEmail(email string) error {
	return nil
}
package repository

import (
	"fmt"
	"time"
)

func (r *AuthRepository) CheckBirthDay(birthDay string) error {
	age, err := CalculateAge(birthDay)
	if err != nil {
		return err
	}
	if age > 150 || age < 18 {

		return fmt.Errorf("age is not correct")
	}

	return nil
}

func (r *AuthRepository) CheckEmail(email string) error {
	return nil
}

func CalculateAge(birthDay string) (int, error) {
	birth, err := time.Parse("2006-01-02", birthDay[:10])
	if err != nil {
		return 0, err
	}
	now := time.Now()
	age := now.Year() - birth.Year()
	if now.Month() < birth.Month() {
		age -= 1
	}
	if now.Month() == birth.Month() && now.Day() <= birth.Day() {
		age -= 1
	}
	return age, nil
}

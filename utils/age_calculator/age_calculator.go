package ageCalc

import (
	"fmt"
	"time"
)

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
	if now.Month() == birth.Month() && now.Day() < birth.Day() {
		age -= 1
	}
	return age, nil
}

func CalculateBirthYear(age int) string {
	dataResult := time.Now()
	return fmt.Sprintf("%d-01-01", dataResult.Year()-age)
}

package repository

import "fmt"

func (r *AuthRepository) CheckUsername(username string) error {
	for _, user := range r.Users {
		if username == user.Username {

			return fmt.Errorf("username is not unique")
		}
	}

	return nil
}

func (r *AuthRepository) CheckEmail(email string) error {
	for _, user := range r.Users {
		if email == user.Email {

			return fmt.Errorf("email is not unique")
		}
	}

	return nil
}

func CheckAge(age int) error {
	if age > 150 || age < 18 {

		return fmt.Errorf("age is not correct")
	}
	return nil
}

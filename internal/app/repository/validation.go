package repository

import "fmt"

func (r *Repository) CheckUsername(username string) error {
	for _, user := range *r.Users {
		if username == user.Username {

			return fmt.Errorf("email is not unique")
		}
	}

	return nil
}

func (r *Repository) CheckEmail(email string) error {
	for _, user := range *r.Users {
		if email == user.Email {

			return fmt.Errorf("email is not unique")
		}
	}

	return nil
}

func CheckAge(age int) error {
	if age > 100 || age < 18 {

		return fmt.Errorf("age is not correct")
	}
	return nil
}

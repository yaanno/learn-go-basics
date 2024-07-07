package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func Default() (*User, error) {
	return &User{}, nil
}

func New(firstName, lastName, birthdate string) (*User, error) {
	err := validateUserInput(firstName, lastName, birthdate)

	if err != nil {
		return nil, err
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func validateUserInput(firstName string, lastName string, birthdate string) error {
	var e, e2, e3 error
	if firstName == "" {
		e = errors.New("first name cannot be empty")
	}

	if lastName == "" {
		e2 = errors.New("last name cannot be empty")
	}

	if birthdate == "" {
		e3 = errors.New("birthdate name cannot be empty")
	}

	return errors.Join(e, e2, e3)
}

func (u User) formatFullName() string {
	return fmt.Sprintf("%s %s", u.firstName, u.lastName)
}

func (u User) formatCreatedAt() string {
	return fmt.Sprintf("%v", u.createdAt.Format(time.DateTime))
}

func (u User) GetUserDetails() string {
	return fmt.Sprintf("%s %s %s", u.formatFullName(), u.birthdate, u.formatCreatedAt())
}

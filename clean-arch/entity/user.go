package entity

import "errors"

type User struct {
	ID   int
	Name string
	Age  int
}

func NewUser(name string, age int) (*User, error) {
	user := &User{
		ID:   0,
		Name: name,
		Age:  age,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if u.Name == "" || u.Age == 0 {
		return errors.New("Invalid params")
	}
	return nil
}

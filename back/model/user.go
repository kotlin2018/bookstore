package model

import (
	"back/repository/ent"
	"back/repository/ent/user"
	"context"
	"github.com/spf13/cast"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func CreateUser(ctx context.Context, client *ent.Client, name string, password string, itype string) (User, error) {
	u, err := client.User.Create().SetUsername(name).SetPassword(password).SetType(user.Type(itype)).Save(ctx)
	if err != nil {
		return User{}, err
	}
	return User{
		ID:   u.ID,
		Name: u.Username,
		Type: cast.ToString(u.Type)}, nil
}

func QueryUser(ctx context.Context, client *ent.Client, name string, password string) (User, error) {
	u, err := client.User.Query().Where(user.Username(name), user.Password(password)).Only(ctx)
	if err != nil {
		return User{}, err
	}
	return User{ID: u.ID, Name: u.Username, Type: cast.ToString(u.Type)}, nil
}

func UpdateUser(ctx context.Context, client *ent.Client, ID int, name string, password string) (User, error) {
	u, err := client.User.UpdateOneID(ID).SetUsername(name).SetPassword(password).Save(ctx)
	if err != nil {
		return User{}, err
	}
	return User{ID: u.ID, Name: u.Username, Type: cast.ToString(u.Type)}, nil
}

func DeleteUser(ctx context.Context, client *ent.Client, ID int) error {
	_, err := client.User.Delete().Where(user.ID(ID)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
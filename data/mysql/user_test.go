// +build integration

package mysql

import (
	"boyter/portfold/data"
	"testing"
)

func TestUserGet(t *testing.T) {
	userModel := UserModel{DB: connect(t)}
	_, err := userModel.Get(2147483647)

	if err.Error() != "data: no matching record found" {
		t.Error("Expected to get no matching record", err.Error())
	}
}

func TestUserInsertNoAccountForeignKeyCheck(t *testing.T) {
	userModel := UserModel{DB: connect(t)}

	user := data.User{
		AccountId: 99999999,
		Name:      "test user",
		Email:     "test@example.com",
	}

	_, err := userModel.Insert(user)

	if err == nil {
		t.Error("Expecting error")
	}
}

func TestUserInsertDelete(t *testing.T) {
	accountModel := AccountModel{DB: connect(t)}
	account, err := accountModel.Insert(data.Account{
		Name:    "test account",
		Email:   "test@example.com",
		Details: "some details about this account",
	})

	if err != nil {
		t.Error(err.Error())
	}

	userModel := UserModel{DB: connect(t)}

	zeuser := data.User{
		AccountId: account.Id,
		Name:      "test user",
		Email:     "test@example.com",
	}
	zeuser.HashPassword("password")

	user, err := userModel.Insert(zeuser)

	if err != nil {
		t.Error(err.Error())
	}

	userModel.Delete(*user)
	accountModel.Delete(*account)
}

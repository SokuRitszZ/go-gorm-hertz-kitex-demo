package pack

import (
	"ghkd/kitex_gen/user"
	"ghkd/module/user/model/db"
)

// User pack user info
func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}
	return &user.User{
		ID: int64(u.ID),
		Name: u.Name,
		Avatar: "test",
	}
}

// Users pack list of user info
func Users(us []*db.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if temp := User(u); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}

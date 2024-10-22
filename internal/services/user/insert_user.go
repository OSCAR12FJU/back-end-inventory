package user

import "back-end-inventory/internal/domains"

func (s Services) InsertUser(user domains.Users) (domains.Users, error) {
	return s.Repo.InsertUser(user)
}

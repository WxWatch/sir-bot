package storage

import (
	"fmt"
	"strings"
)

type InMemoryStorage struct {
	users map[string]*User
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		users: make(map[string]*User),
	}
}

func (s *InMemoryStorage) GetUser(guildID string, userID string) (*User, error) {
	key := fmt.Sprintf("%v-%v", guildID, userID)
	user, ok := s.users[key]
	if !ok {
		return nil, nil
	}

	return user, nil
}

func (s *InMemoryStorage) GetUsers(guildID string) ([]*User, error) {
	users := make([]*User, 0)
	for _, user := range s.users {
		if strings.HasPrefix(user.PrimaryKey(), guildID) {
			users = append(users, user)
		}
	}

	return users, nil
}

func (s *InMemoryStorage) SaveUser(user *User) error {
	s.users[user.PrimaryKey()] = user

	return nil
}

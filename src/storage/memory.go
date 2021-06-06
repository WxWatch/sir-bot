package storage

import "fmt"

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

func (s *InMemoryStorage) SaveUser(user *User) error {
	s.users[user.PrimaryKey()] = user

	return nil
}

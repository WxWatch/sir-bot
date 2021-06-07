package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
)

type FlatFileStorage struct {
	db *leveldb.DB
}

func NewFlatFileStorage() *FlatFileStorage {
	db, err := leveldb.OpenFile("data/db", nil)
	if err != nil {
		log.Fatal(err)
	}
	return &FlatFileStorage{
		db: db,
	}
}

func (s *FlatFileStorage) GetUser(guildID string, userID string) (*User, error) {
	key := fmt.Sprintf("%v-%v", guildID, userID)
	data, err := s.db.Get([]byte(key), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	var user User
	json.Unmarshal(data, &user)

	return &user, nil
}

func (s *FlatFileStorage) GetUsers(guildID string) ([]*User, error) {
	users := make([]*User, 0)
	iter := s.db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		if strings.HasPrefix(string(key), guildID) {
			var user User
			json.Unmarshal(value, &user)
			users = append(users, &user)
		}
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *FlatFileStorage) SaveUser(user *User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = s.db.Put([]byte(user.PrimaryKey()), userBytes, nil)
	if err != nil {
		return err
	}

	return nil
}

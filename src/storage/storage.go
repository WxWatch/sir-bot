package storage

type Storage interface {
	GetUser(guildID, userID string) (*User, error)
	SaveUser(user *User) error
}

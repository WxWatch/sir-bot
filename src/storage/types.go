package storage

import (
	"fmt"
)

const (
	startingExperience = 150
	startingLevel      = 0
)

type options struct {
	userID  string
	guildID string
}

type Option interface {
	apply(*options)
}

type userIDOption struct {
	UserID string
}

func (s userIDOption) apply(opts *options) {
	opts.userID = s.UserID
}

func WithUserID(userID string) Option {
	return userIDOption{UserID: userID}
}

type guildIDOption struct {
	GuildID string
}

func (s guildIDOption) apply(opts *options) {
	opts.guildID = s.GuildID
}

func WithGuildID(guildID string) Option {
	return guildIDOption{GuildID: guildID}
}

// ByLevel implements sort.Interface based on the Level field.
type ByLevel []*User

func (a ByLevel) Len() int           { return len(a) }
func (a ByLevel) Less(i, j int) bool { return a[i].Level < a[j].Level }
func (a ByLevel) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type User struct {
	ID         string `json:"user_id"`
	GuildID    string `json:"guild_id"`
	Experience uint   `json:"experience"`
	Level      uint   `json:"level"`
}

func NewUser(opts ...Option) *User {
	options := options{
		userID:  "",
		guildID: "",
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &User{
		ID:         options.userID,
		GuildID:    options.guildID,
		Experience: startingExperience,
		Level:      startingLevel,
	}
}

func (u *User) PrimaryKey() string {
	return fmt.Sprintf("%v-%v", u.GuildID, u.ID)
}

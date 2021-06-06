package leveling

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"wxwatch.dev/bot/src/storage"
)

type options struct {
	storage storage.Storage
}

type Option interface {
	apply(*options)
}

type storageOption struct {
	Storage storage.Storage
}

func (s storageOption) apply(opts *options) {
	opts.storage = s.Storage
}

func WithStorage(storage storage.Storage) Option {
	return storageOption{Storage: storage}
}

type LevelingListener struct {
	storage storage.Storage
}

func NewLevelingListener(opts ...Option) *LevelingListener {
	options := options{
		storage: storage.NewInMemoryStorage(),
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &LevelingListener{
		storage: options.storage,
	}
}

func (l *LevelingListener) Listen(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	// Fetch user
	user, err := l.storage.GetUser(m.GuildID, m.Author.ID)
	if err != nil {
		return
	}

	// Create a user if there isn't one
	if user == nil {
		user = storage.NewUser(
			storage.WithUserID(m.Author.ID),
			storage.WithGuildID(m.GuildID),
		)
	}

	// Add exp
	exp := rand.Intn(11)
	user.Experience += uint(exp)

	// Check if new level, send message
	threshold := ExpForLevel(user.Level + 1)
	leveledUp := user.Experience >= threshold
	if leveledUp {
		user.Level += 1
	}

	// Save
	err = l.storage.SaveUser(user)
	if err != nil {
		return
	}

	// Respond
	if leveledUp {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Congrats <@%v>, you are now level %v", m.Author.ID, user.Level))
	}

}

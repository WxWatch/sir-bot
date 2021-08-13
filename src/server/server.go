package server

import (
	"wxwatch.dev/bot/src/cache"
	"wxwatch.dev/bot/src/commands"
	"wxwatch.dev/bot/src/discord"
	"wxwatch.dev/bot/src/leveling"
	"wxwatch.dev/bot/src/storage"
)

type Server struct {
	bot   *discord.Bot
	cache *cache.Cache
}

type Options struct {
	DiscordOptions *discord.Options
}

func NewServer(options *Options) *Server {
	botOptions := &discord.Options{
		Token: options.DiscordOptions.Token,
	}
	bot := discord.NewBot(botOptions)

	return &Server{
		bot:   bot,
		cache: cache.NewCache(),
	}
}

func (s *Server) Start() error {
	storage := storage.NewFlatFileStorage()
	routerOptions := &commands.Options{
		Cache:   s.cache,
		Storage: storage,
	}
	router := commands.NewCommandRouter(routerOptions)
	// listener := listener.NewMessageCreateListener()
	levelingListener := leveling.NewLevelingListener(
		leveling.WithStorage(storage),
	)

	handlers := []interface{}{
		router.Route,
		levelingListener.Listen,
		// listener.Listen,
	}

	s.bot.SetupHandlers(handlers)

	err := s.bot.Connect()
	if err != nil {
		return err
	}

	// Need the bot user to set up all this
	router.SetupSlashCommands(s.bot)
	levelingListener.Setup(s.bot)

	return err
}

func (s *Server) Stop() {
	s.bot.Disconnect()
}

package commander

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

type runFunc func(cmd *Command, args []string, s *discordgo.Session, e *discordgo.MessageCreate)

type options struct {
	usage         string
	description   string
	argsValidator ArgValidator
	run           runFunc
	ctx           context.Context
}

type Option interface {
	apply(*options)
}

// Usage
type usageOption struct {
	Usage string
}

func (o usageOption) apply(opts *options) {
	opts.usage = o.Usage
}

func WithUsage(usage string) Option {
	return usageOption{Usage: usage}
}

// Description
type descriptionOption struct {
	Description string
}

func (o descriptionOption) apply(opts *options) {
	opts.description = o.Description
}

func WithDescription(description string) Option {
	return descriptionOption{Description: description}
}

// ArgsValidator
type argsValidatorOption struct {
	ArgsValidator ArgValidator
}

func (o argsValidatorOption) apply(opts *options) {
	opts.argsValidator = o.ArgsValidator
}

func WithArgsValidator(argsValidator ArgValidator) Option {
	return argsValidatorOption{ArgsValidator: argsValidator}
}

// Description
type runOption struct {
	Run runFunc
}

func (o runOption) apply(opts *options) {
	opts.run = o.Run
}

func WithRunFunc(run runFunc) Option {
	return runOption{Run: run}
}

// Context
type ctxOption struct {
	Ctx context.Context
}

func (o ctxOption) apply(opts *options) {
	opts.ctx = o.Ctx
}

func WithContext(ctx context.Context) Option {
	return ctxOption{Ctx: ctx}
}

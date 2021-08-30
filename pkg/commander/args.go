package commander

import "fmt"

type ArgValidator func(cmd *Command, args []string) error

// ExactArgs returns an error if there are not exactly n args.
func ExactArgs(n int) ArgValidator {
	return func(cmd *Command, args []string) error {
		if len(args) != n {
			return fmt.Errorf("accepts %d arg(s), received %d", n, len(args))
		}
		return nil
	}
}

// ArbitraryArgs never returns an error.
func ArbitraryArgs(cmd *Command, args []string) error {
	return nil
}

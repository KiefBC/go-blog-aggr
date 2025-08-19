package commands

import (
	"context"
	"fmt"

	"github.com/KiefBC/blog-aggr/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	// TODO: Uncomment when ready to use command arguments

	// if len(cmd.Args) < 1 {
	// 	return fmt.Errorf("Usage: %s", s.Commands.GetUsage(cmd.Name))
	// }
	//
	// feedURL := cmd.Args[0]

	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to fetch feed %s: %v", feedURL, err)
	}

	fmt.Printf("%+v\n", feed)

	return nil
}

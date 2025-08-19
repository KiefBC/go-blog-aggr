package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/KiefBC/blog-aggr/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Usage: %s", s.Commands.GetUsage(cmd.Name))
	}

	time_between_req, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration format: %v", err)
	}
	fmt.Printf("Collecting feeds every %s\n", time_between_req)

	ticker := time.NewTicker(time_between_req)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			fmt.Printf("Error scraping feeds: %v\n", err)
		}
	}
}

func scrapeFeeds(s *State) error {
	nextFeed, err := s.Db.GetNextFeedToFetch(
		context.Background(),
	)
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %w", err)
	}

	_, err = s.Db.MarkFeedFetched(
		context.Background(),
		nextFeed.ID,
	)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	feed, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed %s: %w", nextFeed.Url, err)
	}

	for _, item := range feed.Channel.Item {
		fmt.Printf("%+v\n", item.Title)
	}

	return nil
}

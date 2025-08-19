package commands

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/KiefBC/blog-aggr/internal/database"
	"github.com/KiefBC/blog-aggr/internal/rss"
	"github.com/google/uuid"
)

// HandlerAgg is a command handler that aggregates feeds at a specified interval.
// It takes a duration as an argument, which specifies how often to scrape the feeds.
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

// ScrapeFeeds fetches the next feed from the database, marks it as fetched,
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
		publishedAt := parseRSSDate(item.PubDate)

		_, err := s.Db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       item.Title,
				Url:         item.Link,
				Description: item.Description,
				PublishedAt: publishedAt,
				FeedID:      nextFeed.ID,
			},
		)
		if err != nil {
			if strings.Contains(err.Error(), "unique constraint") || strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			fmt.Printf("Error creating post %s: %v\n", item.Title, err)
		} else {
			fmt.Printf("Created post: %s\n", item.Title)
		}
	}

	return nil
}

// Parse RSS date strings into time.Time
func parseRSSDate(dateStr string) time.Time {
	formats := []string{
		time.RFC1123Z,               // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC1123,                // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC822Z,                // "02 Jan 06 15:04 -0700"
		time.RFC822,                 // "02 Jan 06 15:04 MST"
		"2006-01-02T15:04:05Z07:00", // ISO 8601
		"2006-01-02T15:04:05Z",      // ISO 8601 UTC
		"2006-01-02 15:04:05",       // Simple format
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t
		}
	}

	fmt.Printf("Warning: Could not parse date '%s', using current time\n", dateStr)
	return time.Now()
}

package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/KiefBC/blog-aggr/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit := int32(2) // Default limit

	if len(cmd.Args) > 0 {
		parsedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %v", err)
		}
		if parsedLimit <= 0 {
			return fmt.Errorf("limit must be a positive number")
		}

		limit = int32(parsedLimit)
	}

	posts, err := s.Db.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  limit,
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found. Make sure you're following some feeds!")
		return nil
	}

	fmt.Printf("Found %d posts:\n", len(posts))
	for _, post := range posts {
		fmt.Printf("* %s from %s\n", post.Title, post.FeedName)
		fmt.Printf("  %s\n", post.Url)
		if post.Description != "" {
			fmt.Printf("  %s\n", post.Description)
		}
		fmt.Println()
	}

	return nil
}

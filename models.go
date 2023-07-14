package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/rodmedeiross/scratch/internal/database"
)

type User struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	ApiKey string `json:"api_key"`
}

type Feed struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	Url string `json:"url"`
	UserID uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID uuid.UUID `json:"user_id"`
	FeedID uuid.UUID `json:"feed_id"`
}

type Posts struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title string `json:"title"`
	Description *string `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url string `json:"url"`
	FeedID uuid.UUID `json:"feed_id"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		ApiKey: dbUser.ApiKey,
	}
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, len(dbFeeds))
	for i, dbFeed := range dbFeeds {
		feeds[i] = databaseFeedToFeed(dbFeed)
	}
	return feeds
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID: dbFeedFollows.ID,
		CreatedAt: dbFeedFollows.CreatedAt,
		UpdatedAt: dbFeedFollows.UpdatedAt,
		UserID: dbFeedFollows.UserID,
		FeedID: dbFeedFollows.FeedID,
	}
}

func databaseFeedsFollowsToFeedsFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedsFollows := make([]FeedFollow, len(dbFeedFollows))
	for i, dbFeedFollow := range dbFeedFollows {
		feedsFollows[i] = databaseFeedFollowsToFeedFollows(dbFeedFollow)
	}
	return feedsFollows
}

func databasePostsToPosts(dbPosts database.Post) Posts {
	var descriptions *string
	if dbPosts.Description.Valid{
		descriptions = &dbPosts.Description.String
	}

	return Posts{
		ID: dbPosts.ID,
		CreatedAt: dbPosts.CreatedAt,
		UpdatedAt: dbPosts.UpdatedAt,
		Title: dbPosts.Title,
		Description: descriptions,
		PublishedAt: dbPosts.PublishedAt,
		Url: dbPosts.Url,
		FeedID: dbPosts.FeedID,
	}
}

func databasePostsToPostsList(dbPosts []database.Post) []Posts {
	posts := make([]Posts, len(dbPosts))
	for i, dbPost := range dbPosts {
		posts[i] = databasePostsToPosts(dbPost)
	}
	return posts
}

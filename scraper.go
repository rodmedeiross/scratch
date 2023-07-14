package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/rodmedeiross/scratch/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration){
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error getting feeds to fetch", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	log.Printf("scraping feed %s", feed.Url)

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("error marking feed as fetched %s", feed.Url)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("error fetching feed %s", feed.Url)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Printf("item %s", item.Title, "on feed", feed.Name)
	}

	log.Printf("Feed %s collected, %v post found", feed.Name, len(rssFeed.Channel.Item))
}


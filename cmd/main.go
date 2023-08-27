package main

import (
	tgClient "github.com/Amore14rn/article-bot/internal/clients/telegram"
	"github.com/Amore14rn/article-bot/internal/config"
	event_consumer "github.com/Amore14rn/article-bot/internal/consumer/event-consumer"
	"github.com/Amore14rn/article-bot/internal/events/telegram"
	"github.com/Amore14rn/article-bot/internal/storage/mongo"
	"log"
	"time"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func main() {
	cfg := config.MustLoad()

	storage := mongo.New(cfg.MongoConnectionString, 10*time.Second)

	eventsProcessor := telegram.New(
		tgClient.NewClient(tgBotHost, cfg.TgBotToken),
		storage,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("mongo_example")

const URI = "mongodb://127.0.0.1:27017"

func getClient(ctx context.Context) (*mongo.Client, error) {
	ctx, span := FollowSpan(ctx, "getClient")
	defer span.End()

	opts := options.Client()
	opts.ApplyURI(URI)
	optsAuth := options.Credential{
		Username:   "root",
		Password:   "example",
		AuthSource: "admin",
	}

	opts.SetAuth(optsAuth)

	opts.Monitor = otelmongo.NewMonitor()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	{
		tc, err := initProvider()
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			errTCStop := tc.Stop(ctx)
			if errTCStop != nil {
				log.Fatal(errTCStop)
			}
		}()
	}

	// =========
	// ==START==
	// =========
	ctx, span := tracer.Start(ctx, "main")
	defer span.End()

	// Создаём клиента
	client, err := getClient(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		if mongoDisconectErr := client.Disconnect(ctx); mongoDisconectErr != nil {
			log.Println(mongoDisconectErr)
		}
	}()

	// Все примеры будут для одной колллекции поэтому сразу создаём соответствующий объект
	col := client.Database("teta").Collection("books")

	// Чтобы точно начать с чистого листа удалим коллекцию вообще.
	err = col.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}

	insertOne(ctx, col)
	insertMany(ctx, col)
	findAll(ctx, col)
	find(ctx, col)
	findWithCondition(ctx, col)
	findWithOrCondition(ctx, col)
	updateOne(ctx, col)
	updateMany(ctx, col)
	replaceOne(ctx, col)
	deleteAllRows(ctx, col)
}

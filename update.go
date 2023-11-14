package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func updateOne(ctx context.Context, col *mongo.Collection) {
	ctx, span := FollowSpan(ctx, "updateOne")
	defer span.End()

	filter := bson.M{
		"title": "Высоконагруженные приложения. Программирование, масштабирование, поддержка",
	}

	update := bson.M{
		"$set": bson.M{"rating": 5},
	}

	result, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updating 1 documents...")
	fmt.Printf("Matched docs: %d. Updated docs: %d\n", result.MatchedCount, result.ModifiedCount)
	fmt.Println("=============================")
}

func updateMany(ctx context.Context, col *mongo.Collection) {
	ctx, span := FollowSpan(ctx, "updateMany")
	defer span.End()

	filter := bson.M{
		"rating": 0,
	}

	update := bson.M{
		"$set": bson.M{"rating": 3},
	}

	result, err := col.UpdateMany(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updating many documents...")
	fmt.Printf("Matched docs: %d. Updated docs: %d\n", result.MatchedCount, result.ModifiedCount)
	fmt.Println("=============================")
}

func replaceOne(ctx context.Context, col *mongo.Collection) {
	ctx, span := FollowSpan(ctx, "replaceOne")
	defer span.End()

	fmt.Println("replaceOne document...")
	var updatedBook Book = Book{
		Author: "Ньюмен Сэм",
		Title:  "Создание микросервисов",
		Year:   2016,
	}

	result, err := col.ReplaceOne(ctx, bson.M{"author": "Ньюмен Сэм"}, updatedBook, options.Replace().SetUpsert(true))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched docs: %d. Updated docs: %d. Upserted docs: %d\n", result.MatchedCount, result.ModifiedCount, result.UpsertedCount)
	fmt.Println("=============================")
}

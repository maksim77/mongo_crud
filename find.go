package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func checkFindErr(err error) {
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(err)
	}
}

func findAll(ctx context.Context, col *mongo.Collection) {
	ctx, span := FollowSpan(ctx, "findAll")
	defer span.End()

	fmt.Println("replaceOne document...")
	cursor, err := col.Find(ctx, bson.M{})
	checkFindErr(err)

	var books []Book
	err = cursor.All(ctx, &books)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Finded %d docs:\n%v\n", len(books), books)
	fmt.Println("=============================")
}

func find(ctx context.Context, col *mongo.Collection) {
	ctx, span := FollowSpan(ctx, "find")
	defer span.End()

	opts := options.Find().SetSort(bson.M{"rating": 1})

	cursor, err := col.Find(ctx, bson.M{"year": 2022}, opts)
	checkFindErr(err)

	var books []Book
	err = cursor.All(ctx, &books)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Search documents where year equal 2022...")
	fmt.Printf("Finded %d docs:\n%v\n", len(books), books)
	fmt.Println("=============================")
}

func findWithCondition(ctx context.Context, col *mongo.Collection) {
	ctx, span := FollowSpan(ctx, "findWithCondition")
	defer span.End()

	filter := bson.M{"year": bson.M{"$gt": 2020}}

	cursor, err := col.Find(ctx, filter)
	checkFindErr(err)

	var books []Book
	err = cursor.All(ctx, &books)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Search documents where year greather then 2020...")
	fmt.Printf("Finded %d docs:\n%v\n", len(books), books)
	fmt.Println("=============================")
}

func findWithOrCondition(ctx context.Context, col *mongo.Collection) {
	ctx, span := FollowSpan(ctx, "findWithOrCondition")
	defer span.End()

	filter := bson.M{
		"$or": bson.A{
			bson.M{
				"year": bson.M{"$gte": 2020},
			},
			bson.M{
				"author": "Касун Индрасири",
			},
		},
	}

	findOptions := options.Find()
	findOptions.SetProjection(bson.M{"author": 0})

	cursor, err := col.Find(ctx, filter, findOptions)
	checkFindErr(err)

	var books []Book
	err = cursor.All(ctx, &books)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Search documents where year greather then 2020 AND author is 'Касун Индрасири'...")
	fmt.Printf("Finded %d docs:\n%v\n", len(books), books)
	fmt.Println("=============================")
}

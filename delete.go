package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func deleteAllRows(ctx context.Context, coll *mongo.Collection) {
	ctx, span := FollowSpan(ctx, "deleteAllRows")
	defer span.End()

	result, err := coll.DeleteMany(ctx, bson.D{})
	if err != nil {
		log.Fatal(result)
	}

	fmt.Println("Remove all documents...")
	fmt.Printf("%d documents removed\n", result.DeletedCount)
	fmt.Println("=============================")
}

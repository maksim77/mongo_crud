package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func insertOne(ctx context.Context, col *mongo.Collection) {
	ctx, span := FollowSpan(ctx, "insertOne")
	defer span.End()

	fmt.Println("Inserting 1 documents...")
	result, err := col.InsertOne(ctx, book)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("One document inserted with id: %s\n", result.InsertedID)
	fmt.Println("=============================")
}

func insertMany(ctx context.Context, col *mongo.Collection) {
	ctx, span := FollowSpan(ctx, "insertMany")
	defer span.End()

	fmt.Println("Inserting 2 documents...")
	inserts := make([]interface{}, 0, len(books))
	for _, book := range books {
		inserts = append(inserts, book)
	}

	result, err := col.InsertMany(ctx, inserts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d documents inserted with ids: %v\n", len(result.InsertedIDs), result.InsertedIDs)

	// ObjectID описание
	for _, v := range result.InsertedIDs {
		id := [12]byte(v.(primitive.ObjectID))
		byteTime := id[0:4]
		byteRandomID := id[4:9]
		byteInc := id[9:12]
		fmt.Printf("Timestamp: %d\n", binary.BigEndian.Uint32(byteTime))
		fmt.Printf("Timestamp to date: %v\n", time.Unix(int64(binary.BigEndian.Uint32(byteTime)), 0))
		fmt.Printf("Random val per process and machine: %d\n", binary.BigEndian.Uint32(byteRandomID))
		fmt.Printf("Inc counter: %d\n", binary.BigEndian.Uint16(byteInc))
		fmt.Println("*******")
	}

	fmt.Println("=============================")
}

// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"ent-todo/internal/ent"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "test:test@tcp(10.50.23.194:3306)/ent?parseTime=true")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.TODO()
	err = Do(ctx, client)
	if err != nil {
		log.Fatalf("create info failed %v", err)
	}

	log.Println("create done")
}

func Do(ctx context.Context, client *ent.Client) error {


	name := fmt.Sprintf("Mashraki - %d", time.Now().Nanosecond())

	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName(name).
		SetUsername(name).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	log.Println("user:", a8m)
	card1, err := client.Card.
		Create().
		SetOwner(a8m).
		SetNumber("1020").
		SetExpired(time.Now().Add(time.Minute)).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating card: %w", err)
	}
	log.Println("card:", card1)
	// Only returns the card of the user,
	// and expects that there's only one.
	card2, err := a8m.QueryCard().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying card: %w", err)
	}
	log.Println("card:", card2)
	// The Card entity is able to query its owner using
	// its back-reference.
	owner, err := card2.QueryOwner().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying owner: %w", err)
	}
	log.Println("owner:", owner)
	return nil
}


func createTables()  {
	client, err := ent.Open("mysql", "test:test@tcp(10.50.23.194:3306)/ent")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
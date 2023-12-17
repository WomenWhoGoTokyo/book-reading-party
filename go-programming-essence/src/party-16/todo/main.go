package main

import (
	"context"
	"fmt"
	"log"
	"todo/ent"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// スキーマ作成 (マイグレーション)
	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Todoを列挙
	for _, e := range client.Todo.Query().AllX(context.Background()) {
		fmt.Println(e.Text)
	}

	// Todoを作成
	_, err = client.Todo.Create().SetText("Go言語を学ぶ").Save(context.Background())
	if err != nil {
		log.Fatalf("failed creating a todo: %v", err)
	}
}

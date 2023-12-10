package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type Todo struct {
	bun.BaseModel `bun:"table:todos,alias:t"`

	ID        int64     `bun:"id,pk,autoincrement"`
	Content   string    `bun:"content,notnull"`
	Done      bool      `bun:"done"`
	Until     time.Time `bun:"until,nullzero"`
	CreatedAt time.Time
	UpdatedAt time.Time `bun:",nullzero"`
	DeletedAt time.Time `bun:",soft_delete,nullzero`
}

func main() {
	sqldb, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// 接続情報の例
	// sqldb, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer sqldb.Close()

	db := bun.NewDB(sqldb, pgdialect.New())
	defer db.Close()

	// TODO 抽出
	var todos []Todo
	ctx := context.Background()
	err = db.NewSelect().Model(&todos).Order("created_at").Where("until is not null").Where("done is false").Scan(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(todos) == 0 {
		return
	}

	from := mail.Address{Name: "TODO Reminder", Address: os.Getenv("MAIL_FROM")}
	var buf bytes.Buffer
	buf.WriteString("From: " + from.String() + "\r\n")
	buf.WriteString("To: " + os.Getenv("MAIL_TO") + "\r\n")
	buf.WriteString("Subject: TODO Reminder\r\n")
	buf.WriteString("\r\n")
	buf.WriteString("This is your todo list\n\n")
	for _, todo := range todos {
		fmt.Fprintf(&buf, "%s %s\n", todo.Until, todo.Content)
	}

	smtpAuth := smtp.PlainAuth(
		os.Getenv("MAIL_DOMAIN"),
		os.Getenv("MAIL_USER"),
		os.Getenv("MAIL_PASSWORD"),
		os.Getenv("MAIL_AUTHSERVER"),
		// 接続情報の例
		// "example.com",
		// "user",
		// "password",
		// "auth.example.com",
	)

	err = smtp.SendMail(
		os.Getenv("MAIL_SERVER"),
		// 接続情報の例
		// "mail.example.com:10025",
		smtpAuth,
		from.Address,
		[]string{os.Getenv("MAIL_TO")},
		// 接続情報の例
		// []string{"todo-to@example.com"},
		buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}

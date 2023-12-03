package main

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

//go:embed templates
var templates embed.FS

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

type Data struct {
	Todos  []Todo
	Errors []error
}

type Template struct {
	templates *template.Template
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

	db.AddQueryHook(bundebug.NewQueryHook(
		//bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	ctx := context.Background()
	_, err = db.NewCreateTable().Model((*Todo)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Renderer = &Template{
		templates: template.Must(template.New("").
			Funcs(template.FuncMap{
				"FormatDateTime": formatDateTime,
			}).ParseFS(templates, "templates/*")),
	}

	e.GET("/", func(c echo.Context) error {
		var todos []Todo
		ctx := context.Background()
		err := db.NewSelect().Model(&todos).Order("created_at").Scan(ctx)
		if err != nil {
			e.Logger.Error(err)
			return c.Render(http.StatusBadRequest, "index", Data{
				Errors: []error{errors.New("cannot get todos")},
			})
		}
		return c.Render(http.StatusOK, "index", Data{Todos: todos})
	})

	e.POST("/", func(c echo.Context) error {
		var todo Todo
		// フォームパラメータをフィールドにバインドする
		errs := echo.FormFieldBinder(c).
			Int64("id", &todo.ID).
			String("content", &todo.Content).
			Bool("done", &todo.Done).
			CustomFunc("until", CustomFunc(&todo)).
			BindErrors()
		if errs != nil {
			e.Logger.Error(errs)
			return c.Render(http.StatusBadRequest, "index", Data{Errors: errs})
		} else if todo.ID == 0 {
			// IDが0の場合は新規作成する
			ctx := context.Background()
			if todo.Content == "" {
				err = errors.New("Todo not found")
			} else {
				_, err = db.NewInsert().Model(&todo).Exec(ctx)
				if err != nil {
					e.Logger.Error(err)
					err = errors.New("Cannot update")
				}
			}
		} else {
			ctx := context.Background()
			if c.FormValue("delete") != "" {
				// 削除
				_, err = db.NewDelete().Model(&todo).Where("id = ?", todo.ID).Exec(ctx)
			} else {
				// 更新
				var orig Todo
				err = db.NewSelect().Model(&orig).Where("id = ?", todo.ID).Scan(ctx)
				if err == nil {
					orig.Done = todo.Done
					_, err = db.NewUpdate().Model(&orig).Where("id = ?", todo.ID).Exec(ctx)
				}
			}
			if err != nil {
				e.Logger.Error(err)
				err = errors.New("Cannot update")
			}
		}
		if err != nil {
			return c.Render(http.StatusBadRequest, "index", Data{Errors: []error{err}})
		}
		return c.Redirect(http.StatusFound, "/")
	})
	e.Logger.Fatal(e.Start(":8989"))
}

func CustomFunc(todo *Todo) func([]string) []error {
	return func(values []string) []error {
		if len(values) == 0 || values[0] == "" {
			return nil
		}
		dt, err := time.Parse("2006-01-02T15:04 MST", values[0]+" JST")
		if err != nil {
			return []error{echo.NewBindingError("until", values[0:1], "failed to decode time", err)}
		}
		todo.Until = dt
		return nil
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func formatDateTime(d time.Time) string {
	if d.IsZero() {
		return ""
	}
	return d.Format("2006-01-02 15:04")
}

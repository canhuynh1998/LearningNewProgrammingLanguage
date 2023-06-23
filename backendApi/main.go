package main
import (
	// "context"
	// "fmt"
	// "log"
	// "time"
	// "github.com/jackc/pgx/v5"
	// "github.com/gofiber/fiber"
	"go-practice/backendApi/app"
)

func main() {
	// dsn := "postgresql://can:bophtyKcIvwkh3IOV0U1Kw@steel-lizard-7333.6wr.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	// ctx := context.Background()
	// conn, err := pgx.Connect(ctx, dsn)
	// defer conn.Close(context.Background())
	// if err != nil {
	// 	log.Fatal("failed to connect database", err)
	// }

	// var now time.Time
	// err = conn.QueryRow(ctx, "SELECT NOW()").Scan(&now)
	// if err != nil {
	// 	log.Fatal("failed to execute query", err)
	// }

	// fmt.Println(now)
  // Fiber instance
  app.AppInit()
	  }
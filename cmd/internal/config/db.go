package config

import (
    "database/sql"
    "log"
    "time"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
    var err error
    const maxRetries = 5

    for i := 0; i < maxRetries; i++ {
        DB, err = sql.Open("postgres", "host=db port=5432 user=user password=password dbname=price_alert_db sslmode=disable")
        if err != nil {
            log.Printf("Failed to open database connection: %v. Retrying in 5 seconds... (%d/%d)", err, i+1, maxRetries)
            time.Sleep(5 * time.Second)
            continue
        }

        err = DB.Ping()
        if err == nil {
            log.Println("Database connection established successfully")
            return
        }

        log.Printf("Failed to ping database: %v. Retrying in 5 seconds... (%d/%d)", err, i+1, maxRetries)
        time.Sleep(5 * time.Second)
    }

    log.Fatalf("Could not connect to the database after %d attempts: %v", maxRetries, err)
}

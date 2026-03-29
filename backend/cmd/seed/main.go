package main

import (
	"fmt"
	"os"

	"chukotka/core/config"
	"chukotka/core/db"
)

// Тот же сценарий, что в main.go: config.Load() + db.InitDatabase()
// (подключение, миграции, сиды). Запуск из каталога backend: go run ./cmd/seed
func main() {
	config.Load()
	db.InitDatabase()
	fmt.Fprintln(os.Stderr, "Done: migrate + seeds.")
}

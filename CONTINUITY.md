# Continuity Ledger

- **Goal:** Документировать, как поднять PostgreSQL и создать таблицы для бэкенда Chukotka.
- **Stack:** Go, GORM, PostgreSQL; `InitDatabase()` вызывает `AutoMigrate`, затем `SeedAdmin`, `SeedAbout`, `SeedMapData()` (районы/сёла из `seed_geo.json`, встраивается в бинарник).
- **Env:** `godotenv` грузит только `.env` из cwd (не `.env.local` автоматически) — UNCONFIRMED: нужен ли дубль/симлинк для локальной разработки.

## Как «запустить сид»

- **Вместе с API:** при старте бэкенда `db.InitDatabase()` → `Connect` + `Migrate` + `RunAllSeeds()`.
- **Отдельно (без HTTP):** из каталога `backend`: `go run ./cmd/seed` — вызывает `db.InitDatabase()` (как `main.go`), без дублирования кода БД.

## Примечания

- Сиды районов: проверка «есть ли запись» через `Count`/`Find`, не через `First`, чтобы GORM не логировал шумный `record not found` (это не сбой, а отсутствие строки).

## Open questions

- Нужно ли явно подключать `.env.local` в коде — сейчас в `config.Load()` только `godotenv.Load()`.

# README.md

## Проект CHUKOTKA

### Структура проекта

```
CHUKOTKA/
├── .vscode/
├── backend/
│   ├── controllers/
│   ├── core/
│   ├── middlewares/
│   ├── models/
│   ├── repositories/
│   ├── routes/
│   ├── utils/
│   ├── .env
│   ├── .env.local
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── nohup.out
├── frontend/
│   ├── node_modules/
│   ├── public/
│   ├── src/
│   ├── .gitignore
│   ├── index.html
│   ├── package-lock.json
│   ├── package.json
│   ├── postcss.config.js
│   ├── README.md
│   ├── tailwind.config.js
│   ├── tsconfig.app.json
│   ├── tsconfig.json
│   ├── tsconfig.node.json
│   ├── tsconfig.tsbldinfo
│   └── vite.config.ts
└── README.md
```

## Настройка проекта

### 1. Настройка PostgreSQL

#### Установка PostgreSQL

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib -y
```

#### Настройка удаленного подключения (если необходимо)

```bash
sudo nano /etc/postgresql/15/main/postgresql.conf
```

Измените параметр `listen_addresses` на `'*'`

#### Настройка аутентификации

```bash
sudo nano /etc/postgresql/15/main/pg_hba.conf
```

Добавьте строку для разрешения удаленных подключений:

```
host all all 0.0.0.0/0 md5
```

#### Перезапуск PostgreSQL

```bash
sudo systemctl restart postgresql
```

#### Инициализация базы данных из дампа

```bash
# Восстановление базы данных из дампа
psql -U postgres -d your_database -f /path/to/dump.sql
```

### 2. Настройка бэкенда (Go)

#### Установка Go

```bash
sudo apt install golang -y
```

#### Настройка переменных окружения для Go

```bash
echo "export GOPATH=$HOME/go" >> ~/.bashrc
echo "export PATH=$PATH:$GOPATH/bin" >> ~/.bashrc
source ~/.bashrc
```

#### Клонирование репозитория

```bash
git clone <URL_ВАШЕГО_РЕПОЗИТОРИЯ>
cd backend
```

#### Установка зависимостей

```bash
go mod download
```

#### Настройка переменных окружения (.env)

Создайте или отредактируйте файл `.env`:

```
DB_PORT="5432"
DB_PASS="your_db_password"
DB_USER="your_db_user"
DB_NAME="your_db_name"
DB_HOST="localhost"

TOKEN_TTL="2000"
JWT_PRIVATE_KEY="THE_KEY"
ADMIN_PASSWORD="your_admin_password"
```

#### Компиляция и запуск

```bash
# Скомпилируйте проект
go build -o main .

# Запустите приложение
./main
```

#### Настройка systemd сервиса (для автозапуска)

Создайте файл `/etc/systemd/system/chukotka-backend.service`:

```ini
[Unit]
Description=Chukotka Backend Service
After=network.target

[Service]
User=your_user
WorkingDirectory=/path/to/backend
ExecStart=/path/to/backend/main
Restart=always
EnvironmentFile=/path/to/backend/.env

[Install]
WantedBy=multi-user.target
```

Активируйте сервис:

```bash
sudo systemctl daemon-reload
sudo systemctl start chukotka-backend.service
sudo systemctl enable chukotka-backend.service
```

### 3. Настройка фронтенда (Vue.js)

#### Установка Node.js и npm

```bash
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
sudo apt install nodejs -y
sudo npm install -g npm
```

#### Клонирование репозитория

```bash
git clone <URL_ВАШЕГО_ФРОНТЕНД_РЕПОЗИТОРИЯ>
cd frontend
```

#### Установка зависимостей

```bash
npm install
```

#### Сборка проекта для продакшена

```bash
npm run build
```

### 4. Настройка веб-сервера (Nginx)

#### Установка Nginx

```bash
sudo apt install nginx -y
```

#### Настройка конфигурации

Создайте файл `/etc/nginx/sites-available/chukotka`:

```nginx
server {
    listen 80;
    server_name your_domain;

    location / {
        root /path/to/frontend/dist;
        index index.html index.htm;
        try_files $uri $uri/ /index.html;
    }

    location /api {
        proxy_pass http://localhost:8080; # Адрес вашего Go бэкенда
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
```

Создайте символическую ссылку:

```bash
sudo ln -s /etc/nginx/sites-available/chukotka /etc/nginx/sites-enabled/
```

Перезагрузите Nginx:

```bash
sudo systemctl reload nginx
```

### 5. Настройка SSL (Let's Encrypt)

```bash
sudo apt install certbot python3-certbot-nginx -y
sudo certbot --nginx -d your_domain
```

### 6. Настройка фаервола (ufw)

```bash
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw enable
```

## Финальные шаги

Теперь ваш проект CHUKOTKA должен быть полностью настроен и доступен по указанному доменному имени!

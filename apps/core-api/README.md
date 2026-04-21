```sh
apps/core-api/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── config/          # load env, config struct
│   ├── handler/         # HTTP layer (Fiber)
│   ├── usecase/         # business logic
│   ├── repository/      # database access
│   ├── model/           # entity/domain model
│   ├── dto/             # request/response schema
│   ├── middleware/      # auth, logging
│   └── validator/       # input validation
│
├── pkg/
│   ├── database/        # DB connection
│   ├── logger/          # logging system
│   └── utils/           # helper
│
├── routes/
│   └── router.go
│
├── migrations/          # SQL migration
│
├── .env
├── go.mod
└── go.sum
```


# Alur Pengerjaan sistem
### 1. Migration
### 2. Model
### 3. DTO
### 4. Repo
### 5. Use Case
### 6. Handler
### 7. Router

# Swagger
```sh
swag init -g cmd/server/main.go
```

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
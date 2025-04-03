# minigoapi

```
minigoapi/ folder structure
├── cmd/                # Main app entry (main.go is here)
│   └── http/
│       └── main.go
├── internal/           # business logic for internal (other domains cannot access here)
│   ├── userapi/           # A sample domain
│   │   ├── handler.go      # HTTP layer
│   │   ├── service.go      # Business logic
│   │   ├── repository.go   # DB repos
│   │   └── model.go        # Data models
│   └── ...
├── pkg/               # utility codes may be used for someone else as well
│   └── logger/
│       └── logger.go
├── configs/           # Config files (.env, yml, vs.)
├── api/               # Docs like Swagger/OpenAPI, Postman
├── scripts/           # Build, deploy, migration scripts
├── docs/              # Technical documentation
├── go.mod
├── go.sum
└── README.md
```

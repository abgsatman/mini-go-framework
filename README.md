# Mini Go Framework

```
mini-go-framework/ folder structure
├── cmd/                # Main app entry (main.go is here)
│   └── http/
│       └── main.go
├── internal/           # business logic for internal (other domains cannot access here)
│   └── http/
│       └── userapi/           # A sample domain
│           ├── handler.go      # HTTP layer
│           ├── service.go      # Business logic
│           ├── repository.go   # DB repos
│           ├── model.go        # Data models
│           └── ...
├── pkg/               # utility codes may be used for someone else as well
│   └── logger/
│       └── logger.go
├── configs/           # Config files (.env, yml, vs.)
│   ├── dev.env
│   ├── prod.envv
│   └── test.env
├── api/               # Docs like Swagger/OpenAPI, Postman
├── scripts/           # Build, deploy, migration scripts
├── docs/              # Technical documentation
├── test/              # Test scripts etc...
├── assets/
│   ├── pages/
│   ├── css/
│   ├── js/
│   └── images/
├── go.mod
├── go.sum
└── README.md
```

Command
```
$ go run ./cmd/http //running for http domain
```
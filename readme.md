# Learning Go

## File structure
```
my-go-project/
├── go.mod
├── go.sum
├── cmd/
│   └── myapp/
│       └── main.go
├── pkg/
│   ├── mypackage/
│   │   └── mypackage.go
└── README.md
```

1. Creating Go Project

### NOTE: Go has
    a. Packages - A folder that contains a collection of go files
    b. Modules - Collection of packages 

### NOTE: When we are initializing a project we are really initializing a new module

- Create a folder for your project
```
- Cd <project>
- go mod init <project> / (usually) github.com/user/<project>
```
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
1. **Folder Structure:**
- project directory, named "my-go-project," contains the following files and directories:
    - `go.mod`: This file manages your project's dependencies and specifies the Go module name.
    - `go.sum`: The checksum file for your dependencies.
    - `cmd/`: A directory where you can place your application-specific commands (e.g., executables).
       - Inside `cmd/myapp/`, there's a `main.go` file, which likely contains your application's entry point.
    - `pkg/`: A directory for your custom packages.
        - Inside `pkg/mypackage/`, there's a `mypackage.go` file, which defines your custom package.

2. **Creating a Go Project:**
- To initialize a new Go project (which is essentially a new module), follow these steps:
    - Open your terminal or command prompt.
    - Navigate to the root folder of your project (in your case, the "my-go-project" folder).
    - Run the following command to initialize a Go module:
       ```
       go mod init github.com/user/my-go-project
       ```
- Replace `github.com/user/my-go-project` with your desired module path (usually in the format `github.com/username/projectname`).


# Packages and modules in Go:

1. **Packages:**
   - A package in Go is a way to organize related code files. It contains a collection of Go source files that provide a set of functions, types, and variables.
   - Packages help modularize your code, making it easier to manage and reuse. You can think of a package as a namespace for related functionality.
   - Commonly used packages (e.g., `fmt`, `os`, `net/http`) are part of the Go standard library and can be imported directly into your code.

2. **Modules:**
   - A module in Go is a higher-level unit that encompasses one or more packages. It defines a versioned collection of related packages.
   - When you create a new Go project, you're essentially initializing a new module.
   - Modules are managed using the `go.mod` file, which specifies the module name, dependencies, and their versions.
   - Modules allow you to track and manage external dependencies more effectively.

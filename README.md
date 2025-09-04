# Docker Container Inspector

### This is a simple command-line tool written in Go to inspect running and stopped Docker containers on your local machine.

#### Prerequisites
- Go (version 1.18 or later)
- Docker Desktop or Docker Engine installed and running.

#### How to RunClone or download the project files.
Place main.go and go.mod in the same directory.
Open your terminal and navigate to the project directory.
Tidy up the dependencies. This command will download the necessary Docker client library.
```go mod tidy```

Run the application.
```go run main.go```

#### What it does
The tool connects to your local Docker daemon and performs the following actions:Lists all containers (both running and stopped).

For each container, it prints detailed information, including:
- Container ID
- Name
- Image
- Current State (e.g., running, exited)
- StatusPort mappings

This provides a quick and easy way to get an overview of your containers directly from your terminal.
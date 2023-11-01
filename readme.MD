# Go gRPC User Service

This is a simple Go gRPC user service that demonstrates how to create a gRPC service for managing user data. It includes mock data for users and provides endpoints to fetch user details based on user IDs.

## Getting Started

These instructions will help you set up and run the gRPC service on your local machine.

### Prerequisites

- Go (1.17 or higher)
- Docker (if you want to run the service in a Docker container)

### Installation

1. Clone the repository to your local machine:
or use the zip file to extract the code
```bash
   git clone  https://github.com/Shalu-Kushwaha/user_micriservice.git
   cd your-repo
```
2. Resolve the dependencies

. Go to the project directory
```bash
cd user_croservice

```
. Run
```bash
 go mod tidy
 ```

### Build the Go application:
   ```bash
   go build
   ```

### Usage
1. Run the Go gRPC service:

```bash
./user_microservice
```

The service will start and listen on port 50051.
2. To interact with the gRPC service, you can use a gRPC client application.
3. If you want to run the service in a Docker container, follow these steps:

  . Build the Docker image:
   ```bash
   docker build -t your-image-name:latest .
   ```
  . Run the Docker container:
   ```bash
   docker run -p 50051:50051 your-image-name
   ```

### gRPC Endpoints
The service provides the following gRPC endpoints:

. `GetUserById`: Retrieve a user's details by ID.

. `GetUsersByIds`: Retrieve a list of user details based on a list of IDs.

### Built With
. Go - The programming language used.

. gRPC - A high-performance, open-source framework for building remote procedure call (RPC) services.

#### Authors
. Shalu Kushwaha

### License

This project is licensed under the MIT License

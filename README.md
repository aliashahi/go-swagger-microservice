# Golang Microservice with Swagger and gRPC

Welcome to the Golang Microservice project utilizing Swagger for API documentation and gRPC for communication, with Gin as the web server. This project aims to simplify the setup process for integrating Swagger documentation into a microservice architecture in Golang. It includes a Task CRUD API as a demonstration of a minimal microservice system.

## Features

- **Swagger Integration**: Easily document your APIs with Swagger UI.
- **gRPC Communication**: Utilize efficient communication between client and server using gRPC protocol.
- **Gin Web Server**: Handle HTTP requests seamlessly with Gin framework.
- **Task CRUD API**: Implement basic CRUD operations for tasks.

## Project Structure

```
.
├── client                  # Client implementation with Gin
│   ├── main.go             # Entry point for the client
│   └── ...
├── protobuf/taskpb         # Protobuf definitions
│   ├── taskpb.proto        # Protobuf file for Task service
│   └── ...
├── server                  # Server implementation with gRPC
│   ├── main.go             # Entry point for the server
│   └── ...
└── ...
```

## Setup

1. **Clone the Repository**: 
   ```
   git clone https://github.com/aliashahi/go-swagger-microservice.git
   ```

2. **Install Dependencies**:
   - Make sure you have Go installed. If not, you can download it [here](https://golang.org/dl/).
   - Install necessary Go packages:
     ```
     go mod tidy
     ```

3. **Run the Server**:
   ```
   cd server
   go run main.go
   ```

4. **Run the Client**:
   ```
   cd client
   go run main.go
   ```

5. **Access Swagger UI**:
   Open your web browser and navigate to `http://localhost:8002/api/swagger/index.html` to view Swagger documentation.

## Related Links

- [Swagger Documentation](https://swagger.io/)
- [gRPC Documentation](https://grpc.io/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)

- [Declarative Comments Format](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format)


## Protobuf Setup

To set up Protobuf for your project, follow these steps:

1. Install Protobuf Compiler:
   - For macOS:
     ```
     brew install protobuf
     ```
   - For Linux:
     ```
     sudo apt-get install protobuf-compiler
     ```

2. Install Protobuf Go Plugin:
   ```
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   ```

3. Define your Protobuf messages and services in `.proto` files in the `protobuf` directory.

4. Generate Go code from Protobuf definitions:
   ```
   protoc --go_out=. protobuf/*.proto
   ```

5. Use the generated Go files in your project.
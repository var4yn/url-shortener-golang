# URL Shortener using Go

🌐 README: [English](README.md) | [Русский](README.ru.md)

This is a pet project implementing a URL shortener service using Go. It demonstrates a clean architecture approach.

## Functionality

*   **URL Shortening:** Takes a long URL and generates a shorter, unique identifier.
*   **URL Redirection:** Redirects users from a shortened URL to the original URL.

## Architecture

The project follows a layered architecture, separating concerns for better maintainability and testability.

*   **`cmd/app/main.go`:** Entry point of the application.
*   **`internal/adapters/`:** Contains implementations for different data storage mechanisms. Currently, only an in-memory adapter is implemented.
    *   **`internal/adapters/inmemory/inmemory.go`:** In-memory repository for storing URL mappings (for testing and development).
*   **`internal/app/`:** Contains the application logic, separated into commands and queries.
    *   **`internal/app/command/`:** Defines command handlers for write operations.
        *   **`internal/app/command/create_short_url.go`:** Implements the command to create a shortened URL.
    *   **`internal/app/query/`:** Defines query handlers for read operations.
        *   **`internal/app/query/get_real_url.go`:** Implements the query to retrieve the original URL from a shortened URL.
*   **`internal/depend/state.go`:** Defines the application state, holding the command and query handlers.
*   **`internal/generator/`:** Contains components responsible for generating unique identifiers for shortened URLs.
    *   **`internal/generator/nanoid.go`:** Implements an ID generator using the `nanoid` library.
*   **`internal/ports/`:** Defines how the application interacts with the outside world.
    *   **`internal/ports/httpfiber/`:** Implements a REST API for the URL shortener service using Fiber.
        *   **`internal/ports/httpfiber/router.go`:** Defines the API routes using Fiber.

## Endpoints

*   **`POST /`**: Create a shortened URL. Takes a JSON body with the original URL. Returns a JSON response containing the generated short ID.

    **Example Request:**

    ```
    POST /

    Body:
    {
        "url": "https://github.com"
    }

    Result:
    {
        "id": "9lwSRgDIrY"
    }
    ```

*   **`GET /{id}`**: Redirect to the original URL. `{id}` is the shortened URL identifier. Redirects to the original URL.

    **Example Request:**

    ```
    GET /9lwSRgDIrY

    Result:
    {
        "url": "https://github.com"
    }
    ```

## Storage Adapters

*   [x] In-Memory
*   [ ] PostgreSQL (Planned)

## Dependencies

*   `github.com/gofiber/fiber/v3`: For building the REST API.
*   `github.com/teris-io/go-nanoid`: For generating unique IDs.
*   `github.com/stretchr/testify`: For testings.

## Running the Project
```
go run cmd/app/main.go
```
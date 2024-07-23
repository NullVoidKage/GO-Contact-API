# Contact API

This is a simple HTTP server using Golang and the Gin framework that accepts and persists "Contact" data into a PostgreSQL database.

## Setup

1. Install Go: https://golang.org/doc/install
2. Clone the repository.
3. Install dependencies:
    ```sh
    go mod tidy
    ```
4. Update the database connection string in `database/database.go`.
5. Run the server:
    ```sh
    go run main.go
    ```

## API Endpoints

### Create Contact

- **URL:** `/contacts`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
        "full_name": "Alex Bell",
        "email": "alex@bell-labs.com",
        "phone_numbers": ["03 8578 6688", "1800728069"]
    }
    ```

- **Response:**
    ```json
    {
        "data": {
            "ID": 1,
            "CreatedAt": "2023-07-23T12:00:00Z",
            "UpdatedAt": "2023-07-23T12:00:00Z",
            "DeletedAt": null,
            "full_name": "Alex Bell",
            "email": "alex@bell-labs.com",
            "phone_numbers": null,
            "phones": [
                {
                    "ID": 1,
                    "CreatedAt": "2023-07-23T12:00:00Z",
                    "UpdatedAt": "2023-07-23T12:00:00Z",
                    "DeletedAt": null,
                    "ContactID": 1,
                    "number": "+61385786688"
                },
                {
                    "ID": 2,
                    "CreatedAt": "2023-07-23T12:00:00Z",
                    "UpdatedAt": "2023-07-23T12:00:00Z",
                    "DeletedAt": null,
                    "ContactID": 1,
                    "number": "+611800728069"
                }
            ]
        }
    }
    ```

## Notes

- The phone numbers are stored in E.164 format.
- Only valid Australian phone numbers are accepted.
# GO-Contact-API

![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)

# Indications

## Why did you build it?
I developed this project as a learning exercise to deepen my understanding of the Go programming language and its application in building robust REST APIs. I gained hands-on experience in implementing key API features such as authentication, data retrieval, and data submission. I also explored best practices for error handling, security, and database integration. The process allowed me to enhance my skills in designing scalable and efficient API architectures.

## Description
This REST API, developed in Go, provides secure and reliable access to metering data for electrical devices, utilizing token-based authentication.

__Authentication Process:__

- Users authenticate through the `/api/auth` (POST) endpoint.

- On successful authentication, a JSON Web Token (JWT) is issued, valid for 120 seconds. This token must be included in the headers of all subsequent API requests.

- If authentication fails, the response includes a JSON message: 
    ```json
    {"message": "Invalid login or password"}
    ``` 
    and an HTTP `401` status.

- If the JWT generation fails, the response includes: 
    ```json
    {"message": "Couldn't generate token"}
    ```
    and an HTTP `500` status.

 __Core Functionalities:__

 1. __Retrieve Meter Data:__            
    - Endpoint: `/api/getObjectsByNumLS` (GET).
    - Retrieves device details and the latest metering data using a provided `num_ls` parameter.
    - On success, returns a JSON array containing information such as device address, type, last meter reading, etc.
    - Possible errors:
      -  Missing token (HTTP 401): 
            ```json
            {"error": "Missing authorization token"}
            ```
      - Invalid token (HTTP 401): 
        ```json
        {"error": "Invalid token"}
        ```
      - Non-existent num_ls  (HTTP 404):
        ```json
        {"error": "LS [num_ls] not found"}
        ```
      - Database error: (HTTP 500) 
        ```json
         {"error": "Database error"}
        ```
2. __Submit Meter Readings:__
   - Endpoint: `/api/setMeterIndicationByGuid` (GET)
   - Allows users to submit new metering data by providing device_guid, metering, and number_request.
   - On success, returns (HTTP 200): 
        ```json
        {"id": "operation_id", "message": "true"}
        ``` 
   - Possible errors:
      - Missing or invalid token, or non-existent device_guid, resulting in similar error patterns as described above.

This API makes sure that communication for managing electrical device metering data is always secure and efficient.

## Features
 __1. Return Operation Identifier__

- __Endpoint__: /api/setMeterIndicationByGuid (GET)

- __Description__: Adds functionality to return an operation identifier (operation_id) in the JSON response. This allows users to track and manage submitted meter readings efficiently. Because now the operation identifier returned constant value.

__2. Meter Value Validation__
- __Description__: Implements validation logic to ensure that submitted meter values are accurate and fall within acceptable parameters. The implementation of this feature is under consideration for integration into either the core application layer or the database layer.
- __Benefit__: Enhances data integrity by verifying the correctness of the meter values.
  
__3. Versioned API Design (v2)__
- Introduces an updated version of the API (v2) that accepts metering data as a structured JSON object in the request body. This improves data handling and offers more flexibility for users.

## Usage
Follow these steps to configure, launch, and interact with the application effectively:

__1. Set Up the Database__
- Create a PostgreSQL database named `Indication`.
- Execute the SQL scripts located in the `datafordb` folder to initialize the required tables and schema.
- Ensure the database is properly structured and ready for the application to connect and operate.
  
__2. Configure the Application__
- Locate the `your_main.yaml` file in the folder `configs`  and update its content with your desired configuration values, such as:
    - Server configuration
    - Database connection details
- Once the file has been updated, rename it to `main.yaml`. This file contains all essential settings required for the application to function properly.

__3. Start the Application__
- Run the following command in your terminal __`go mod tidy`__ to download all models in application which it can determine from reading only the go.mod files.
- Run the following command in your terminal to start the server
__`go run main.go`__

__4.Interact with the API__
- __Using Swagger UI__  
    Open your browser and navigate to:
    __`http://localhost:port_form_main.yaml/swagger/index.html`__
    This will provide an interactive interface for testing the API endpoints.
    In the section `Available authorizations` username write `test@test.com` and password `test1` 
- __Using Postman or cURL__ :  
  Create a new request in Postman or use cURL to access the authentication endpoint:
  __`http://localhost:port_form_main.yaml/api/auth`__
  ```bash
  curl -X POST http://localhost:port_form_main.yaml/api/auth -u test@test.com:test1
  ```
  __Access Protected Endpoints__  
  After obtaining the JWT token from the /api/auth endpoint, include it in the Authorization header for all subsequent requests to access protected API methods.  
Example request with JWT token:
  ```bash
  curl -X GET http://localhost:8000/api/getObjectsByNumLS/11112 -H "Authorization:<JWT_TOKEN>"
  ```
By following these steps, you can successfully set up, start, and interact with the application.

## Technologies Used
- Programming Language: GO
- Web and API Frameworks:
    - __Gorilla Mux__: HTTP router
    - __UrFave CLI__:Command-line interface
- Swagger and Documentation:
   - __Swaggo__: Swagger documentation
   - __Http-Swagger__:Serve Swagger UI
- Database: 
   - __PostgreSQL__
   - __SQLx__: SQL with extensions
   - __pq__: PostgreSQL driver
   - __SQLMock__: Mock SQL driver for testing
- Logging and Observability:
  - __Logrus__: Structured logging
  - __Lumberjack__: Rolling logs
- Utilities:
    - __Testify__: Testing utilities
- Tools: Git / Docker 
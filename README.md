# InstaShopAssessment
INSTASHOP ASSESSMENT

To run the application, run `go run main.go`

# Implementation Details
## Language and Framework
The language is Go and the framework is Gin.

## Environment Variables
A `.env.example` file is provided as a model of the required environment variables.
- Create a `.env` file with the required variables to run the application.
- Create a `.env.test` file with the required variables to run the test suite.

## Database
The database is a PostgreSQL database and the ORM is GORM.

## Authentication
- The authentication is done using JWT tokens.
- The token is valid for 24 hours
- The token is sent in the header of the request and the format is "Bearer <token>"
- Hashing is done using Argon2.
- Refresh endpoint, logout endpoint, and token rotation/invalidation are not implemented

# Features
## User
There are signup and login endpoints. In signing up, the `role` field can only either be
- "regular": for normal users
- "admin": for admin users

## Products
- The endpoints to get product is open to authenticated users
- The create, update and delete endpoints are admin only access

## Order
- Authenticated users can create, get order and cancel an order
- Only an admin can update the status of an order

# Documentation
The documentation is done with Swagger. You can find it available at http://localhost:8080/docs/index.html

# Testing
To run the test suite, run `go test ./tests`

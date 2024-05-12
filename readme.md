# CRM-Backend-Project

Project for Udacity's Golang course - JOSHUA CALLARY

# Usage

## Required Packages

    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/gorilla/mux"

## Running the API server

To run the server, run:

```
go run main.go
```

# API Endpoints

## GET /customers

Retrieve all customer data.

## GET /customers/{id}

Retrive data for a customer

## POST /customers

Create a new customer

Request (JSON) Body fields:

- **Name**: (required, string): Customer's name
- **Role**: (required, string): Customer's role
- **Email**: (required, string): Customer's email address
- **Phone**: (required, number): Customer's phone number
- **Contacted**: (required, boolean): True if customer has been contacted

## UPDATE /customers/{id}

Update an exiting customer

Request (JSON) Body fields:

- **Name**: (required, string): Customer's name
- **Role**: (required, string): Customer's role
- **Email**: (required, string): Customer's email address
- **Phone**: (required, number): Customer's phone number
- **Contacted**: (required, boolean): True if customer has been contacted

## DELETE /customers/{id}

Deletes a customer
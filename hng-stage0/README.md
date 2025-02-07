# HNG Stage 1 Backend Task

# Number Classification API

This API classifies a given number and returns interesting mathematical properties about it, along with a fun fact.

## Technology Stack

- Go

## API Specification

### Endpoint

`GET /api/classify-number?number={number}`

### Response Format

#### 200 OK
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```
[Go Lang Devs]("https://hng.tech/hire/golang-developers")
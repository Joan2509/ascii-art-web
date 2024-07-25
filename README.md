# Ascii-art-web

## Description

Ascii-art-web is a web application that provides a graphical user interface (GUI) for creating ASCII art banners. The application allows users to input text and select a banner style (shadow, standard, or thinkertoy) to generate an ASCII art banner.


## Installation and Setup

To install, clone the repository locally using Git:
```go
git clone https://learn.zone01kisumu.ke/git/jwambugu/ascii-art-web.git
```

Alternatively, download the project directly from Gitea and access it through your file manager.

## Usage

To run the application:


1. Navigate to the project directory:
```go
cd ascii-art-web
```

2. Run the application:
```go
go run main.go
```

3. Open a web browser at the provided port: [http://localhost:8080](http://localhost:8080)

## Implementation

The application is implemented in Go programming language. It creates an HTTP server that handles GET and POST requests. HTML templates are used to display data to the user.

### Algorithm

1. User sends a GET request to the root URL (/) to retrieve the main page.
2. Server responds with an HTML template including a text input, radio buttons for banner styles, and a submit button.
3. When the form is submitted, the client sends a POST request to /ascii-art endpoint with text and banner style as form data.
4. Server processes the request, generates ASCII art banner using the selected style, and responds with an HTML template displaying the banner.

## Code Structure

The code is organized into the following directories:

- `ascii`: contains the ASCII art generation
- `banner`: contains resources related to banner styles 
- `handlers`: contains server endpoint handlers and related logic for HTTP requests and responses.
- `templates`: contains HTML templates for main page and ASCII art banner display


## HTTP Endpoints

- `GET /`: returns the main page HTML template
- `POST /ascii-art`: processes form data and generates ASCII art banner

## HTTP Status Codes

- `200 OK`: returned when request is successful
- `404 Not Found`: returned when requested resource is not found
- `400 Bad Request`: returned when request is invalid
- `500 Internal Server Error`: returned when an unhandled error occurs

## Notes

- Application uses Go templates for displaying data to the user.
- Application uses form data to send text and banner style to the server.
- Application uses HTTP server to handle GET and POST requests.

## Authors

- [Joan Wambugu](https://github.com/Joan2509)
- [Cherrypick14](https://github.com/Cherrypick14)
- [Raymond](https://github.com/anxielray)


# Snippetbox

Snippetbox is a secure web application that allows developers to store, manage, and share code snippets privately. This project demonstrates idiomatic Go programming, secure coding practices, and key web development features like user authentication and dynamic templating.

## Live Demo

Visit [https://snippetbox.alwaysdata.net/](https://snippetbox.alwaysdata.net/) to see the application in action. (Note: Sign-in required to view snippets)

## Key Features

- **Secure User Authentication:** Implements signup, login, and logout functionality with secure password hashing (bcrypt) and CSRF protection.
- **SQL Injection Prevention:** Uses parameterized queries and SQL transactions to ensure secure database interactions.
- **Dynamic HTML Templating:** Renders HTML pages with dynamic data using Go templates, including template inheritance for maintainability.
- **Custom Middleware:** Employs custom middleware for logging, security headers (X-Content-Type-Options, etc.), panic recovery, and request handling.
- **RESTful Routing:** Structures the application using a RESTful approach.
- **HTTPS & TLS:** Configures the application to use HTTPS with TLS encryption.
- **Efficient Database Queries:** Utilizes connection pooling and prepared statements for optimal database performance.

## Requirements

- Go 1.21+
- MySQL 8.0+

## Installation

1. Clone the repository:

    `git clone https://github.com/Devkumar194Aghera/snippetbox.git`

2. Navigate to the project directory:

    `cd snippetbox`

3. Install dependencies:

    `go mod download`

4. Set up your MySQL database and update the connection details in the configuration file.

5. Build and run the application:

    `go run ./cmd/web/.`


## Usage

1. Create an account or log in at [https://snippetbox.alwaysdata.net/](https://snippetbox.alwaysdata.net/).
2. Click "New Snippet" to add a new code snippet. Provide a title, content, and expiration time.
3. View, edit, or delete your snippets from your dashboard.
4. Snippets are private by default, ensuring your code remains secure.

## Security Measures

- Secure password hashing using bcrypt
- CSRF protection for all POST requests
- SQL injection prevention through parameterized queries
- Secure session management with SameSite and HTTP-only flags
- Implementation of various security headers (X-Frame-Options, X-XSS-Protection, etc.)

## Project Structure

The project follows a clean, modular structure:

- `cmd/web`: Main application entry point
- `internal/models`: Database models and queries
- `internal/validator`: Input validation logic
- `ui/html`: HTML templates
- `ui/static`: Static assets (CSS, JavaScript)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Images

### Project Screenshots

- **Website View:**  
  ![image](https://github.com/user-attachments/assets/654be7ae-e0fa-4b3f-a746-0f5aadb70048)

### Traffic Metrics

- **Best Traffic Received:**  
  ![number-of-visits](https://github.com/user-attachments/assets/5568da39-a634-41cc-abde-a02b5a3aff45)



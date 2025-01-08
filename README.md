# **Snippet Box**

This project implements a simple web application that allows users to store, manage, and share code snippets. The project demonstrates the use of idiomatic Go programming practices, structuring applications in a maintainable and secure way, and integrating common web development features.

## Features

- **Configuration Settings** : Set configuration options at runtime using command-line flags.

- **Enhanced Logging**: Improve log messages by categorizing them by type (e.g., info, error) for better debugging and monitoring.

- **Dependency Management**: Make dependencies available to handlers in a type-safe, extensible manner, making testing and scalability easier.
Centralized Error Handling: Handle errors consistently across the application to avoid repetitive code.

- **HTML Rendering**: Render HTML pages and use template inheritance to keep markup DRY and maintainable.

- **Static File Serving**: Serve images, CSS, and JavaScript files from the application.

- **SQL Transactions & Security**: Use transactions to group SQL statements and prevent SQL injection attacks with safe parameterized queries.

- **Template Handling**: Dynamically pass data to HTML templates and handle template rendering errors gracefully.

- **Custom Middleware**: Implement middleware for logging, security headers, panic recovery, and request handling.

- **RESTful Routing**: Structure your application using a RESTful approach with a third-party router.

- **Form Handling**: Parse and validate form data with user-friendly feedback for validation errors.

- **Session Management**: Use sessions to securely store user data and handle session timeouts and cookie settings.

- **TLS Support**: Set up HTTPS with self-signed certificates and adjust TLS settings for security.

- **User Authentication**: Implement basic signup, login, and logout functionality with secure password storage and CSRF protection.

- **Request Context**: Pass contextual information (like user data) between handlers using request context.

## Requirements

- Go 1.16 or higher
- MySQL or compatible database server
- A web browser to test the application

## Installation

1. Clone the repository:

    `git clone https://github.com/Devkumar194Aghera/snippetbox`

2. Navigate to the project directory:

    `cd snippet-box`

3. Install the necessary dependencies:

    `go mod tidy`

4. Configure your MySQL database connection in the "config" file.

5. Run the application:
    `go run cmd/web/*`

6. Visit http://localhost:4000 in your web browser.


## Usage

- **Store Snippets**: Use the web interface to add new code snippets with titles, descriptions, and code.

- **Edit Snippets**: Edit existing snippets directly from the user interface.

- **Delete Snippets**: Remove snippets when no longer needed.

- **User Authentication**: Users can sign up, log in, and log out with secure sessions.

- **Access Control**: Users will only be able to view their own snippets unless otherwise specified.

## Security Features
- **SQL Injection Prevention**: Safe queries are executed using Go's database/sql package with parameterized queries.

- **Cross-Site Request Forgery (CSRF)**: CSRF protection is in place to prevent malicious form submissions.

- **Password Security**: Passwords are encrypted using Bcrypt before being stored in the database.

- **TLS Encryption**: The application serves all traffic over HTTPS with secure TLS settings.


## Middleware 

- **Logging**: Logs all incoming HTTP requests for monitoring.

- **Panic Recovery**: Catches panics and gracefully recovers, preventing crashes and providing error information.

- **Security Headers**: Adds HTTP headers to improve security, such as X-Content-Type-Options, X-Frame-Options, and Strict-Transport-Security.

## Template Features
- **Template Caching**: Templates are cached to avoid redundant reading from the disk, improving performance.

- **Dynamic Data**: Pass dynamic data to templates in a type-safe way.

- **Custom Template Functions**: Create custom functions to format or transform data in your templates.


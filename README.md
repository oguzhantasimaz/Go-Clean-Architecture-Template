# Golang Backend Clean Architecture Template

This is a template repository for building a Golang backend application following the Clean Architecture principles. It includes the following features:

- **Google Auth**: Integration with Google OAuth for user authentication.
- **JWT Auth**: Token-based authentication and authorization system.
- **MySQL Connection**: Database connectivity to MySQL for data storage.
- **User CRUD Operations**: Basic user management operations (Create, Read, Update, Delete).

**You can use this project as a template to build your Backend project in the Go language on top of this project.**

**Edit the project to suit your needs after you use it as a template.**


Please ⭐️ this repository if this project helped you, also contributions are welcome 🙏🏼

## Getting Started

Follow the steps below to set up and start using this template for your Golang backend project.

### Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install)
- [MySQL](https://dev.mysql.com/downloads/installer/) (or a compatible database)
- [Google Cloud Console](https://console.cloud.google.com/) project for Google Auth setup.

### Installation
There are two ways to start using this template:

1. Create a new repository based on this template by clicking the **Use this template** button at the top of this page.


----


1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/oguzhantasimaz/Go-Clean-Architecture-Template.git
   cd your-repo-name
   ```

2. Create a `.env` file in the root directory based on the `.env.example` template and fill in your configuration details.

3. Install the necessary Go packages:

   ```bash
   go mod tidy
   ```

### Configuration

1. Set up your Google OAuth credentials in the Google Cloud Console, and update your `.env` file with the corresponding client ID and client secret.

2. Configure your MySQL database connection in the `.env` file with the appropriate credentials.

### Usage

Run the following command to start the Golang backend server:

```bash
go run main.go
```

Your server should now be running at port you specified in the `.env` file.

## Project Structure

```
├── api/
│   ├── controller/
│   ├── middleware/
│   └── route/
├── bootstrap/
├── cmd/
├── domain/
├── internal/
│   ├── tokenutil/
├── repository/
├── usecase/
├── utils/
└── main.go
```

## About Me
Hello, I am Oguzhan Tasimaz, a software engineer from Turkey. I am interested in backend development, distributed systems, and cloud computing.
You can find me on [LinkedIn](https://www.linkedin.com/in/oguzhantasimaz).

## Contributing

Feel free to contribute to this project.

## If this project helps you in anyway, show your love ❤️ by putting a ⭐ on this project ✌️

### Contributing to Go Backend Clean Architecture

All pull requests are welcome.

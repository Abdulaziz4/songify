# Songify
> This project serves as an educational resource to gain familiarity with Docker, Kubernetes, and AWS deployment. 

Songify is a Go-based RESTful API that allows you to manage songs. You can create, edit, delete, and retrieve songs using this API. It is built using the Gin framework and utilizes PostgreSQL as the database for storing song information. The API is designed to provide a simple and efficient way to manage your song library.

## Technology Stack

- Go: The programming language used to build the API.
- Gin: A lightweight and fast HTTP web framework for building APIs in Go.
- PostgreSQL: A powerful and popular open-source relational database used for storing song data.
- Docker: The API is containerized using Docker, allowing for easy deployment and scalability.
- AWS EC2 and EBS: The API is deployed on an AWS EC2 instance with EBS for storage.

## Getting Started
To get started with the Songify, follow the steps below:

### Prerequisites
- Go 1.16 or later
- docker

### Installation

1. Clone the repository:

```sh
git clone <repository-url>
```


2. Install the required dependencies:
```sh
go mod download
```

3. Start development environment:
```sh
docker compose up
```

## API Endpoints
The following endpoints are available in the Songify API:

- GET /songs: Retrieves a list of all songs.
- GET /songs/{id}: Retrieves a specific song by its ID.
- POST /songs: Creates a new song. Requires a JSON payload with song details.
- PUT /songs/{id}: Updates an existing song by its ID. Requires a JSON payload with updated song details.
- DELETE /songs/{id}: Deletes a song by its ID.

## Deployment
*WIP*

## License
This project is licensed under the MIT License.
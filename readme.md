# Songify
> This sample project showcases the integration of key AWS services to create a robust application infrastructure. It leverages AWS RDS for the database, AWS ECR for container image management, AWS Secret Manager for secure secrets storage, and AWS EKS for deploying the application as a Kubernetes cluster.

Songify is a Go-based RESTful API that allows you to manage songs. You can create, edit, delete, and retrieve songs using this API. It is built using the Gin framework and utilizes PostgreSQL as the database for storing song information. The API is designed to provide a simple and efficient way to manage your song library.

## Technology Stack

- Go: The programming language used to build the API.
- Gin: A lightweight and fast HTTP web framework for building APIs in Go.
- PostgreSQL: A powerful and popular open-source relational database used for storing song data.
- Docker: The API is containerized using Docker, allowing for easy deployment and scalability.


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
**AWS RDS (Relational Database Service):** Utilize AWS RDS as the primary database solution, ensuring efficient data management and persistence.

**AWS ECR (Elastic Container Registry):** Push automated container images to AWS ECR, a secure and scalable container registry, simplifying image management and deployment.

**AWS Secret Manager:** Store and manage sensitive information securely in AWS Secret Manager, ensuring secure access to confidential data without exposing them directly in the codebase.

**AWS EKS (Elastic Kubernetes Service):** Deploy the application as a Kubernetes cluster using AWS EKS, a managed Kubernetes service, ensuring high availability, fault tolerance, and scalability.

## License
This project is licensed under the MIT License.
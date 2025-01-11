# Application Configuration

This section explains how to configure and run the application using environment variables. The application uses
structured configuration via the `ApplicationEnv` struct, which is divided into application-level and database-level
settings.

## Application Environment Variables

The application reads its configuration from the following environment variables:

### General Configuration (`ApplicationEnv`)

| Variable                   | Description                                                             | Default Value              | Required? |
|----------------------------|-------------------------------------------------------------------------|----------------------------|-----------|
| `API_SECRET_AUTHORIZATION` | Secret used to encrypt and decrypt API token claims.                    | N/A                        | Yes       |
| `API_SECRET_KEY`           | The api private key used for API security.                              | N/A                        | Yes       |
| `PORT`                     | The port on which the application runs.                                 | `1080`                     | No        |
| `CONTEXT_PATH`             | The base path used for API routing.                                     | `/api/private-network/v1/` | No        |
| `PEERS_RESOURCE_PATH`      | Filesystem path for peer resource connections.                          | `/etc/wireguard/`          | No        |
| `API_INIT_FILE`            | Add the first user data that will be created at the application startup | N/A                        | No        |

### Database Configuration (`DBEnv`)

| Variable         | Description                                               | Required? |
|------------------|-----------------------------------------------------------|-----------|
| `MONGO_DATABASE` | Name of the MongoDB database the application connects to. | Yes       |
| `MONGO_USERNAME` | Username for MongoDB authentication.                      | Yes       |
| `MONGO_PASSWORD` | Password for MongoDB authentication.                      | Yes       |
| `MONGO_URI`      | MongoDB connection URI.                                   | Yes       |

## Default Behavior

- The default **server port** is `1080` if `PORT` is not specified.
- The default **API context path** is `/api/private-network/v1/` if `CONTEXT_PATH` is not provided.
- The default path for **peer resources** is `/etc/wireguard/` if `PEERS_RESOURCE_PATH` is not provided.

## Example `.env` File

Here is an example `.env` file for setting up the environment variables:

``` env
# Server Configuration
API_SECRET_AUTHORIZATION=my-secret-for-auth
API_SECRET_KEY=my-api-secret-key
PORT=3000
CONTEXT_PATH=/custom/api/path/
PEERS_RESOURCE_PATH=/config/peers/

# Database Configuration
MONGO_DATABASE=my-app-database
MONGO_USERNAME=my-db-user
MONGO_PASSWORD=my-db-password
MONGO_URI=mongodb://localhost:27017
```

## Startup app with docker

- **Note**: The following commands assume you have Docker installed on your machine and use the `./dev/.env` file that
  is not recommended for **production**.

```shell
make rebuild
```

## Startup app with a custom `.env` file

- Check the [`./dev/.env`](./dev/.env) file for the environment variables

```shell
make rebuild ENV=<< .env file path >>
```

## Configuration Initialization

The application initializes by loading environment variables during startup. Here’s how the variables work:

1. **General Configuration**:
    - Variables such as `API_SECRET_AUTHORIZATION` and `API_SECRET_KEY` are essential for securing API operations.
    - Optional fields like `PORT` and `CONTEXT_PATH` fallback to default values if not provided.

2. **Database Configuration**:
    - All `MONGO_*` variables (e.g., database name, URI, username, and password) are required to establish a successful
      MongoDB connection.
    - If a variable is missing, application startup will likely fail.

## Deployment Recommendations

- **Sensitive Variables**: Avoid committing sensitive information (e.g., secrets, API keys, passwords) to version
  control. Use tools like `.env` files (local development) or secret management systems (production).
- **Default Paths**: Make sure the paths, such as `PEERS_RESOURCE_PATH`, align with your intended filesystem structure
  in production environments.
- **Error Handling**: Ensure you have runtime checks for missing or misconfigured environment variables to prevent
  application crashes.

This documentation is concise, yet it contains all the necessary details for configuring and running the application.
Let me know if you need further adjustments!

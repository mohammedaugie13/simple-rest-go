# Bitwyre Golang Project Boilerplate

This is Go Clean Architecture Template with gin framework as HTTP Server, gqlgen as GraphQL Server, gorm as orm for database related operations.

This project already included Sample Rest API and GraphQL Query:
- `http://localhost:3000/healthc` Sample Health Check
- `http://localhost:3000/graph/play` GraphQL Playground
- `http://localhost:3000/graph` GraphQL Server

Read about Clean Architecture here: https://betterprogramming.pub/the-clean-architecture-beginners-guide-e4b7058c1165?gi=7d1db8c02f61

## Running Project
If you are first time with this project, execute `./run.sh setup` to init the project and download some dependencies. Then follow the next command below:
- Copy `.env.example` to `.env`
- Run `docker-compose up`
- Everything will be set up, go to `localhost:3001`

## Generate GraphQL
To define query, resolver , schemes please read official docs: https://github.com/99designs/gqlgen

To generate the server please run: `gqlgen generate`

## Available Command

You can use this available command
- To start the app with hot reloading use `./run.sh hot-serve`
- To start the app without hot reloading `./run.sh serve <args>`

Available args

| Command     | Options          | Description                                    |
|-------------|------------------|------------------------------------------------|
| --env       | local, dev, prod | Define specific environment                    |
| --migration | y, n             | Wheter you want to migrate the database or not | 
| --seeds     | y, n             | Execute database seeder                        |

For another command you can pas `--help`

## Feature
- HTTP Rest Server with GIN Framework
- GraphQL Server and Playground with GQLGEN
- Centralized config files via `.env` and controlled by `config.go`
- Logging with logrus
- Postgres Database (Including Auto Migration and Seeder)
- GORM as ORM
- Already included Sample Rest API and GraphQL Query
- Dockerized Project

## Sample GraphQL Query

Query:
```graphql
query($input: InfoQueryInput){
  infoQuery(input: $input){
    isEmailVerified
    isRegistered
    email
    actionToken
  }
}
```

Variables
```graphql
{
	"input": {
		"email": "dock@collins.biz"
	}
}

```
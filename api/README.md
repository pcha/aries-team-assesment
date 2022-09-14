# API 
## Setup
### Modules
`go get ./...`

### Database
This project uses a mysql db, you can start a local database using the docker-compose.yml included in this repo. 
The initial data (including tables) for this project is defined in ./init.sql. The docker-compose.yml will start up 
running this init.sql file

### Environment
The vars you need are defined in `sample.env`, you just need to copy this file and rename it to `.env`. Keep in mind that 
the docker-compose.yml file is using this file to inject variables too.

## Start
`go run main.go`
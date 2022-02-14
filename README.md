# ddd-value-objects-in-postgres

## Prerequisites

You need Docker and docker-compose installed on your computer in addition to psql or similar.

## How to use

1. Run docker-compose to start Postgres on your computer `docker-compose -f build/docker-compose.yaml up -d`.
1. Connect to the database by running `psql -h localhost -U tobbstr -d ddd`.
1. Run the migration script manually by copy pasting it into the terminal.
1. Run the integration tests in the `./code` folder by entering `go test ./code/... -v`.
1. Note that the tests pass which proves that it's possible to both store and retrieve JSON data from a JSONB column in Postgres.
1. Run `docker-compose -f build/docker-compose.yaml down -v` to shut down the Postgres instance.

## Notes

Please read through the code files in the `./code` package. They contain comments that explain what they are for.

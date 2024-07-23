# MQTT Studio

Website: [mqtt.studio](https://mqtt.studio)

## API Requirements

- [golang-migrate](https://github.com/golang-migrate/migrate) should be installed on your local machine in order to use migrations.

## API Usage

- Copy `.env.example` as `.env` and configure it.
- `go get` to install go packages.
- `make migrate-up` to run migrations.
- `make run` to run the API.

# :trophy: Go-Trophy Server [![CircleCI](https://circleci.com/gh/hirsch88/go-trophy-server.svg?style=svg)](https://circleci.com/gh/hirsch88/go-trophy-server)

## ❯ Why

TODO

## ❯ Installation

```shell
brew install go
brew install glide
brew install goose
brew install mysql
```

## ❯ Getting Started

First we need to install the dependencies.
```shell
make install
```


To start the server run.
```shell
make run
```


## ❯ Dependency Management

Our dependency management tool is [glide](https://glide.sh/);

### Add Dependency

```shell
glide get github.com/foo/bar
```

### Updating Dependencies

```shell
glide update
```

## ❯ Database Migrations

To manage our database schema we used the library [goose](https://github.com/pressly/goose).

### Run migrations

```shell
make db-migrate
```

### Create a migration

```shell
make db-create-migration name=create_user_table
```

## ❯ Testing

For testing our awesome app we decided to use the toolkit [stretchr/testify](https://github.com/stretchr/testify).

### Create Mocks

To automatically generate the mocks we use [vektra/mockery](https://github.com/vektra/mockery). 
The below command example generates the mock files of our interfaces inside the /app folder.

```shell
make mocks
```

## ❯ License

[MIT](/LICENSE)

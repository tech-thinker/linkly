<div align="center">
  <img src="views/assets/favicon.png" width="100px" alt="logo" />
  <h1><code>linkly</code></h1>
  <p>
    <strong>
    A simple, fast, and powerful link shortening service.
    </strong>
  </p>
</div>

## Requirements

- [Golang](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/) (Optional but recommended)
- [Postgresql](https://www.postgresql.org/download/) (**Only if you are not using `docker`**)

## Features

- Shorten any url
- View stats of any url

## Development

To get started, you can clone the repository and run the following commands:

```bash
git clone https://github.com/tech-thinker/linkly.git
```

To install the dependencies, run the following commands:

```bash
cd linkly
go mod download
```

Copy environment variables file and replace the values with your own.

```bash
cp .env.example .env
```

Generate OpenAPI spec:

```bash
swag init --parseDependency --parseInternal
```

To run the application, run the following commands:

```bash
export $(cat .env | xargs)
go build
./linkly
```

## Contributing

To contribute, please open an issue or pull request.

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Author

- Injamul Mohammad Mollah <mrinjamul@gmail.com>

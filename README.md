# Tech Service

Federated graphQL sub-graph written schema-first with gqlgen. This service exposes queries for technologies.

## Setup

This service makes heavy use of docker and docker compose. Setting up docker is beyond the scope of this readme.

```shell
mkdir ~/projects/dudo-website
cd ~/projects/dudo-website
git clone git@github.com:dudo-website/tech-service.git
cd ~/projects/dudo-website/tech-service

docker compose up -d
```

## Development

gqlgen provides a schema-first approach to graphQL development. Start with `./graph/schema.graphqls`, and modify the schema as necessary.

To generate code:

```shell
docker compose run --rm gen
```

## Deployment

This should be done for you via CI/CD, but in case you need to manually push an image

```shell
docker build -t docker.pkg.github.com/dudo-website/tech-service/tech-service:0.1.0
docker push docker.pkg.github.com/dudo-website/tech-service/tech-service:0.1.0
```

## Reading

Learn about the various tech powering this application:

- [GraphQL](https://graphql.org)
- [gqlgen](https://github.com/99designs/gqlgen)
- [go](https://go.dev)
- [GraphQL Federation](https://www.apollographql.com/docs/federation)
- [Docker](https://docs.docker.com/compose/gettingstarted)

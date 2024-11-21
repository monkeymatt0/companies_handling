# companies_handling

To use start the application you need to set:

```
cmd/.env -> There is an example file
.env -> There is an example file
```

## cmd/.env

You need to set this env variable in order to make it work.
GIN_MODE should be setted to release
TRUSTED_PROXY should contain the IP of the reverse proxy, if used(the code will adjust accordingly on strat up phase)

```
POSTGRES_USER=<USER>
POSTGRES_PASSWORD=<PASSWORD>
POSTGRES_DB=<DB>
DB_HOST=<DB_HOST> This on here should be "db" since we are using the service in the docker compose
DB_PORT=<DB_PORT>

SECRET=<SECRET>
GIN_MODE=<MODE> # can be "debug" or "release"
TRUSTED_PROXY= # Here is possible add trusted proxy
```

## .env

These right here are the same of the above file only that this file has been placed in the root of the folder to remove the warning from the postrges image

```
POSTGRES_USER=<USER>
POSTGRES_PASSWORD=<PASSWORD>
POSTGRES_DB=<DB>
DB_HOST=<DB_HOST>
DB_PORT=<DB_PORT>
```

Once these env variable are setted you can start the application by running:
`docker compose build && docker compose up -d`
this command needs to be ran at the root of the project folder.

## Endpoints

```
POST   /users
POST   /login
GET    /users/:id
DELETE /users/:id
DELETE /users/:id/hard
GET    /users/:id/companies/:uuid
POST   /users/:id/companies
PATCH  /users/:id/companies/:uuid
DELETE /users/:id/companies/:uuid
DELETE /users/:id/companies/:uuid/hard
```

## User POST JSON example

All the user related post endpoint will use this JSON obect

```
{
    "email": "test@testing.com",
    "password": "test12345"
}
```

## Company POST/PATCH JSON example

```
{
  "name": "Tech Inc.",
  "description": "A company focused on cutting-edge technology solutions.",
  "amountOfEmployees": 250,
  "registered": true,
  "type": "Corporations"
}
```

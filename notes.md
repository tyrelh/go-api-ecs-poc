# Progress

## Local database running in Docker

Using the official MySQL image
```bash
docker pull mysql:8.0
```
Not sure why I didn't start with the latest version??

`MYSQL_ROOT_PASSWORD` env variable needs to be set in Docker image container for MySQL8


## GORM

[GORM](https://gorm.io/docs/index.html) is the Go ORM. It's the most popular ORM solution for Go.

Able to connect to the database running locally. Have it currently setup on port 5000.

Created effectively an admin user to use.

Using the root user, run:
```sql
CREATE USER 'go'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'go'@'%';
flush privileges;
```


### Models

GORM can use your structs as database Models. https://gorm.io/docs/models.html

You can optionally inject `gorm.Model` into your struct to get the automatic default fields `ID`, `CreatedAt`, `UpdatedAt`, and `DeletedAt`.


### Deleting

It seems like if you have the `DeletedAt` field on your model, that will automatically enable soft-delete for that table. When deleting a timestamp gets set in the `DeletedAt` field and then this object is not returned in other queries, though it still remains in the database.

I think if you don't have the `DeletedAt` field, hard-delete is the default.


### Auto Migration

You can call the `AutoMigrate` function on the database connection, passing the models you're interested in, to force modification or creation of the database schema


## OpenAPI & oapi-codegen


### OpenAPI

The OpenAPI 3.0 specification docs seem easy enough to follow: https://swagger.io/docs/specification/v3_0/basic-structure/

One gotcha that caught me initially is that `number` in the spec is float32. There's a separate `integer` type for signed ints.


#### Schemas

You can define schemas in-line in each response for request and response format, but more usefully you can define them separately in the `components:` block so that you can reuse them across endpoints. This also gives you control over their naming which can help keep them generic or directly related to internal data models instead of specific endpoints.

Example:
```yaml
components:
  schemas:
    RewardResponse:
      type: object
      properties:
        id:
          type: integer
          example: 1
        brand:
          type: string
          example: Pokemart
        currency:
          type: string
          example: PKD
        denomination:
          type: number
          example: 1000.00
    RewardRequest:
      type: object
      properties:
        brand:
          type: string
          example: Pokemart
        currency:
          type: string
          example: PKD
        denomination:
          type: number
          example: 1000.00
```


## oapi-codegen

[oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) is a tool used to generate Go backend or client boilerplate code given an OpenAPI specification.

In this project we're using it to generate server-side backend boilerplate for the standard _net/http_ server.


### Command

```bash
oapi-codegen --config=api/oapi-codegen.yml api/api.yml
```
The two args here are the config file flag, and then the input OpenAPI specification.


### Code generated

This tool generates a good chunk of the boilerplate necessary to hook up all your individual endpoints into your router.

Essentially what it gives you is an interface called _ServerInterface_ that you implement with your business logic. You build out the necessary functions and stick them into a struct which you can then pass to the oapi function that wraps that up in the standard _http.Handler_ interface used by the _net/http_ standard library.

Example:
```go
type ServerInterface interface {
	// Get a list of active rewards for your account
	// (GET /go/reward)
	GetGoReward(w http.ResponseWriter, r *http.Request)
	// Create a reward
	// (POST /go/reward)
	PostGoReward(w http.ResponseWriter, r *http.Request)
	// Delete a Reward by ID
	// (DELETE /go/reward/{id})
	DeleteGoRewardId(w http.ResponseWriter, r *http.Request, id int)
	// Get a Reward by ID
	// (GET /go/reward/{id})
	GetGoRewardId(w http.ResponseWriter, r *http.Request, id int)
}
```

Wrapping that custom implemented _ServerInterface_ in the cutomized _http.Handler_ and passing it to the http server just takes a handful of lines of code in the _main.go_ file. You probably rarely ever need to interact with that code unless you're adding new custom middleware.

By default the _ServerInterface_ you're given to implement has generic functions in it that take as args a basic request and response objects `w http.ResponseWriter, r *http.Request`, in addition to any URL parameters. You then need to parse any JSON yourself from the request and format the response JSON yourself.


### Strict server

Instead of handling the request/response JSON manually yourself, you can optionally generate what oapi-codegen calls a _StrictServerInterface_. This is the same as the _ServerInterface_ except that it strictly enforces Request and Response types with some auto-generated structs.

Example interface:
```go
type StrictServerInterface interface {
	// Get a list of active rewards for your account
	// (GET /go/reward)
	GetGoReward(ctx context.Context, request GetGoRewardRequestObject) (GetGoRewardResponseObject, error)
	// Create a reward
	// (POST /go/reward)
	PostGoReward(ctx context.Context, request PostGoRewardRequestObject) (PostGoRewardResponseObject, error)
	// Delete a Reward by ID
	// (DELETE /go/reward/{id})
	DeleteGoRewardId(ctx context.Context, request DeleteGoRewardIdRequestObject) (DeleteGoRewardIdResponseObject, error)
	// Get a Reward by ID
	// (GET /go/reward/{id})
	GetGoRewardId(ctx context.Context, request GetGoRewardIdRequestObject) (GetGoRewardIdResponseObject, error)
}
```
This way your code is much more strongly typed, and there's less boilerplate you need to write to parse JSON or create JSON. It should really be the default.

I'm not totally sure the best way to organize this, but currently in the project I have a _controllers_ package. _controllers_ contains files each of which implement their domain-specific part of the overall _StrictServerInterface_.

There's also a simple central file that defines this custom struct, _controllers/server.go_. This feels a bit weird and in general the directory and package structure of this project is up for iteration.


### Config

Some basic configuration for the tool can be set in a yaml file, for example _api/oapi-codegen.yml_:
```yml
# yaml-language-server: $schema=https://raw.githubusercontent.com/oapi-codegen/oapi-codegen/HEAD/configuration-schema.json
package: api
generate:
  std-http-server: true
  models: true
  strict-server: true
output: api/api.gen.go
```

In this file you can define the package you want to give the generated code, the path to where you want the code output.

You can also technically pass all of these via CLI flags to the command itself.

We're using the flag `std-http-server: true` which tells the tool to generate code for the standard _net/http_ server. You can also generate code for other popular http servers like _fiber_, _gin_, and others.

The _models: true_ flag tells the tool to generate code for the models described in the spec. You might want to disable this if you wanted more control over those. Though generating them here I think holds you closer to the spec which is kind of the point.

The final flag in the oapi-codegen config file is `strict-server: true`. This generates a second optional interface that strictly adhears to the defined schema.

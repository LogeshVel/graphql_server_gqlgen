# graphql_server_gqlgen
Golang GraphQL server using Schema first approach

Go package used for building the GraphQL server is [gqlgen](github.com/99designs/gqlgen)


Steps:

- create your project dir
- initialize go module (```go mod init <dirname>```)
- get gqlgen (```go get github.com/99designs/gqlgen```)
- initialize gqlgen config and generate the models.(```go run github.com/99designs/gqlgen init```)
- now run ```go mod tidy``` to get the pacakges

At this point, you should see various files and folders under the directory graphql. The directory graph were generated by gqlgen after you typed the ```init``` command. 

_model/model\_gen.go_: this is file with structs generated by gqlgen, defined by the schema file **schema.graphqls**

_generated/generated.go_: this is a file with generated code that injects context and middleware for each query and mutation.

**You should not modify either of those files since they will be modified by gqlgen as you update your schema.**

_schema.graphqls_: a GraphQL schema file where types, queries, and mutations are defined. The schema file uses schema-definition-language (SDL) to describe data types and operations (queries/mutations) in a human-readable way.

_schema.resolvers.go_: a go file with wrapper code for queries and mutations defined in **schema.graphqls**


now that we have initialized and got our models from gqlgen.

Proceed to define our own schema. 

- Delete the contents of the file graph/schema.graphqls

- Define your own schema in that file. After defining your schema now run the generate commands to generate the some boilerplate codes for our schema.

- Delete the example code in _schema.resolvers.go_

- ```go run github.com/99designs/gqlgen generate``` run this command. While running this command if you got some error like missing go.sum entry then get those packages to be added in go.sum. Mostly ```go get github.com/99designs/gqlgen``` this will resolve the error. then execute the generate command.

At this point, we have generated go files for our schemas.

- Proceed to define our resolvers in **schema.resolvers.go** file.

- once the resolver implementation is done.Server is ready to spin up ```go run server.go```

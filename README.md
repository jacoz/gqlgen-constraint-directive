## Package
This package allows you to define constraints to your input parameters directly in the graphql schema.

## Requirements
- [go](https://go.dev) >= 1.18
- [99designs/gqlgen](https://github.com/99designs/gqlgen) >= 0.17 
- [go-playground/validator](https://github.com/go-playground/validator) >= 10
- [go-playground/universal-translator](https://github.com/go-playground/universal-translator) >= 0.18

## Setup
1. Add this library to your go project with:
```shell
go get -u github.com/jacoz/gqlgen-constraint-directive
go mod tidy
```

2. Add to your `schema.graphqls` file the directive and then use it like the example below:
```graphql
directive @constraint(constraint: String!) on INPUT_FIELD_DEFINITION | ARGUMENT_DEFINITION

# ...

input SearchInput {
    offset: Int! @constraint(constraint: "min=0")
    limit: Int! @constraint(constraint: "max=50")
}
```

3. Regenerate the graphql schema with the command
```shell
go run github.com/99designs/gqlgen generate
```

4. Define your validator and your translator (optional).
In case you already are using a validator and a translator in your project, you can pass them as parameters, you can pass `nil` otherwise.
```go
constraintDirective := gqlgen_constraint_directive.New(nil, nil)

// or

constraintDirective := gqlgen_constraint_directive.New(yourValidator, yourTranslator)
```

5. Load the directive into your graphql server configuration: 
```go
c := generated.Config{}
c.Directives.Constraint = constraintDirective.Contraint
```

## Validators
You can find all the available validators in the [validator](https://github.com/go-playground/validator) package

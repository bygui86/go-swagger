
# go-swagger

Sample project to test Swagger-OpenAPI integration in Golang

---

## Prerequisites

- install `go-swagger` to run `go-swagger-rest-api`, instructions [here](https://goswagger.io/install.html)
- install `swaggo` to run `swaggo-rest-api`, instructions [here](https://github.com/swaggo/swag#getting-started)
- install `swagger2markup-cli` to convert Swagger JSON/YAML files to Markdown, instructions [here](https://github.com/Swagger2Markup/swagger2markup-cli#quick-usage)

---

## Applications

### go-swagger

```shell script
cd go-swagger-rest-api && make start
```

### swaggo

```shell script
cd swaggo-rest-api && make start
```

---

## Generate Swagger documentation

### go-swagger

`WARN`: it's easier if `swagger.json` is located in the same folder as the Swagger-UI. 

1. enter `go-swagger-rest-api`

    ```shell script
    cd go-swagger-rest-api
    ```

2. generate Swagger documentation

    ```shell script
    make generate-swagger
    ```

3. start application

    ```shell script
    make start
    ```

4. open Swagger UI

    ```shell script
    make open-swagger-ui
    ```

### swaggo

1. enter `swaggo-rest-api`

    ```shell script
    cd swaggo-rest-api
    ```

2. generate Swagger documentation

    ```shell script
    make generate-swagger
    ```

3. start application

    ```shell script
    make start
    ```

4. open Swagger UI

    ```shell script
    make open-swagger-ui
    ```

### swagger2markup

1. enter `swaggo-rest-api` or `go-swagger-rest-api`

    ```shell script
    cd swaggo-rest-api
    
    # or
    
    cd go-swagger-rest-api
    ```

2. generate Markdown documentation as file or folder

    ```shell script
    make generate-markdown-doc-file
    
    # or
    
    make generate-markdown-doc-folder
    ```

---

## Links

### officials
- https://github.com/go-openapi
- https://editor.swagger.io/

### libraries
#### go-swagger
- https://goswagger.io/
- https://github.com/go-swagger/go-swagger
#### swaggo
- https://github.com/swaggo/swag
- https://github.com/swaggo/http-swagger
- https://github.com/swaggo/http-swagger/issues/3

### tutorials
#### go-swagger
- https://medium.com/@supun.muthutantrige/lets-go-everything-you-need-to-know-about-creating-a-restful-api-in-go-part-iv-52666c5221d4
#### swaggo
- https://www.soberkoder.com/swagger-go-api-swaggo/

### ui
- https://github.com/swagger-api/swagger-ui

### swagger2markup
- https://github.com/Swagger2Markup
- https://github.com/Swagger2Markup/swagger2markup-cli
- http://swagger2markup.github.io/swagger2markup/1.3.3/#_command_line_interface

### cli
- https://www.npmjs.com/package/swagger-cli
- https://github.com/APIDevTools/swagger-cli

### tools
- https://github.com/OpenAPITools
- https://openapi-generator.tech/
- https://github.com/OpenAPITools/openapi-generator

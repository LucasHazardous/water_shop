# *Water Shop* - REST API in Go with Gin

![https://www.pexels.com/photo/clean-clear-cold-drink-416528/](./water.jpg)

---

## How to run

1. Install latest verion of golang.
2. Clone/Download this repository and go to the root of this project.
3. Download dependencies with:

```
go mod download
```

4. Make sure that port 8080 on your machine is free and start the program with:

```
go run .
```

or build it and run the .exe file:

```
go build .
```

## Content

This program provides five endpoints:

* **GET**: water bottle menu, water bottle by id
* **PATCH**: buy water (lower amount of bottles by 1), give water (increase amount of bottles by 1)
* **POST**: water request (one is already defined in a slice, you can send only one additional, example is provided in a json file)

## Testing

For sending requests to test the API I recommend using **cURL**, use below examples for reference:

- `curl localhost:8080/request --include --header "Content-Type: application/json" -d @water_request_example.json --request POST`
- `curl localhost:8080/buy?id=1 --request PATCH`
- `curl localhost:8080/menu`
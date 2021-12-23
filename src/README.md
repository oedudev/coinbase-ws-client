# VWAP Calculation

### Go version used on this project: 1.17

## Code Structure

### Adapter
The `adapter` is responsible for communication with coinbase websocket.

### Model
On `model` folder, we have all models of the application.

### Processor
On `processor` folder we have the core of processing for coinbase responses.
(For this module, we have an unit test and a benchmark test)

### Use Case
In this module, we join functions of the `adapter` and `processor` to get infos from coinbase and process the data.

In this example, I preferred to use function types instead of creating structs with methods because it's more easy to isolate responsabilities.  

## To Run
Run this app is very simple, just type: `go run main.go` 

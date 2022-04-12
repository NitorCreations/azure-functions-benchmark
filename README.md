# Azure Functions Benchmark

Performance benchmark tests for different Azure Functions runtimes.

Currently supported runtimes:

- Golang: Custom Go Runtime [azure-functions-go-handler](https://github.com/NitorCreations/azure-functions-go-handler)
- Javascript: Azure Node runtime

## Benchmark Tests

### EstimatePi

Compute the mathematical Pi value using Monte Carlo Simulation.

```
GET /api/estimatepi
```

Query Parameters
| Name | Description | Optional | Default |
| :--- | :---------- | :------- |:------- |
| `n` | number of iterations | `true` | 100_000 |

### Fibonacci

Compute Fibonacci numbers.

```
GET /api/fibonacci
```

### ImageCopyStorage

Read image data from blob storage, write image data to blob storage.

```
GET /api/imagecopystorage
```

### ImageResizeBind

Read image data from blob input binding, resize the image, write image data to blob output binding.

```
GET /api/imageresizebind
```

### ImageResizeStorage

Read image data from blob storage, resize the image, write image data to blob storage.

```
GET /api/imageresizestorage
```

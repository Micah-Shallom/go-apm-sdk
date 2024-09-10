Here's a comprehensive `README.md` for your APM SDK project:

---

# Telex APM SDK

The **Telex APM SDK** is an Application Performance Monitoring (APM) middleware built in Go, designed to collect and send application metrics and error logs to a specified webhook URL. The SDK is built to be integrated into various Go web applications, specifically supporting the Gin web framework.

## Features

- **Request Metrics Collection**: Track HTTP request metrics including endpoint, latency, status codes, and HTTP method.
- **Error Reporting**: Automatically capture panics and errors, report them to a webhook URL, and optionally rethrow panics based on configuration.
- **Performance Metrics Collection**: Monitor application performance metrics such as memory usage, CPU usage, garbage collection cycles, and goroutine count.
- **Customizable**: Supports configurable options such as timeouts and synchronous/asynchronous metric delivery.
- **Flexible Status Handling**: Categorize requests as `success` (for 2xx status codes) or `error` (for 3xx, 4xx, and 5xx status codes).

## Installation

To install the Telex APM SDK, run:

```bash
go get github.com/Micah-Shallom/go-apm-sdk
```

## Getting Started

### Step 1: Initialize the SDK

Before integrating the middleware, initialize the SDK in your application by calling the `Telex.Init()` function.

```go
package main

import (
    "github.com/Micah-Shallom/go-apm-sdk/telex"
)

func main() {
    apmOptions := telex.Options{
        WebhookURL:      "https://your-webhook-url",
        WaitForDelivery: false,
        Timeout:         5 * time.Second,
        Repanic:         true, // Control whether to rethrow panics
    }

    telex.Init(apmOptions)
}
```

### Step 2: Integrate with Gin

You can now integrate the Telex APM middleware with your Gin application by using the provided `handlegin` handler.

```go
package main

import (
    "github.com/Micah-Shallom/go-apm-sdk/telexgin"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    handler := &telexgin.Handler{
        Options: telexgin.Options{
            // Your handler options
        },
    }

    // Apply middleware to track request and error metrics
    router.Use(handler.handlegin)

    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    router.Run()
}
```

### Step 3: Sending Metrics

The SDK automatically collects request metrics, performance metrics, and error metrics. You can customize the webhook integration by configuring the following collectors:

- **Request Metrics Collector**: Captures endpoint path, method, latency, and status code.
- **Performance Metrics Collector**: Monitors CPU usage, memory usage, garbage collection cycles, and more.
- **Error Metrics Collector**: Tracks error messages, error counts, and details of the last error.

#### Example of Collected Metrics:

```json
{
    "EventName": "request_completed",
    "Message": "Endpoint: /ping | StatusCode: 200 | Latency: 150ms",
    "Status": "success",
    "Username": "APM",
    "PerformanceMetrics": {
        "CPUUsage": 4,
        "MemoryUsage": 2048000,
        "GCCycles": 12,
        "Goroutines": 10
    }
}
```

### Configuration Options

The `telex.Options` structure provides various configuration options for the SDK:

- **WebhookURL**: The URL to which metrics and errors are sent.
- **WaitForDelivery**: Boolean to determine if the SDK should wait for delivery of metrics before returning a response.
- **Timeout**: Time duration to wait before timing out while sending metrics.
- **Repanic**: Boolean to control whether to rethrow a panic after capturing it.
  
```go
type Options struct {
    WebhookURL      string        // Webhook URL for sending metrics
    WaitForDelivery bool          // Synchronous or asynchronous delivery
    Timeout         time.Duration // Request timeout for webhook
    Repanic         bool          // Rethrow panic after capturing it
}
```

## Usage

### Error Handling

The SDK captures errors and panics, sends them to the webhook, and optionally allows the application to recover from them. By default, if `Repanic` is set to `true`, the SDK will rethrow the panic after capturing it.

### Performance Metrics

The SDK tracks key performance metrics using Go's runtime stats such as memory allocation, goroutine count, and garbage collection cycles.

### Request Metrics

Each HTTP request is automatically tracked, including the following data:
- **Endpoint Path**: The API path being accessed.
- **Method**: HTTP method used (GET, POST, etc.).
- **Latency**: Time taken to process the request.
- **Status Code**: HTTP status code of the response.

### Error Metrics

Errors that occur during request handling, including panics, are captured and sent to the webhook.

## Example Webhook Payload

For a successful request:

```json
{
    "EventName": "request_completed",
    "Message": "Endpoint: /ping | StatusCode: 200 | Latency: 150ms",
    "Status": "success",
    "Username": "APM"
}
```

For an error:

```json
{
    "EventName": "error_occurred",
    "Message": "Error: runtime error: index out of range",
    "Status": "error",
    "Username": "APM"
}
```


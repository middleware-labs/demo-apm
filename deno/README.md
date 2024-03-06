## Tracing in Deno

Tracing in Deno allows you to instrument your code to monitor and log network requests and other events. This guide outlines the steps to set up tracing using middleware in Deno.

### Prerequisites

Before you begin, ensure you have Deno installed on your system. You can install Deno by following the instructions on the [official Deno website](https://deno.land/#installation).

### Steps

1. **Import Dependencies**: Start by importing the necessary packages into your Deno project. You'll need the `serve` function from the Deno standard library and middleware functions from the `middlewareio` module for instrumentation.

    ```typescript
    import { serve } from "https://deno.land/std@0.120.0/http/server.ts";
    import { track, httpTracer, info, warn, error, debug } from "https://deno.land/x/middlewareio@v1.0.7/mod.ts";
    ```

2. **Configure Tracing**: Initialize tracing by calling the `track` function with the appropriate configuration options. This sets up the middleware to trace network requests.

    ```typescript
    track({
    serviceName: '{SERVICE_NAME}',
    target: 'https://{ACCOUNT-UID}.middleware.io:443',
    accessToken: '{ACCOUNT_KEY}',
    })
    ```

3. **Define Request Handler**: Create a request handler function that defines how your server responds to incoming HTTP requests.

    ```typescript
    function handler(_req: Request): Response {
       const data = {
           message: `Hello world!`,
       };
       info('Request received in test_function_2!');
       return new Response(JSON.stringify(data), { headers: { 'Content-Type': 'application/json' } });
    }
    ```

4. **Start Server with Tracing Middleware**: Use the `serve` function to start an HTTP server, passing the `httpTracer` middleware along with your request handler.

    ```typescript
    await serve(httpTracer(handler));
    ```

5. **Optional Logging**: You can also log additional information using the `info`, `warn`, `error`, and `debug` functions provided by the middlewareio module.

    ```typescript
    info("info");
    warn("warn");
    error("error");
    debug("debug");
    ```

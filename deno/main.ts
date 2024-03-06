import { serve } from "https://deno.land/std@0.120.0/http/server.ts";
import {track,httpTracer,info,warn,error,debug} from "https://deno.land/x/middlewareio@v1.0.7/mod.ts";

track({
    serviceName: '{SERVICE_NAME}',
    target: 'https://{ACCOUNT-UID}.middleware.io:443',
    accessToken: '{ACCOUNT_KEY}',
})


info("info message")
warn("warn message")
error("error message")
debug("debug message")

function handler(_req: Request): Response {
    const data = {
        message: `Hello world!`,
    }
    return new Response(JSON.stringify(data), { headers: { 'Content-Type': 'application/json' } })
}
await serve(httpTracer(handler))

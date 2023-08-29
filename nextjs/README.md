# Next.js APM Setup

[![npm](https://img.shields.io/npm/v/%40middleware.io%2Fagent-apm-nextjs)](https://www.npmjs.com/package/@middleware.io/agent-apm-nextjs)

|  Traces  |  Metrics  |  Profiling  |  Logs (App/Custom)  |
|:--------:|:---------:|:-----------:|:-------------------:|
|   Yes    |    No     |     Yes     |       No/Yes        |

## Prerequisites

Make sure you have installed the latest version of Next.js or a version greater than 13.4+, as Vercel introduced their experimental feature in that release.

Before proceeding with the Next.js APM setup, make sure you have the `@opentelemetry/api` package installed. If it's not already installed, run the following command:

```
npm install @opentelemetry/api@">=1.3.0 <1.5.0"
```

## Guide

### Step 1: Install Next.js APM package

Run the command below in your terminal to install Middleware's Next.js APM package:
```
npm install @middleware.io/agent-apm-nextjs
```

### Step 2: Modify the `next.config.js` file

As this feature is experimental, you need to explicitly opt-in by adding the following code to your **next.config.js** file:
```
const nextConfig = {
     ---
     ---
     experimental: {
         instrumentationHook: true
     }
     ---
     ---
}
module.exports = nextConfig
```

### Step 3: Create an `Instrumentation` file

Create a custom `instrumentation.ts` file in your project root directory, and add the following code as per your choice:

- If you are using [Vercel](https://vercel.com/) platform to deploy your projects, then use the code snippet below for serverless functions:
```
// @ts-ignore
import tracker from '@middleware.io/agent-apm-nextjs';

export function register() {
    tracker.track({
        projectName: "<PROJECT-NAME>",
        serviceName: "<SERVICE-NAME>",
        accessToken: "<ACCESS-TOKEN>",
        target: "vercel",
    });
}
```
**Note: After Deploying your project on Vercel, you need to integrate the [Middleware](https://vercel.com/integrations/middleware) from the marketplace. You can find more details [here](https://docs.middleware.io/docs/apm-configuration/next-js/vercel-integration). To get a better idea, you can clone the sample project from the [GitHub](https://github.com/middleware-labs/demo-apm/tree/master/nextjs/setup) repository.**
- If you are using [Middleware's Host-agent](https://docs.middleware.io/docs/installation) on your machine then use code snippet below:
```
// @ts-ignore
import tracker from '@middleware.io/agent-apm-nextjs';

export function register() {
    tracker.track({
        projectName: "<PROJECT-NAME>",
        serviceName: "<SERVICE-NAME>",
        accessToken: "<ACCESS-TOKEN>",
    });
}
```
- If you want to instrument your project without installing any host, then use the code snippet below:
```
// @ts-ignore
import tracker from '@middleware.io/agent-apm-nextjs';

export function register() {
    tracker.track({
        projectName: "<PROJECT-NAME>",
        serviceName: "<SERVICE-NAME>",
        accessToken: "<ACCESS-TOKEN>",
        target: "https://{ACCOUNT-UID}.middleware.io"
    });
}
```
### Step 4: Enable Logging
To enable logging in your project, add the following code in your file:
```javascript
// @ts-ignore
import tracker from '@middleware.io/agent-apm-nextjs';

export default async function handler(req, res) {
    // ...
    // Your existing code

    tracker.info("Info Sample");
    tracker.warn("Warn Sample", {
        "tester": "Alex",
    });
    tracker.debug("Debug Sample");
    tracker.error("Error Sample");

    // ...
    // Your existing code
}
```
---------------------

## Note for APM inside Kubernetes

If you are using APM in a Kubernetes cluster make sure to follow these 2 steps:

### Step 1 : Find your Middleware Service namespace
For older setup, your "mw-service" can be inside "mw-agent-ns-{FIRST-5-LETTERS-OF-API-KEY}" namespace

For newer setup, we simplified the namespace name to "mw-agent-ns"

### Step 2 : Set this ENV variable in your application deployment YAML
```
MW_AGENT_SERVICE=mw-service.NAMESPACE.svc.cluster.local
```
Please replace "NAMESPACE" with the correct value that you found from Step 1.

# Next.js APM Setup

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
import { track } from '@middleware.io/agent-apm-nextjs';

export function register() {
    track({
        projectName: "<PROJECT-NAME>",
        serviceName: "<SERVICE-NAME>",
        target: "vercel",
    });
}
```
**Note: After Deploying your project on Vercel, you need to install the [Middleware](https://vercel.com/integrations/middleware) app from the marketplace. You can find more details [here](https://docs.middleware.io/docs/vercel).**
- If you are using [Middleware's Host-agent](https://docs.middleware.io/docs/installation) on your machine then use code snippet below:
```
// @ts-ignore
import { track } from '@middleware.io/agent-apm-nextjs';

export function register() {
    track({
        projectName: "<PROJECT-NAME>",
        serviceName: "<SERVICE-NAME>",
    });
}
```
- If you want to instrument your project without installing any host, then use the code snippet below:
```
// @ts-ignore
import { track } from '@middleware.io/agent-apm-nextjs';

export function register() {
    track({
        projectName: "<PROJECT-NAME>",
        serviceName: "<SERVICE-NAME>",
        accountKey: "{ACCOUNT_KEY}",
        target: "https://{ACCOUNT-UID}.middleware.io"
    });
}
```
---------------------

## Note:

If you are using APM in a Kubernetes cluster, make sure to pass this ENV variable:

```
MW_AGENT_SERVICE=mw-service.mw-agent-ns-{FIRST-5-LETTERS-OF-API-KEY}.svc.cluster.local
```
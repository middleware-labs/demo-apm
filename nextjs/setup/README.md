# Next.js Demo Project & APM Setup

## Guide

***Note: This Demo project is specially denotes the integration between [Middleware](https://middleware.io/) and [Vercel](https://vercel.com/). This demo can help you visualize your traces of Vercel-deployed projects on the Middleware platform.***


### Step 1: Set up a sample project

Make sure you have installed the latest version of Next.js or a version greater than 13.4+, as Vercel introduced their experimental feature in that release.

Run the command below in your terminal to install Next.js sample project, i.e. `my-app`:
```bash
npx create-next-app@latest --example api-routes
```

### Step 2: Install Next.js APM package

Now go to `my-app` directory and run the command below in your terminal to install Middleware's Next.js APM package:
```bash
npm i @opentelemetry/api@">=1.3.0 <1.5.0"
```

```bash
npm install @middleware.io/agent-apm-nextjs
```

### Step 3: Modify the `next.config.js` file

As this feature is experimental, you need to explicitly opt-in by adding the following code to your **next.config.js** file:
```javascript
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

### Step 4: Create an `Instrumentation` file

Create a custom `instrumentation.ts` file in your project root directory, and add the following code:
```javascript
// @ts-ignore
import tracker from '@middleware.io/agent-apm-nextjs';

export function register() {
    tracker.track({
        projectName: "<PROJECT-NAME>",
        serviceName: "<SERVICE-NAME>",
        accountKey: "<ACCOUNT-KEY>",
        target: "vercel",
    });
}
```
*Note: You can find your &lt;ACCOUNT-KEY&gt; on the Installation screen for [NextJs / Vercel](https://app.middleware.io/installation#apm/nextjs).*

### Step 5: Enable Logging
To enable logging in your project, add the following code in your file where you want to log, like: `/pages/api/people/[id].ts`:
```javascript
// @ts-ignore
import tracker from '@middleware.io/agent-apm-nextjs';

export default function personHandler(...) {
    // ...
    // Your existing code

    if (person) {

        tracker.info(`Person with id ${id} found successfully.`, person);
        return res.status(200).json(person)

    } else {

        racker.error(`Requested person with id ${id} not found.`, { id: id });
        return res.status(404).json({ message: `User not found.` })

    }
}
```

### Step 6: Integrate Middleware on Vercel 
 Now you can deploy your project on Vercel. Afterward, you need to integrate the [Middleware](https://vercel.com/integrations/middleware) from the marketplace. You can find more details [here](https://docs.middleware.io/docs/apm-configuration/next-js/vercel-integration).

***Unlock the power of seamless integration between Middleware and Vercel, and elevate your Next.js projects with advanced APM capabilities.*** 🚀
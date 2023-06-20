// @ts-ignore
import tracker from '@middleware.io/agent-apm-nextjs';

export function register() {
    tracker.track({
        projectName: "sock-shop",
        serviceName: "socket-service",
        accountKey: "abcdefghijklmnopqrstuvwxyz0123456789",
        target: "vercel",
    });

    // Please replace "accountKey" with your own Account key / Access-token key.

    tracker.info("Build completed and triggered instrumentation.ts file.");
}

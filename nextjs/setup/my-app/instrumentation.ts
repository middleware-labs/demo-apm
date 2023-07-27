// @ts-ignore
import tracker from '@middleware.io/agent-apm-nextjs';

export function register() {
    tracker.track({
        projectName: "sock-shop",
        serviceName: "socket-service",
        accessToken: "abcdefghijklmnopqrstuvwxyz0123456789",
        target: "vercel",
    });

    // Please replace "accessToken" with your own Access-token key.

    tracker.info("Build completed and triggered instrumentation.ts file.");
}

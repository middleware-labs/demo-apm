import  {init,track} from "@middleware.io/agent-apm-worker";

export interface Env {
	OTLP_ENDPOINT: string;
}

export default {
	async fetch(
		request: Request,
		env: Env,
		ctx: ExecutionContext
	): Promise<Response> {

		init({
			projectName:"demo-cloudflare-project",
			serviceName:"demo-cloudflare-service",
			accountKey:"{ACCOUNT_KEY}",
            target:"https://{ACCOUNT-UID}.middleware.io"
		})
		
		const sdk = track(request, ctx);
		
		sdk.logger.error("demo error log")
		sdk.logger.info("demo info log")
		sdk.logger.warn("demo warn log")
		sdk.logger.debug("demo debug log")

		const url = new URL(request.url);

		const response = await sdk.fetch(`https://httpbin.org${url.pathname}`);
		return sdk.sendResponse(response);
	},
};

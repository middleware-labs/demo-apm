variable "region" {
  type    = string
  default = "eu-north-1"
}

variable "prefix" {
  description = "Name prefix for every resource."
  type        = string
  default     = "tracing-demo"
}

variable "allowed_cidr" {
  description = "CIDR allowed to reach order-service on :8080. Defaults to your current public IP."
  type        = string
  default     = null
}

variable "traffic_interval_ms" {
  description = "How often order-service calls itself to generate traces."
  type        = number
  default     = 3000
}

# ── Middleware ────────────────────────────────────────────────
variable "enable_middleware" {
  description = "Add the Middleware agent sidecar, OTel auto-instrumentation init container and FireLens log routing."
  type        = bool
  default     = true
}

variable "mw_api_key" {
  description = "Middleware API key (required when enable_middleware = true)."
  type        = string
  default     = ""
  sensitive   = true
}

variable "mw_target" {
  description = "Middleware target URL, e.g. https://<uid>.middleware.io:443"
  type        = string
  default     = "https://sandbox.middleware.io:443"
}

variable "mw_agent_image" {
  type    = string
  default = "ghcr.io/middleware-labs/mw-host-agent:1.20.1"
}

variable "python_autoinstrumentation_image" {
  description = "OpsAI build: captures the function body when an exception is recorded."
  type        = string
  default     = "ghcr.io/middleware-labs/opentelemetry-operator/autoinstrumentation-python:0.64b0-opsai"
}

variable "java_autoinstrumentation_image" {
  type    = string
  default = "ghcr.io/open-telemetry/opentelemetry-operator/autoinstrumentation-java:2.19.0"
}

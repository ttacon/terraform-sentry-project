variable "sentry_api_token" {
  description = "An API token to manage the Sentry resources"
  type        = string
}

variable "sentry_base_url" {
  description = "The base URL of the Sentry deployment to interact with. This is primarily useful if you are running on-prem Sentry. Otherwise, the default value will do."
  type        = string
  default     = "https://app.getsentry.com/api/"
}

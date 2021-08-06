variable "name" {
  description = "Sentry project name"
  type        = string
}

variable "team" {
  description = "Team to assign to the Sentry project"
  type        = string
}

variable "organization" {
  description = "The Sentry organization slug"
  type        = string
}

variable "slug" {
  description = "The project slug"
  type        = string
}

variable "platform" {
  description = "The platform that will report errors to Sentry"
  type        = string
}

variable "resolve_age" {
  description = "The number of hours before issues should auto-resolve"
  type        = number
  default     = 720
}

variable "has_slack_notification_rule" {
  description = "Whether or not to include a Slack notification rule"
  type        = bool
  default     = true
}

variable "event_frequency_threshold" {
  description = "The threshold at which to trigger the Slack notification"
  type        = number
  default     = 100
}

variable "event_interval" {
  description = "The interval within which to consider the event frequency threshold"
  type        = string
  default     = "1h"
}

variable "action_slack_channel_name" {
  description = "The name of the Slack channel to post events to"
  type        = string
  default     = "#sentry"
}

variable "action_slack_channel_id" {
  description = "The ID of the Slack channel to post events to"
  type        = string
  default     = ""
}

variable "action_slack_workspace_id" {
  description = "The ID of the Slack workspace to post events to"
  type        = string
  default     = ""
}

variable "action_slack_alert_tags" {
  description = "Event tags to display in slack notifications"
  type        = list(string)
  default     = []
}

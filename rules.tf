resource "sentry_rule" "slack_notification" {
  count        = var.has_slack_notification_rule ? 1 : 0
  name         = "Slack notifications"
  organization = var.organization
  project      = var.slug
  action_match = "any"
  frequency    = 60

  conditions = [
    {
      id   = "sentry.rules.conditions.first_seen_event.FirstSeenEventCondition"
      name = "A new issue is created"
    },
    {
      id       = "sentry.rules.conditions.event_frequency.EventFrequencyCondition"
      name     = "The issue is seen more than 100 times in 1h"
      interval = var.event_interval
      value    = var.event_frequency_threshold
    }
  ]

  filters = []

  actions = [
    {
      channel    = var.action_slack_channel_name
      channel_id = var.action_slack_channel_id
      id         = "sentry.integrations.slack.notify_action.SlackNotifyServiceAction"
      name       = "Send a notification to the Mixmax Slack workspace to #sentry and show tags [] in notification"
      tags       = ""
      workspace  = var.action_slack_workspace_id
    }
  ]
}

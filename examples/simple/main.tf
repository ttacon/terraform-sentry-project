module "sentry_project" {
  ## In a real service, use the following line instead of the relative source path:
  #source = "git::ssh://git@github.com/ttacon/terraform-sentry-project.git?ref=vX.X.X"
  source = "../.."

  name         = "Sample project"
  team         = "sample-team"
  organization = "sample-org"
  slug         = "sample-project"
  platform     = "node"

  # Extra configuration for the Slack event rule
  action_slack_channel_id   = "C4934F0F3"
  action_slack_workspace_id = "53846"
}

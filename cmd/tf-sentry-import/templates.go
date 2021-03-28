package main

const projectTemplate = `
module "sentry_project_{{.ProjectSlug}}" {
  source = "git::ssh://git@github.com/ttacon/terraform-sentry-project.git?ref={{.Version}}"

  name         = "{{.ProjectName}}"
  team         = "{{.ProjectTeam}}"
  organization = "{{.OrgSlug}}"
  slug         = "{{.ProjectSlug}}"
  {{if .Platform}}platform     = "{{.Platform}}"{{end}}
  {{if .IncludeSlackRule}}
  # Extra configuration for the Slack event rule
  action_slack_channel_name = "{{.SlackChannelName}}"
  action_slack_channel_id   = "{{.SlackChannelID}}"
  action_slack_workspace_id = "{{.SlackWorkspaceID}}"
  {{end}}
}
`

type TerraformProjectInfo struct {
	Version          string
	ProjectName      string
	ProjectTeam      string
	OrgSlug          string
	ProjectSlug      string
	Platform         string
	IncludeSlackRule bool
	SlackChannelName string
	SlackChannelID   string
	SlackWorkspaceID string
}

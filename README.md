# terraform-sentry-project
An opinionated take on how to structure Sentry projects in terraform
## Requirements

The following requirements are needed by this module:

- terraform (>= 0.13)

- sentry (~> 0.6.0)

## Providers

The following providers are used by this module:

- sentry (~> 0.6.0)

## Required Inputs

The following input variables are required:

### name

Description: Sentry project name

Type: `string`

### organization

Description: The Sentry organization slug

Type: `string`

### platform

Description: The platform that will report errors to Sentry

Type: `string`

### slug

Description: The project slug

Type: `string`

### team

Description: Team to assign to the Sentry project

Type: `string`

## Optional Inputs

The following input variables are optional (have default values):

### action\_slack\_channel\_id

Description: The ID of the Slack channel to post events to

Type: `string`

Default: `""`

### action\_slack\_channel\_name

Description: The name of the Slack channel to post events to

Type: `string`

Default: `"#sentry"`

### action\_slack\_workspace\_id

Description: The ID of the Slack workspace to post events to

Type: `string`

Default: `""`

### event\_frequency\_threshold

Description: The threshold at which to trigger the Slack notification

Type: `number`

Default: `100`

### event\_interval

Description: The interval within which to consider the event frequency threshold

Type: `string`

Default: `"1h"`

### has\_slack\_notification\_rule

Description: Whether or not to include a Slack notification rule

Type: `bool`

Default: `true`

### resolve\_age

Description: The number of hours before issues should auto-resolve

Type: `number`

Default: `720`

## Outputs

No output.


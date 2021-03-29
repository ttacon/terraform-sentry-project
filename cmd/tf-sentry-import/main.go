package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/url"
	"path/filepath"

	"github.com/jianyuan/go-sentry/sentry"
)

// This code is a bunch of hacks to make using Sentry terraform simpler
// when you have an existant Sentry deployment.
//
// Use at your own caution!!!
//
// Notes for later development:
//  - Should support dry run of identifying what resources we would generate
//    tf code for.
//  - We should be able to say which scopes our API token will need.
//  - Support generation of API token and retrieval of it.
//  - Support flag for calling "terraform import"
var (
	dryRun = flag.Bool("dry-run", false, "run in dry run mode")

	sentryOrg = flag.String("org", "", "The organization in Sentry to generate code for")

	generateTeams = flag.Bool("gen-teams", true, "whether or not to generate team tf code")
	projects      = flag.String("projects", "", "specific projects to generate terraform code for")

	determineScopes = flag.Bool("determine-scopes", false, "Identify which scopes are needed to run terraform code")

	importTFState = flag.Bool("import-state", false, "Import terraform state")

	apiToken = flag.String("api-token", "", "Sentry API token")

	generateAPIToken = flag.Bool("gen-api-token", false, "Generate an API token from Sentry")

	outputDirectory = flag.String("out-dir", "", "Output directory")

	baseURL = flag.String("base-url", sentry.DefaultBaseURL, "The base URL of the Sentry deployment to talk to")

	sentryModuleVersion = flag.String("module-version", "v1.0.0", "The module version to use")

	// Slack flags if we can't determine rules on our own.
	includeSlackRule = flag.Bool("include-slack", false, "Whether or not to include a CLI specified Slack rule")
	slackChannelName = flag.String("slack-channel-name", "#sentry", "Slack channel name")
	slackChannelID   = flag.String("slack-channel-id", "", "Slack channel ID")
	slackWorkspaceID = flag.String("slack-workspace-id", "", "Slack workspace ID")
)

func main() {
	flag.Parse()

	if len(*apiToken) == 0 {
		fmt.Println("must provide -api-token")
		return
	} else if len(*sentryOrg) == 0 {
		fmt.Println("must provide -org")
		return
	}

	parsedURL, err := url.Parse(*baseURL)
	if err != nil {
		fmt.Println("failed to parse URL, err: ", err)
		return
	}
	client := sentry.NewClient(nil, parsedURL, *apiToken)

	// Get projects
	//
	// We should specify which ones we want to import.
	projects, resp, err := client.Projects.List()
	if err != nil {
		fmt.Println("failed to list Sentry projects: ", err)
		return
	}

	_ = resp.Body.Close()

	// Get rules
	//
	// This should rely on the section above.
	var projectRuleMap = make(map[string][]sentry.Rule)
	for i, proj := range projects {
		fmt.Printf("[%d/%d] retrieving rules for %q\n", i, len(projects), proj.Slug)
		rules, resp, err := client.Rules.List(
			*sentryOrg,
			proj.Slug,
		)
		if err != nil {
			fmt.Printf("failed to retrieve rules for project %q, err: %s\n", proj.Slug, err)
		}
		_ = resp.Body.Close()

		projectRuleMap[proj.Slug] = rules
	}

	// Get teams
	//
	// Can opt in or out of this resource type.
	var teams []sentry.Team
	if *generateTeams {
		teams, resp, err = client.Teams.List(*sentryOrg)
		if err != nil {
			fmt.Println("failed to retrieve teams, err: ", err)
		}
	}
	_ = teams // Silence unused messages for now

	// Generate tf code.
	var templ = template.Must(
		template.New("projectTemplate").
			Parse(projectTemplate),
	)
	projectBuf := bytes.NewBuffer(nil)
	for _, proj := range projects {
		data := TerraformProjectInfo{
			Version:     *sentryModuleVersion,
			ProjectName: proj.Name,
			ProjectTeam: proj.Team.Slug,
			ProjectSlug: proj.Slug,
			OrgSlug:     *sentryOrg,
			Platform:    proj.Platform,
		}

		if rules, ok := projectRuleMap[proj.Slug]; ok && len(rules) > 0 {
			// TODO(ttacon): convert rules and add them in for
			// Slack rule.

			// Hit a snag with Sentry 9 to Sentry 10 migration.
			//
			// Will play with this more when get to Sentry 10 upgrades.
		} else if *includeSlackRule {
			data.IncludeSlackRule = true
			data.SlackChannelName = *slackChannelName
			data.SlackChannelID = *slackChannelID
			data.SlackWorkspaceID = *slackWorkspaceID
		}

		if err := templ.Execute(projectBuf, data); err != nil {
			fmt.Println("failed to execute template, err: ", err)
			return
		}
	}

	if err := ioutil.WriteFile(
		filepath.Join(
			*outputDirectory,
			"projects.tf",
		),
		projectBuf.Bytes(),
		0755,
	); err != nil {
		fmt.Println("failed to write output file, err: ", err)
		return
	}

	// Call `terraform import.`

}

package all

import (
	// Register notifiers.
	_ "github.com/stackrox/rox/central/notifiers/awssh"
	_ "github.com/stackrox/rox/central/notifiers/cscc"
	_ "github.com/stackrox/rox/central/notifiers/email"
	_ "github.com/stackrox/rox/central/notifiers/jira"
	_ "github.com/stackrox/rox/central/notifiers/pagerduty"
	_ "github.com/stackrox/rox/central/notifiers/slack"
	_ "github.com/stackrox/rox/central/notifiers/splunk"
	_ "github.com/stackrox/rox/central/notifiers/sumologic"
	_ "github.com/stackrox/rox/central/notifiers/teams"
	_ "github.com/stackrox/rox/pkg/notifiers/generic"
	_ "github.com/stackrox/rox/pkg/notifiers/syslog"
)

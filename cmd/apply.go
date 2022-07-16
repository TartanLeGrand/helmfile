package cmd

import (
	"github.com/helmfile/helmfile/pkg/app"
	"github.com/helmfile/helmfile/pkg/config"
	"github.com/urfave/cli"
)

func addApplySubcommand(cliApp *cli.App) {
	cliApp.Commands = append(cliApp.Commands, cli.Command{
		Name:  "apply",
		Usage: "apply all resources from state file only when there are changes",
		Flags: []cli.Flag{
			cli.StringSliceFlag{
				Name:  "set",
				Usage: "additional values to be merged into the command",
			},
			cli.StringSliceFlag{
				Name:  "values",
				Usage: "additional value files to be merged into the command",
			},
			cli.IntFlag{
				Name:  "concurrency",
				Value: 0,
				Usage: "maximum number of concurrent helm processes to run, 0 is unlimited",
			},
			cli.BoolFlag{
				Name:  "validate",
				Usage: "validate your manifests against the Kubernetes cluster you are currently pointing at. Note that this requires access to a Kubernetes cluster to obtain information necessary for validating, like the list of available API versions",
			},
			cli.IntFlag{
				Name:  "context",
				Value: 0,
				Usage: "output NUM lines of context around changes",
			},
			cli.StringFlag{
				Name:  "output",
				Value: "",
				Usage: "output format for diff plugin",
			},
			cli.BoolFlag{
				Name:  "detailed-exitcode",
				Usage: "return a non-zero exit code 2 instead of 0 when there were changes detected AND the changes are synced successfully",
			},
			cli.StringFlag{
				Name:  "args",
				Value: "",
				Usage: "pass args to helm exec",
			},
			cli.BoolFlag{
				Name:  "retain-values-files",
				Usage: "DEPRECATED: Use skip-cleanup instead",
			},
			cli.BoolFlag{
				Name:  "skip-cleanup",
				Usage: "Stop cleaning up temporary values generated by helmfile and helm-secrets. Useful for debugging. Don't use in production for security",
			},
			cli.BoolFlag{
				Name:  "skip-crds",
				Usage: "if set, no CRDs will be installed on sync. By default, CRDs are installed if not already present",
			},
			cli.BoolTFlag{
				Name:  "skip-needs",
				Usage: `do not automatically include releases from the target release's "needs" when --selector/-l flag is provided. Does nothing when when --selector/-l flag is not provided. Defaults to true when --include-needs or --include-transitive-needs is not provided`,
			},
			cli.BoolFlag{
				Name:  "include-needs",
				Usage: `automatically include releases from the target release's "needs" when --selector/-l flag is provided. Does nothing when when --selector/-l flag is not provided`,
			},
			cli.BoolFlag{
				Name:  "include-transitive-needs",
				Usage: `like --include-needs, but also includes transitive needs (needs of needs). Does nothing when when --selector/-l flag is not provided. Overrides exclusions of other selectors and conditions.`,
			},
			cli.BoolFlag{
				Name:  "skip-diff-on-install",
				Usage: "Skips running helm-diff on releases being newly installed on this apply. Useful when the release manifests are too huge to be reviewed, or it's too time-consuming to diff at all",
			},
			cli.BoolFlag{
				Name:  "include-tests",
				Usage: "enable the diffing of the helm test hooks",
			},
			cli.StringSliceFlag{
				Name:  "suppress",
				Usage: "suppress specified Kubernetes objects in the diff output. Can be provided multiple times. For example: --suppress KeycloakClient --suppress VaultSecret",
			},
			cli.BoolFlag{
				Name:  "suppress-secrets",
				Usage: "suppress secrets in the diff output. highly recommended to specify on CI/CD use-cases",
			},
			cli.BoolFlag{
				Name:  "show-secrets",
				Usage: "do not redact secret values in the diff output. should be used for debug purpose only",
			},
			cli.BoolFlag{
				Name:  "suppress-diff",
				Usage: "suppress diff in the output. Usable in new installs",
			},
			cli.BoolFlag{
				Name:  "skip-deps",
				Usage: `skip running "helm repo update" and "helm dependency build"`,
			},
			cli.BoolFlag{
				Name:  "wait",
				Usage: `Override helmDefaults.wait setting "helm upgrade --install --wait"`,
			},
			cli.BoolFlag{
				Name:  "wait-for-jobs",
				Usage: `Override helmDefaults.waitForJobs setting "helm upgrade --install --wait-for-jobs"`,
			},
		},
		Action: Action(func(a *app.App, c config.ConfigImpl) error {
			return a.Apply(c)
		}),
	})
}
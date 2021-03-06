package cmd

import (
	"io"
	"strings"

	"github.com/jenkins-x/jx/pkg/gits"
	"github.com/jenkins-x/jx/pkg/jx/cmd/templates"
	cmdutil "github.com/jenkins-x/jx/pkg/jx/cmd/util"
	"github.com/jenkins-x/jx/pkg/kube"
	"github.com/spf13/cobra"
)

// GetQuickstartLocationOptions containers the CLI options
type GetQuickstartLocationOptions struct {
	GetOptions
}

const (
	quickstartLocation  = "quickstartlocation"
	quickstartLocations = quickstartLocation + "s"
)

var (
	quickstartLocationsAliases = []string{
		quickstartLocation, "quickstartloc", "qsloc",
	}

	getQuickstartLocationLong = templates.LongDesc(`
		Display one or many Quickstart Locations for the current Team.

		For more documentation see: [https://jenkins-x.io/developing/create-quickstart/#customising-your-teams-quickstarts](https://jenkins-x.io/developing/create-quickstart/#customising-your-teams-quickstarts)

`)

	getQuickstartLocationExample = templates.Examples(`
		# List all the quickstart locations
		jx get quickstartlocations

		# List all the quickstart locations via an alias
		jx get qsloc

	`)
)

// NewCmdGetQuickstartLocation creates the new command for: jx get env
func NewCmdGetQuickstartLocation(f cmdutil.Factory, out io.Writer, errOut io.Writer) *cobra.Command {
	options := &GetQuickstartLocationOptions{
		GetOptions: GetOptions{
			CommonOptions: CommonOptions{
				Factory: f,
				Out:     out,
				Err:     errOut,
			},
		},
	}
	cmd := &cobra.Command{
		Use:     quickstartLocations,
		Short:   "Display one or many Quickstart Locations",
		Aliases: quickstartLocationsAliases,
		Long:    getQuickstartLocationLong,
		Example: getQuickstartLocationExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			cmdutil.CheckErr(err)
		},
	}

	options.addGetFlags(cmd)
	return cmd
}

// Run implements this command
func (o *GetQuickstartLocationOptions) Run() error {
	jxClient, ns, err := o.JXClientAndDevNamespace()
	if err != nil {
		return err
	}
	err = o.registerEnvironmentCRD()
	if err != nil {
		return err
	}

	locations, err := kube.GetQuickstartLocations(jxClient, ns)
	if err != nil {
		return err
	}

	table := o.CreateTable()
	table.AddRow("GIT SERVER", "KIND", "OWNER", "INCLUDES", "EXCLUDES")

	for _, location := range locations {
		kind := location.GitKind
		if kind == "" {
			kind = gits.KindGitHub
		}
		table.AddRow(location.GitURL, kind, location.Owner, strings.Join(location.Includes, ", "), strings.Join(location.Excludes, ", "))
	}
	table.Render()
	return nil
}

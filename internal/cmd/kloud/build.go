package kloud

import (
	"github.com/kloud-team/kloud/internal/pkg/kloud"
	"github.com/spf13/cobra"
)

var (
	build    *kloud.Build
	buildCmd = &cobra.Command{
		Use:     "build",
		Aliases: []string{"b"},
		Short:   "Build your project and create a docker image",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if err := build.Validate(); err != nil {
				kloud.Exit(err)
			}

			if err := build.Execute(); err != nil {
				kloud.Exit(err)
			}
		},
	}
)

func init() {
	var err error
	if build, err = kloud.NewBuild(); err != nil {
		kloud.Exit(err)
	}

	buildCmd.Flags().StringVarP(&build.Dockerfile, "dockerfile", "d", build.Dockerfile, "path of Dockerfile")
	buildCmd.Flags().StringVarP(&build.WorkDirectoryPath, "work_directory", "w", build.WorkDirectoryPath, "path of project")
	rootCmd.AddCommand(buildCmd)
}

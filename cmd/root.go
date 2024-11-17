package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Config is the configuration for the application.
type Config struct {
	Version int `yaml:"version"`
	Dbs     []struct {
		SchemasPath   string `yaml:"schemas_path"`
		QueriesPath   string `yaml:"queries_path"`
		SeedsPath     string `yaml:"seeds_path"`
		OutputSchema  string `yaml:"output_schema"`
		OutputQueries string `yaml:"output_queries"`
		OutputSeeds   string `yaml:"output_seeds"`
	} `yaml:"dbs"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sqlcaush",
	Short: "A cli tools to squash sql files into single files based on type.",
	Long: `A cli tools to squash sql files into single files based on type configured in config file.

Example:

sqlcaush combine -c config.yaml
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(combineCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

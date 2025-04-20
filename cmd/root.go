package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Config is the configuration for the application.
type Config struct {
	Version int        `yaml:"version"`
	Dbs     []DbConfig `yaml:"dbs"`
}

// DbConfig contains the configuration for a single database.
type DbConfig struct {
	SchemasPath   string `yaml:"schemas"`
	QueriesPath   string `yaml:"queries"`
	SeedsPath     string `yaml:"seeds"`
	OutputSchema  string `yaml:"schema"`
	OutputQueries string `yaml:"query"`
	OutputSeeds   string `yaml:"seed"`
	Fmt           string `yaml:"fmt"`
	FmtContains   string `yaml:"fmt-contains"`
	MaxGoroutines int    `yaml:"max-goroutines"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sqlcquash",
	Short: "A cli tools to squash sql files into single files based on type.",
	Long: `A cli tools to squash sql files into single files based on type configured in config file.

Example:

sqlcquash combine -c config.yaml
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

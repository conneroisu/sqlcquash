package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// combineCmd represents the combine command
var combineCmd = &cobra.Command{
	Use:   "combine",
	Short: "Combines defined schemas, queries and seeds into single output files as defined in the config file. (`sqlcquash.yaml`)",
	Long: `
Combines defined schemas, queries and seeds into single output files as defined in the config file ("sqlcquash.yaml").
	`,
	Example: `
	sqlcaush combine
	`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		configPath, err := FindConfigFile()
		if err != nil {
			return err
		}
		cmd.Printf("Using config file: %s\n", configPath)
		configCnt, err := readFile(configPath)
		if err != nil {
			return err
		}
		cmd.Printf("Reading config file: %s\n", configCnt)
		config := Config{}
		err = yaml.Unmarshal([]byte(configCnt), &config)
		if err != nil {
			return err
		}
		cmd.Printf("Found %d instances in config file: %s\n", len(config.Dbs), configPath)
		dir, err := FindConfigDir(configPath)
		if err != nil {
			return err
		}
		cmd.Printf("Working in directory: %s\n", dir)
		err = os.Chdir(dir)
		if err != nil {
			return err
		}
		cmd.Printf("Changing directory to: %s\n", dir)
		for _, inst := range config.Dbs {
			schemas, err := catFiles(inst.SchemasPath)
			if err != nil {
				return err
			}
			queries, err := catFiles(inst.QueriesPath)
			if err != nil {
				return err
			}
			seeds, err := catFiles(inst.SeedsPath)
			if err != nil {
				return err
			}
			schemaFile, err := os.Create(inst.OutputSchema)
			if err != nil {
				return err
			}
			defer schemaFile.Close()
			if err != nil {
				return err
			}
			queryFile, err := os.Create(inst.OutputQueries)
			if err != nil {
				return err
			}
			defer queryFile.Close()
			if err != nil {
				return err
			}
			seedFile, err := os.Create(inst.OutputSeeds)
			if err != nil {
				return err
			}
			defer seedFile.Close()
			if err != nil {
				return err
			}
			_, err = schemaFile.WriteString(strings.Join(schemas, "\n"))
			if err != nil {
				return err
			}

			_, err = queryFile.WriteString(strings.Join(queries, "\n"))
			if err != nil {
				return err
			}
			_, err = seedFile.WriteString(strings.Join(seeds, "\n"))
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func catFiles(path string) ([]string, error) {
	files, err := filepath.Glob(path)
	if err != nil {
		return nil, err
	}
	var content []string
	for _, file := range files {
		cnt, err := readFile(file)
		if err != nil {
			return nil, err
		}
		content = append(content, cnt)
	}
	return content, nil
}

func readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
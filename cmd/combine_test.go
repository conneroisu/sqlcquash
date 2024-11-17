package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

const content = `
---
version: 1
dbs:
  - schemas_path: ./schemas/*.sql
    queries_path: ./queries/*.sql
    seeds_path: ./seeds/*.sql
    output_schema: ./combined/schema.sql
    output_queries: ./combined/queries.sql
    output_seeds: ./combined/seeds.sql
`

func TestMarshal(t *testing.T) {
	var c Config
	err := yaml.Unmarshal([]byte(content), &c)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, c.Dbs, 1)
}

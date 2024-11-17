package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

const content = `
# This file is used to configure sqlcquash.
#
---
version: 1
dbs:
  - schemas: ./schemas/*.sql
    queries: ./queries/*.sql
    seeds: ./seeds/*.sql
    schema: ./combined/schema.sql
    query: ./combined/queries.sql
    seed: ./combined/seeds.sql
`

func TestMarshal(t *testing.T) {
	var c Config
	err := yaml.Unmarshal([]byte(content), &c)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, c.Dbs, 1)
	assert.Equal(t, "./combined/schema.sql", c.Dbs[0].OutputSchema)
	assert.Equal(t, "./combined/queries.sql", c.Dbs[0].OutputQueries)
	assert.Equal(t, "./combined/seeds.sql", c.Dbs[0].OutputSeeds)
	assert.Equal(t, "./schemas/*.sql", c.Dbs[0].SchemasPath)
	assert.Equal(t, "./queries/*.sql", c.Dbs[0].QueriesPath)
	assert.Equal(t, "./seeds/*.sql", c.Dbs[0].SeedsPath)
}

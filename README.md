# sqlcquash

A simple tool to squash sql files into single files based on type.

## Installation

```bash
go install github.com/conneroisu/sqlcquash@latest
```

## Usage

```bash
sqlcaush combine -c config.yaml
```

```go
//go:generate github.com/conneroisu/sqlcquas@latest combine
```

## Configuration

The configuration file is `sqlcquash.yaml` and should be placed in the same directory as the executable.

```yaml
---
version: 1
dbs:
  - schemas: ./schemas/*.sql
    queries: ./queries/*.sql
    seeds: ./seeds/*.sql
    schema: ./combined/schema.sql
    query: ./combined/queries.sql
    seed: ./combined/seeds.sql
        
```

### Schemas

The `schemas` field is a glob pattern that matches the files to be combined into a single file.

### Queries

The `queries` field is a glob pattern that matches the files to be combined into a single file.

### Seeds

The `seeds` field is a glob pattern that matches the files to be combined into a single file.

### Schema

The `schema` field is the path to the file where the combined schema will be written.

### Query

The `query` field is the path to the file where the combined query will be written.

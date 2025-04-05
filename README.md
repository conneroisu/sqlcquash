# sqlcquash

A simple tool to squash sql files into single files based on type.

## Installation

### Go

```bash
go install github.com/conneroisu/sqlcquash@latest
```

### Nix

#### Flakes

```nix
{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    sqlcquash.url = "github:conneroisu/sqlcquash";
    flake-utils.url = "github:numtide/flake-utils";
  }
  outputs = { self, sqlcquash }:
  flake-utils.lib.eachDefaultSystem (system:
    let
      pkgs = import nixpkgs {
          inherit system; 
          overlays = [ sqlcquash.overlays.default ];
      };
    in
      {
        devShells.deffault = pkgs.mkShell {
          packages = with pkgs; [
            sqlcquash
          ];
        };
      }
  );
}
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
    fmt: sleek -i 4
    fmt-contains: queries
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

{
  description = "A golang cli sqlc schema, query and migration collector.";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?tag=24.11";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = {flake-utils, ...} @ inputs:
    flake-utils.lib.eachDefaultSystem (
      system: let
        overlays = [(final: prev: {final.go = prev.go_1_24;})];
        pkgs = import inputs.nixpkgs {inherit system overlays;};
        buildGoModule = pkgs.buildGoModule.override {go = pkgs.go_1_24;};
        buildWithSpecificGo = pkg: pkg.override {inherit buildGoModule;};
      in {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            # Go Tools
            go_1_24
            air
            templ
            pprof
            revive
            golangci-lint
            (buildWithSpecificGo gopls)
            (buildWithSpecificGo templ)
            (buildWithSpecificGo golines)
            (buildWithSpecificGo golangci-lint-langserver)
            (buildWithSpecificGo gomarkdoc)
            (buildWithSpecificGo gotests)
            (buildWithSpecificGo gotools)
            (buildWithSpecificGo reftools)
          ];
        };
        packages = {
          inherit (pkgs) sqlcquash;
          default = pkgs.sqlcquash;
        };

        overlays.default = final: prev: {
          sqlcquash = prev.buildGoModule {
            pname = "sqlcquash";
            version = "0.1.0";
            src = ./.;
            vendorHash = "sha256-/GsKVjvxQ97OrH04zM8tBnaElpOPrToYsgFWAtZLyLo";
          };
        };
      }
    );
}

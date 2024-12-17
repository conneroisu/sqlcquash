{
  description = "A golang cli sqlc schema, query and migration collector.";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?tag=24.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = nixpkgs.legacyPackages.${system};
      in {
        packages =  {
          sqlcquash = pkgs.buildGoModule {
            pname = "sqlcquash";
            version = "0.1.0";
            src = ./.;
            vendorHash = "sha256-/GsKVjvxQ97OrH04zM8tBnaElpOPrToYsgFWAtZLyLo";
          };
        };
        defaultPackage = self.packages.${system}.sqlcquash;
      }
    );
}

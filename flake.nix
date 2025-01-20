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
  }: let
    overlay = final: prev: {
      sqlcquash = prev.buildGoModule {
        pname = "sqlcquash";
        version = "0.1.0";
        src = ./.;
        vendorHash = "sha256-/GsKVjvxQ97OrH04zM8tBnaElpOPrToYsgFWAtZLyLo";
      };
    };
  in
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [overlay];
        };
      in {
        packages = {
          inherit (pkgs) sqlcquash;
          default = pkgs.sqlcquash;
        };
      }
    )
    // {
      overlays.default = overlay;
    };
}

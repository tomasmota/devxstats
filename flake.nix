{
  description = "devxstats flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    gitignore = {
      url = "github:hercules-ci/gitignore.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, flake-utils, gitignore }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system ; };
      in
      rec {
        packages.devxstats = pkgs.buildGoModule {
          name = "devxstats";
          src = gitignore.lib.gitignoreSource ./.;
          vendorHash = "sha256-CMpwsoKV5IsAum8BZGrbweuEntgeRYn5ODMgS4ocvL0=";
        };

        packages.default = packages.devxstats;

        devShell = pkgs.mkShellNoCC {
          packages = with pkgs; [
            go_1_22
            gotools
            gopls
            golangci-lint
          ];
        };
      }
    );
}

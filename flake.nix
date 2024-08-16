{
  description = "Solid project workspace";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let 
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages.default = 
        let
            inherit (pkgs) stdenv lib;
        in
        stdenv.mkDerivation {
          buildInputs = [
            pkgs.go
          ];
          name = "backend";
          src = "./.";
          cleanPhase = "rm -rf $sourceRoot/*";
          unpackPhase = ''
            sourceRoot=$PWD
            mkdir $sourceRoot
            cp -r $src/* $sourceRoot/
          '';
        };

        packages.frontend = 
        let
            inherit (pkgs) stdenv lib;
        in

        stdenv.mkDerivation {
          buildInputs = [
            pkgs.nodejs_20
          ];
          name = "frontend";
          src = "./front";
          cleanPhase = "rm -rf $sourceRoot/*";
          unpackPhase = ''
            sourceRoot=$PWD
            mkdir $sourceRoot
            cp -r $src/* $sourceRoot/
          '';
          buildPhase = "npm run biild";
        };
        devShells.default =
        let
            inherit (pkgs) stdenv lib;
        in
        pkgs.mkShell {
            name = "go";
            buildInputs = [
              pkgs.cowsay
              pkgs.lolcat
              pkgs.go
              pkgs.gotools
              pkgs.delve
              pkgs.protoc-gen-go
              pkgs.protobuf
            ];

            shellHook = ''
              export PATH=$(go env GOPATH)/bin:$PATH
              go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 
              echo "Go Shell" | cowsay | lolcat
            '';
          };
        devShells.frontend =
        let
            inherit (pkgs) stdenv lib;
        in
        pkgs.mkShell {
            name = "sveltkit";
            buildInputs = [
              pkgs.cowsay
              pkgs.lolcat
              pkgs.nodejs_20
            ];

            shellHook = ''
              echo "Sveltkit Shell" | cowsay | lolcat
            '';
          };
      }
      );
}

{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    wasm-spec = {
      url = "github:WebAssembly/spec";
      flake = false;
    };
  };

  outputs =
    {
      flake-utils,
      nixpkgs,
      wasm-spec,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        overlays = [ ];
        pkgs = import nixpkgs {
          inherit system overlays;
        };
        testsuite-json = pkgs.runCommand "json-testsuite" { } ''
          mkdir -p $out
          find ${wasm-spec}/test/core -name '*.wast' | while read f; do
            ${pkgs.wabt}/bin/wast2json $f -o "$out/$(basename $f).json"
          done
        '';
      in
      {
        devShells.default =
          with pkgs;
          mkShell {
            GOTOOLCHAIN = "local";
            WASMIUM_TEST_DIR = testsuite-json;
            buildInputs = [
              go
              gotools
              wabt
            ];
          };
      }
    );
}

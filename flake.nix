{
  description = "Advent Of Code 2024";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.05";
  };
  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
    in
    {
      packages.${system}.default = [ pkgs.figlet ] ;

      devShells.${system} = {
        go = pkgs.mkShell {
          name = "default";
          packages = [
		    self.packages.${system}.default
			pkgs.go_1_22
		  ];
          shellHook = "go version | figlet";
        };
      };
    };
}


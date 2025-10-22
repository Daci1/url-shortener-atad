{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-25.05";
  };

  outputs = { self, nixpkgs }: {
    devShells.x86_64-linux.default =
            let
                pkgs = import nixpkgs { system = "x86_64-linux"; };
            in
                pkgs.mkShell {
                    packages = with pkgs; [
                        nodejs_24
                        nodePackages."@angular/cli"

                    ];
                };
  };
}

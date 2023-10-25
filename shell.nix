let
  pkgs = import (fetchTarball {
    name = "nixos-23.05";
	url = "https://github.com/NixOS/nixpkgs/archive/b7cde1c47b73.tar.gz";
    sha256 = "0xry0farxln7cjas4zfip4s1hfl2x9lbwb0imsla0zqvzhvbz729";
  }) { };

in pkgs.mkShell {
	buildInputs = [
		pkgs.gnumake
		pkgs.go
		pkgs.golangci-lint
	];
	shellHook = ''
		if [ -e ~/.gitconfig ] && [ -f ~/.git-prompt.sh ]; then
			source ~/.git-prompt.sh
			PS1='\[\033[01;33m\]>>> nix:\w\[\033[00m\]$(__git_ps1 " %s")\n\[\033[00m\]$ '
		fi
	'';
}

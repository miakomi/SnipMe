{ pkgs }: {
    deps = [
	pkgs.git
        pkgs.nano
        pkgs.unzip
        pkgs.wget
        pkgs.go
        pkgs.gopls
    ];
}

## Build

> go build -o ../host/bin/main.exe -ldflags "-H=windowsgui -s -w"

You don't know what it means?

> $ go tool link

- ``-ldflags "-H=windowsgui"``: no console screen
- ``-s`` disable symbol table
- ```-w``` disable DWARF generation

see more:
  - [What does the w flag mean when passed in via the ldflags option to the go command?](https://stackoverflow.com/a/47967258/9935654)
  - [go build](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)


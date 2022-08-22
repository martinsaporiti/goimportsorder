## Go Imports Order

Go Imports Order is a tool that prevents you from having imports out of order mantaing the company imports (the mod package) at the end of the file.

### Installation

```bash
go install github.com/martinsaporiti/go-imports-order
```
### Usage

```bash
./goimportsorder -pattern=github.com/mycompany/myproject ./...
```
Also you can analize the imports order of a single file:

```bash
./goimportsorder -pattern=github.com/mycompany/myproject ./myproject/myproject.go
```

### Adding as private linter

Run this:
```bash
go build -buildmode=plugin -ldflags "-X plaugin.pattern=github.com/mycompany/myproject" plugin.go
```
then put this in your .golangci.yml:

```yaml
  custom:
    importsorder:
      path: ./plugin.so
      description: The description of the linter
      original-url: https://github.com/martinsaporiti/goimportsorder

```
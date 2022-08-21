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

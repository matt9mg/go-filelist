# File List
Returns a list of files.

### Installation
```
go get github.com/matt9mg/go-filelist
```

### Examples
OS Filesystem

```go
fileList := filelist.NewFileList()

fl, err := fileList.ListFromLocation("/path/to/base_directory")

if err != nil {
    log.Fatalln(err)
}

log.Println(fl)
```

Embed Filesystem

```go
//go:embed a_dir
var files embed.FS

fileList := filelist.NewFileList()

fl, err := fileList.ListFromFS(files, ".")

if err != nil {
    log.Fatalln(err)
}

log.Println(fl)
```

### LICENSE

This project is licensed under the MIT License - see the LICENSE.md file for details

### Disclaimer

We take no legal responsibility for anything this code is used for.
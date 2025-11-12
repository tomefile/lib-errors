# Tomefile Errors Library

Detailed, pretty-printed errors for tomefile.

## Example usage

```go
_, err := os.ReadFile("example.txt")
derr := &liberrors.DetailedError{
    Name:    liberrors.ERROR_IO,
    Details: err.Error(),
    Trace: []liberrors.TraceItem{
        {
            Name: "example_file.txt",
            Col:  1,
            Row:  1,
        },
        {
            Name: "parent_file.md",
            Col:  1,
            Row:  9,
        },
    },
    Context: "    1 |  this is some context",
}
derr.Print(os.Stderr)
```

Will produce:
```
[!] I/O Error
    in example_file.txt:1:1
    └─ from parent_file.md:9:1

    1 |  this is some context

[?] Details
    open example.txt: no such file or directory
```

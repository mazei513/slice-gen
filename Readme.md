# Slice Gen

A code generator that can generate slice tricks for a struct type.

## Usage

Pass in the slice type and run in the desired directory

```bash
go run github.com/mazei513/slice-gen Gopher
```

That will create a file named `gopher_gen.go` in the same directory.

## Functions created

Functions taken from https://github.com/golang/go/wiki/SliceTricks

- Delete
- DeleteRange
- Insert
- Push
- Pop
- Unshift
- Shift
- Filter

## Further enhancements

- Support for basic types
- Support for pointer types
- Add tests

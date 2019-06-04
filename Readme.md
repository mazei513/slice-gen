# Slice Gen

A code generator that can generate slice tricks for a struct type.

## Usage

### CLI

Pass in the slice type and run in the desired directory

```bash
go run github.com/mazei513/slice-gen Droid
```

That will create a file named `droid_gen.go` in the same directory.

### Go Generate

Add the following line at the top of a Go file to use with `go generate`. The example is for a file `droid.go`.

```Go
//go:generate go run github.com/mazei513/slice-gen Droid
```

In the same directory as the file, run:

```bash
go generate
```

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

- Add tests for nil array situations
- Add checks for valid arguments into slice functions
- Error out of slice functions instead of panic

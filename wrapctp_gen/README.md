# Generate

__For C and platform__

```
go run wrapper_gen.go -csys macos -lang c -srcpath ../ctpapi/ -outpath ../csrc/macos

go run wrapper_gen.go -csys windows -lang c -srcpath ../ctpapi/ -outpath ../csrc/windows

go run wrapper_gen.go -csys linux -lang c -srcpath ../ctpapi/ -outpath ../csrc/linux
```

__For Python and platform__

```
go run wrapper_gen.go -csys macos -lang python -srcpath ../ctpapi/ -outpath ../pyctp/macos

go run wrapper_gen.go -csys windows -lang python -srcpath ../ctpapi/ -outpath ../pyctp/windows

go run wrapper_gen.go -csys linux -lang python -srcpath ../ctpapi/ -outpath ../pyctp/linux
```

__For Golang and platform__

```
go run wrapper_gen.go -csys macos -lang golang -srcpath ../ctpapi/ -outpath ../ctpgo/

go run wrapper_gen.go -csys windows -lang golang -srcpath ../ctpapi/ -outpath ../ctpgo/

go run wrapper_gen.go -csys linux -lang golang -srcpath ../ctpapi/ -outpath ../ctpgo/
```

module example

go 1.24.0

toolchain go1.24.7

replace ctpgo => ../ctpgo

require (
	ctpgo v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
)

require (
	github.com/ebitengine/purego v0.9.1 // indirect
	golang.org/x/text v0.32.0 // indirect
)

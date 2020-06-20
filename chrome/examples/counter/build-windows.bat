@echo off
go generate
go build -ldflags "-H windowsgui" -o chromeui-example.exe

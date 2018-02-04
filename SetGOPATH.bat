@ECHO Setting GOPATH to current path...
@ECHO.
:: @SET GOPATH=%cd%\ :: alternate method just using the current directory
@FOR /F %%I IN ('git rev-parse --show-toplevel') DO @SET GOPATH=%%I
@SET GOPATH
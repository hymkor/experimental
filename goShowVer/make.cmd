@echo off
setlocal
    where rsrc >nul || go get github.com/akavel/rsrc
    where rsrc >nul || call :setgobin
    if not exist rsrc.syso rsrc -manifest test.manifest -o rsrc.syso
    go build -ldflags="-H windowsgui"
endlocal
exit /b

:setgobin
    setlocal
    if "%GOPATH%" == "" (
        set "GOBIN=%USERPROFILE%\go\bin"
    ) else (
        for /F "delims=;" %%I in ("%GOPATH%") ; do set "GOBIN=%%I"
    )
    endlocal & set "PATH=%GOBIN%;%PATH%"
exit/b

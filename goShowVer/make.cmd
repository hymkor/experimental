@echo off
setlocal
    where rsrc >nul || go get github.com/akavel/rsrc
    where rsrc >nul || call :setgobin
    if not exist rsrc.syso rsrc -manifest test.manifest -o rsrc.syso
    call :"%1"
endlocal
exit /b

:""
    go build -ldflags="-H windowsgui"
    exit /b

:"update"
    go get -u github.com/akavel/rsrc
    go get -u github.com/lxn/walk
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

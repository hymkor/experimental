@echo off
call :"%1"
exit /b

:"upgrade"
    setlocal
    for %%I in (%CD%) do set "EXE=%%~nI"
    for /F %%I in ('where %EXE%') do (
        if not "%CD%\%EXE%.exe" == "%%I" copy /-Y "%CD%\%EXE%.exe" %%I
    )
    endlocal
    exit /b

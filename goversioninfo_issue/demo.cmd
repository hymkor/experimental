@setlocal
@set PROMPT=$G
go build
for /F %%I in ('cd') do set EXE=%%~nI
powershell -ExecutionPolicy RemoteSigned -file %~dp0showver.ps1 %EXE%.exe
@endlocal

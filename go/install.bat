@echo off
:: go.exe may exist on the HKLM, so you should run as admin.
echo "Make sure you are run as an admin."
pause
cd %~dp0
echo %cd%
go build -o ../host/bin/main.exe -ldflags "-H=windowsgui -s -w"
robocopy  .  ../host/bin/ manifest.json
pause > nul && echo done

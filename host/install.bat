@echo off
:: If you add "/f" then you can force write.
REG ADD "HKCU\Software\Google\Chrome\NativeMessagingHosts\com.google.chrome.extension.thunder" ^
 /ve /t REG_SZ ^
 /d "%~dp0manifest.json"

:: Run with the admin. Otherwise, you will get the error of permission denied.
REG ADD "HKLM\Software\Google\Chrome\NativeMessagingHosts\com.google.chrome.extension.thunder" /ve /t REG_SZ /d "%~dp0manifest.json"
pause > nul && echo done

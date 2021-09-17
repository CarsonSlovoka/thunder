:: Deletes the entry created by install.bat
@echo off
REG DELETE "HKCU\Software\Google\Chrome\NativeMessagingHosts\com.google.chrome.extension.thunder" /f
REG DELETE "HKLM\Software\Google\Chrome\NativeMessagingHosts\com.google.chrome.extension.thunder" /f
pause > nul && echo done

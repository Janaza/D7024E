﻿$file = "$pwd\main.exe" 
$bootstrap = "10.0.75.1:1"
go build .\main.go
Start-Sleep -Seconds 2
Start-Process powershell.exe -ArgumentList $file, "1"
Start-Sleep -Milliseconds 1000
For ($i=2; $i -le 30; $i++) {
Start-Process -NoNewWindow powershell.exe -ArgumentList $file, $i, $bootstrap
Start-Sleep -Milliseconds 500
}
Start-Sleep -Milliseconds 500
Start-Process powershell.exe -ArgumentList $file, "40", $bootstrap

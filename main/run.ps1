$file = "$pwd\main.exe" 
$bootstrap = "-bootstrapIP localhost:1"
go build .\main.go
Start-Sleep -Seconds 2
Start-Process powershell.exe -ArgumentList $file, "-port 1"
Start-Sleep -Milliseconds 1000
For ($i=2; $i -le 30; $i++) {
Start-Process -NoNewWindow powershell.exe -ArgumentList $file, "-port $i", $bootstrap
Start-Sleep -Milliseconds 500
}
Start-Sleep -Milliseconds 500
Start-Process powershell.exe -ArgumentList $file, "-port 40", $bootstrap

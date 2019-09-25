$file = "$pwd\main.exe" 
$bootstrap = "localhost:1"
go build .\main.go
Start-Sleep -Seconds 2
Start-Process powershell.exe -ArgumentList $file, "1"
Start-Sleep -Milliseconds 100
For ($i=2; $i -le 30; $i++) {
Start-Process -NoNewWindow powershell.exe -ArgumentList $file, $i, $bootstrap
Start-Sleep -Milliseconds 200
}
Start-Sleep -Milliseconds 500
Start-Process powershell.exe -ArgumentList $file, "40", $bootstrap

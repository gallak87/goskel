# --> https://medium.com/blokur/how-to-implement-a-grpc-client-and-server-in-typescript-fa3ac807855e
# Kill any previously running servers
$existing1 = Get-NetTCPConnection -LocalPort 9090 -EA SilentlyContinue
$existing2 = Get-NetTCPConnection -LocalPort 8080 -EA SilentlyContinue
if ($existing1) {
  kill -f (Get-Process -Id $existing1.OwningProcess) -EA SilentlyContinue
}
if ($existing2) {
  kill -f (Get-Process -Id $existing2.OwningProcess) -EA SilentlyContinue
}

$cmd1 = "grpcwebproxy"
$args1 = "--backend_addr=localhost:9090 --allow_all_origins --run_tls_server=false"
$cmd2 = "go"
$args2 = "run ./cmd/server"

$id1 = start-process $cmd1 -PassThru -ArgumentList $args1
$id2 = start-process $cmd2 -PassThru -ArgumentList $args2

Write-Host "Press any key to exit.."
Read-Host

kill -f $id1
kill -f $id2
kill -f (get-process server).Id -EA SilentlyContinue

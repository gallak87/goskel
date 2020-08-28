$web = "$pwd/../web"
$api = "$web/src/api"
# Clean previously generated files
rm -force $api/*

# Generate new clients
# TODO: dynamic per pb? hard coding userpb for now
protoc --proto_path .\userpb .\userpb\user.proto `
  "-IC:/Users/georg/.protoc" `
  "-IC:/Users/georg/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.0/third_party/googleapis/google/api" `
  --go_out=plugins=grpc:userpb `
  "--plugin=protoc-gen-ts=$web/node_modules/.bin/protoc-gen-ts.cmd" `
  --js_out="import_style=commonjs,binary:$api" `
  --ts_out=service=grpc-web:$api

# pre-fix with /* eslint-disable */ to bypass lint issue
$disable = '/* eslint-disable */' + [System.Environment]::NewLine + [System.Environment]::NewLine
ls $api | % {
  $output = Get-Content $_.FullName -Raw
  "$disable" + $output | Set-Content $_.FullName
}
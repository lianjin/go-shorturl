#/bin/bash

project_dir="$(cd "$(dirname "$0")/.." && pwd)"

cd $project_dir
go build -ldflags="-s -w" -o gsurl cmd/gsurl.go

echo "build gsurl done."
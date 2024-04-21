#!/bin/bash

VERSION=0.0.2
APPNAME=telebot_app_serverless_sendtome
echo "package constvar" > ./internal/constvar/version.go
echo "const(APP_NAME = \"${APPNAME}\"" >> ./internal/constvar/version.go
echo "APP_VERSION = \"${VERSION}\")" >> ./internal/constvar/version.go
go fmt ./internal/constvar


rm ./main/main

#git init #
git add .
git commit -m "v${VERSION} debug"
#git branch -M main #
git push -u origin main

git tag "v${VERSION}"
git push --tags -u origin main


# !/bin/bash

echo "++++++++++ go mod tidy ++++++++++"
go mod tidy

echo "++++++++++ go mod build ++++++++++"
go mod build

echo "++++++++++ starting server ++++++++++"
./banking
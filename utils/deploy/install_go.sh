#!/usr/bin/env bash

echo Removing folders...
rm -rf /godist
rm -rf /usr/local/go

echo Removing folders...
rm /usr/bin/go
rm /usr/bin/go-fmt

echo Creating folders...
mkdir /godist
cd /godist

echo Downloading go source...
wget https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz

echo Extracting go source...
sudo tar -xf go1.7.4.linux-amd64.tar.gz
sudo mv go /usr/local

ln -s /usr/local/go/bin/go /usr/bin/go
ln -s /usr/local/go/bin/go-fmt /usr/bin/go-fmt

export GOROOT=/usr/local/go

rm go1.7.4.linux-amd64.tar.gz
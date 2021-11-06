#!/usr/bin/env bash
set -e
# install
sudo apt-get update
sudo apt-get dist-upgrade -y
sudo apt-get install -y git
curl -fsSL https://deb.nodesource.com/setup_16.x | sudo -E bash -
sudo apt-get install -y nodejs
npm install pm2 -g

# project
cd /opt
git clone https://mstip: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@github.com/mstip/solid-pancake.git
cd solid-pancake/agent
cp .env.prod .env
echo "SERVER_ID=3" >> .env
npm install
pm2 start src/index.mjs

curl --header "Content-Type: application/json" --header "Authorization: Basic c29saWRwYW5jYWtlOnNvbGlkcGFuY2FrZQ==" --request POST --data '{"status":"done"}' http://localhost:3000/api/install/3
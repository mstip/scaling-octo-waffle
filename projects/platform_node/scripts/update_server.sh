#!/usr/bin/env bash
sudo apt-get update
sudo apt-get dist-upgrade -y


cd /opt/solid-pancake
git checkout .
git pull origin main
cd service
cp .env.prod .env
npm install
pm2 restart all
#!/usr/bin/env bash
sudo apt-get update
sudo apt-get dist-upgrade -y


cd /opt/isms
git checkout .
git pull origin master
cd frontend
npm install
npm run build
cd ../backend
npm install
npm run build
cp .env.prod .env
pm2 restart all
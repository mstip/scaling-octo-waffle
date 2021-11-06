#!/usr/bin/env bash

# install
sudo apt-get update
sudo apt-get dist-upgrade -y
sudo apt-get install -y git nginx snapd postgresql postgresql-contrib
sudo snap install core; sudo snap refresh core
sudo snap install --classic certbot
curl -fsSL https://deb.nodesource.com/setup_16.x | sudo -E bash -
sudo apt-get install -y nodejs
npm install pm2 -g


# postgres
sudo -u postgres createuser -s isms
sudo -u postgres createdb isms
sudo adduser isms
# ALTER USER isms PASSWORD 'isms';

# project
cd /opt
git clone https://mstip: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@github.com/mstip/isms.git
cd isms/frontend
npm install
npm run build
cd ../backend
npm install
npm run build
cp .env.prod .env
pm2 start dist/main.js


# nginx
cd ..
rm /etc/nginx/sites-enabled/default
rm /etc/nginx/sites-available/default
cp /opt/isms/scripts/isms.conf /etc/nginx/sites-available
ln -s /etc/nginx/sites-available/isms.conf /etc/nginx/sites-enabled/isms.conf
service nginx restart

# https
sudo ln -s /snap/bin/certbot /usr/bin/certbot
sudo certbot --nginx
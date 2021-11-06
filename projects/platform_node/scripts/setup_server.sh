#!/usr/bin/env bash

# install
sudo apt-get update
sudo apt-get dist-upgrade -y
sudo apt-get install -y git nginx snapd mysql-server
sudo snap install core; sudo snap refresh core
sudo snap install --classic certbot
curl -fsSL https://deb.nodesource.com/setup_16.x | sudo -E bash -
sudo apt-get install -y nodejs
npm install pm2 -g


echo "CREATE USER 'solidpancake'@'localhost' IDENTIFIED BY 'solidpancake';" | mysql
echo "CREATE DATABASE solidpancake;" | mysql
echo "GRANT ALL PRIVILEGES ON solidpancake.* TO 'solidpancake'@'localhost';" | mysql


# project
cd /opt
git clone https://mstip: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@github.com/mstip/solid-pancake.git
cd solid-pancake/service
cp .env.prod .env
npm install
pm2 start src/index.mjs


# nginx
rm /etc/nginx/sites-enabled/default
rm /etc/nginx/sites-available/default
cp /opt/solid-pancake/scripts/solid-pancake.conf /etc/nginx/sites-available
ln -s /etc/nginx/sites-available/solid-pancake.conf /etc/nginx/sites-enabled/solid-pancake.conf
service nginx restart

# https
snap install core &&  snap refresh core
snap install --classic certbot
ln -s /snap/bin/certbot /usr/bin/certbot
certbot --nginx --non-interactive --agree-tos -m marcstipcevic@gmail.com --domain solid-pancake.w00p.xyz
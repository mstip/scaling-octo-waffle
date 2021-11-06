ARG VARIANT=16-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/javascript-node:0-${VARIANT}


RUN apt-get update && apt-get dist-upgrade -y
# install mongodb
RUN wget -qO - https://www.mongodb.org/static/pgp/server-5.0.asc | sudo apt-key add - \
    && echo "deb http://repo.mongodb.org/apt/debian buster/mongodb-org/5.0 main" | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list \
    && apt-get update \
    && apt-get install -y mongodb-org \
    && mkdir -p /data/db
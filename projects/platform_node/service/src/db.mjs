import fastifyPlugin from 'fastify-plugin'

import sequelizePkg from 'sequelize';
const { Sequelize, DataTypes } = sequelizePkg;



async function dbConnector(fastify, options) {
    const db = new Sequelize(
        process.env.DB_NAME, process.env.DB_USER, process.env.DB_PASS, {
        host: process.env.DB_HOST,
        dialect: 'mysql'
    });

    db.define('server', {
        name: {
            type: DataTypes.STRING,
            allowNull: false
        },
        ip: {
            type: DataTypes.STRING,
            allowNull: true
        },
        domain: {
            type: DataTypes.TEXT,
            allowNull: true
        },
        agentInstalled:
        {
            type: DataTypes.BOOLEAN,
            allowNull: false,
            defaultValue: false
        },

    }, {
    });

    await db.sync({ alter: true });
    fastify.decorate('db', db)
}

export default fastifyPlugin(dbConnector)
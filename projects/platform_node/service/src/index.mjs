import Fastify from 'fastify';
import pointOfView from 'point-of-view';
import nunjucks from 'nunjucks';
import fastifyStatic from 'fastify-static';
import fastifyBasicAuth from 'fastify-basic-auth';

import dotenv from 'dotenv';
dotenv.config();

import { fileURLToPath } from 'node:url';
import { dirname, join as pathJoin } from 'node:path';

import db from './db.mjs';
import web from './web.mjs'
import api from './api.mjs'


const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const fastify = Fastify({
  logger: true
})

fastify.register(db)
fastify.register(pointOfView, { engine: { nunjucks } });

fastify.register(fastifyBasicAuth, { validate, authenticate: { realm: 'solidpancake' } })
function validate(username, password, req, reply, done) {
  if (username === process.env.BASIC_AUTH_NAME && password === process.env.BASIC_AUTH_PASS) {
    done()
  } else {
    done(new Error('unauthorized'))
  }
}

fastify.register(fastifyStatic, {
  root: pathJoin(__dirname, '../public'),
  prefix: '/public/',
});

fastify.register(import('fastify-formbody'))

fastify.register(web);
fastify.register(api, {prefix: '/api'});



fastify.listen(3000, function (err, address) {
  if (err) {
    fastify.log.error(err)
    process.exit(1);
  }
});
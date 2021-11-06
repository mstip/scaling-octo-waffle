export default async function routes(fastify, options) {
    fastify.after(() => {
        fastify.addHook('onRequest', fastify.basicAuth);

        fastify.get('/', function (request, reply) {
            reply.view('./templates/index.njk', {});
        });

        fastify.get('/server', async function (request, reply) {
            const servers = await fastify.db.models.server.findAll();
            reply.view('./templates/server_index.njk', { servers });
        });
        fastify.get('/server/edit/:serverId', async function (request, reply) {
            const server = await fastify.db.models.server.findByPk(request.params.serverId);
            if (server === null) {
                reply.callNotFound();
                return;
            }
            reply.view('./templates/server_edit.njk', {
                isEdit: true,
                id: server.id,
                name: server.name,
                ip: server.ip,
                domain: server.domain,
                agentInstalled: server.agentInstalled,
                serviceDomain: process.env.SERVICE_DOMAIN
            });
        });

        fastify.post('/server/edit/:serverId', async function (request, reply) {
            await fastify.db.models.server.update({ ...request.body }, {
                where: {
                    id: request.params.serverId
                }
            });
            reply.redirect(`/server/edit/${request.params.serverId}`);
        });

        fastify.get('/server/create', function (request, reply) {
            reply.view('./templates/server_edit.njk', { isEdit: false });
        });
        fastify.post('/server/create', async function (request, reply) {
            await fastify.db.models.server.create({ ...request.body })
            reply.redirect('/server');
        });
    });
};

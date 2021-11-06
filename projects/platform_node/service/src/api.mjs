export default async function routes(fastify, options) {
    fastify.after(() => {
        fastify.addHook('onRequest', fastify.basicAuth);
        
        fastify.get('/install/:serverId/install.sh', function (request, reply) {
            reply.send("insert agent script here");
        });

        fastify.post('/install/:serverId', async function (request, reply) {
            if(request.body.status === 'done') {
                await fastify.db.models.server.update({ agentInstalled: true }, {
                    where: {
                        id: request.params.serverId
                    }
                });
            }
            reply.send();
        });
    });
};
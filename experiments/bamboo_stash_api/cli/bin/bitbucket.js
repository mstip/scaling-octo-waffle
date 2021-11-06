const BitBucketService = require('../src/bitbucket_service');

const command = async function() {
    const service = new BitBucketService( 'https://stash.traviangames.com', 'NzYxMjkwOTc5MTg3OtI1wO5Bc1jcgpGRd4vxoiMbAUIM');
    // console.log(await service.getAllProjects());
    // console.log(await (await service.getAllRepositories('SOL')).data.length);
    console.log(await (await service.getAllPullRequests('SOL','katalon')));
}

command();
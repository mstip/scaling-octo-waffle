// time from env deploy to pr merge, the other way around because bamboo keeps no history on deploys ...
const BitBucketService = require('../src/bitbucket_service');
const BambooService = require('../src/bamboo_service');

const project = 'TGP';
const repo = 'tgp_checkout';
const planId = 'PAY-TGPCHECK';

const command = async function () {
    const bitBucket = new BitBucketService('https://stash.traviangames.com', 'NzYxMjkwOTc5MTg3OtI1wO5Bc1jcgpGRd4vxoiMbAUIM');
    const bamboo = new BambooService('https://bamboo.traviangames.com', 'MjkxNTA1ODk3MzQ3OjGMJzhqwr4t+qsypH1k4+L0Z2ZE');

    const results = [];

    // get deployment infos for a plan in bamboo
    const deploy = await bamboo.getDeployForPlan(planId);
    for (const environment of deploy.data.environments) {
        // get deployment results for each env (dev, stage, prod ...)
        const deployEnvResults = await bamboo.getDeploymentEnvironmentResults(environment.id);
        for (const result of deployEnvResults.data) {
            // their are multiple of them always with the same content so use the first
            const item = result.deploymentVersion.items[0];
            // with this item we have the build id and we can grab the result because the key is unique with project name etc in it
            const planResult = await bamboo.getPlanResults(item.planResultKey.key);
            // from the build we get the commitId
            const commit = await bitBucket.getCommit(project, repo, planResult.data.vcsRevisionKey);
            // a merge commit has always the same commit message and this is the only way to get the prid so parse it !
            const prId = commit.data.message.split(':')[0].split('#')[1].trim();
            // with the pr id we finally could make the connection and get the time
            const pullRequest = await bitBucket.getPullRequest(project, repo, prId);
            results.push({
                project, repo, planId,
                versionName: result.deploymentVersionName,
                envId: environment.id,
                envName: environment.name,
                buildkey: item.planResultKey.key,
                buildNumber: item.planResultKey.resultNumber,
                finishDate: new Date(result.finishedDate),
                vcsRevisionKey: planResult.data.vcsRevisionKey,
                pullRequestId: prId,
                pullRequestFinish: new Date(pullRequest.data.closedDate),
                pullRequestTitle: pullRequest.data.title,
                fromPrToProd: result.finishedDate - pullRequest.data.closedDate,
                fromPrToProdHours: (result.finishedDate - pullRequest.data.closedDate) / (1000 * 60 * 60),
                fromPrToProdDays: (result.finishedDate - pullRequest.data.closedDate) / (1000 * 60 * 60 * 24)
            });
            break;
        }
    }

    console.log(results);
}

command();
// Time from Pull request creation to prod deploy
// example tgp checkout


const BitBucketService = require('../src/bitbucket_service');
const BambooService = require('../src/bamboo_service');


const project = 'TGP';
const repo = 'tgp_checkout';
const planId = 'PAY-TGPCHECK';

const command = async function () {
    const bitBucket = new BitBucketService('https://stash.traviangames.com', 'NzYxMjkwOTc5MTg3OtI1wO5Bc1jcgpGRd4vxoiMbAUIM');
    const bamboo = new BambooService('https://bamboo.traviangames.com', 'MjkxNTA1ODk3MzQ3OjGMJzhqwr4t+qsypH1k4+L0Z2ZE');

    /*
    bamboo speichert keine versions history auf seinen deploys deshalb muss man es von der bamboo env zurück machen und rechnen und sich einfach täglich merken wenn was neues kam ...
    */ 


    // 1. get pull requests state closed
    // 2. get commit history connect pr by commit message pull request #408: Bugfix/PAY-1323


    /*
    Bugfix/PAY-1323 tgp 2021-08-30T06:07:48.000Z 177635333 Development 664 TGP-CHECKOUT-build-664 2021-09-02T08:38:23.000Z
    Bugfix/PAY-1323 tgp 2021-08-30T06:07:48.000Z 177635334 Sandbox 664 TGP-CHECKOUT-1.57.0 2021-08-16T07:01:19.000Z
    Bugfix/PAY-1323 tgp 2021-08-30T06:07:48.000Z 177635335 Prod 664 TGP-CHECKOUT-1.57.0 2021-09-03T08:59:45.000Z
    */
    const matches = [];
    // work with latest PR
    const pullRequests = (await bitBucket.getAllPullRequests(project, repo)).data;
    for (const prIndex in pullRequests) {
        const pullRequest = pullRequests[prIndex];
        console.log(prIndex);

        const prId = pullRequest.id;
        const mergeCommits = await bitBucket.getMergeCommits(project, repo);
        let commitId = null;
        for (const commit of mergeCommits.data) {
            if (commit.message.startsWith(`Pull request #${prId}:`)) {
                commitId = commit.id;
                break;
            }
        }

        if (commitId === null) {
            console.log('could not find merge commit for', pullRequest.id, pullRequest.title, new Date(pullRequest.closedDate))
            continue;
        }


        let buildNumber = null;
        // build mit pr verknüpfen und dann schauen wo deployed
        const planResults = await bamboo.getPlanResults(planId);
        for (const planResult of planResults.data) {
            const buildResult = await bamboo.getPlanBuildResult(planId, planResult.buildNumber);
            if (buildResult.data.vcsRevisionKey === commitId) {
                buildNumber = buildResult.data.number;
                break;
            }
        }

        if (buildNumber === null) {
            console.log('could not find buildnumber', pullRequest.id, pullRequest.title, new Date(pullRequest.closedDate), commitId)
            continue;
        }


        const deploy = await bamboo.getDeployForPlan(planId);
        for (const environment of deploy.data.environments) {
            const results = await bamboo.getDeploymentEnvironmentResults(environment.id);
            for (const result of results.data) {
                // console.log(pullRequest.title, new Date(pullRequest.closedDate), environment.id, environment.name, buildNumber, result.deploymentVersionName, new Date(result.finishedDate));
                let matched = false;
                for (const item of result.deploymentVersion.items) {
                    if (buildNumber === item.planResultKey.resultNumber) {
                        // console.log(item.planResultKey.key, item.planResultKey.resultNumber)
                        matched = true;
                        matches.push([
                            pullRequest.title,
                            new Date(pullRequest.closedDate),
                            environment.id,
                            environment.name,
                            buildNumber,
                            result.deploymentVersionName,
                            new Date(result.finishedDate),
                            item.planResultKey.key,
                            item.planResultKey.resultNumber])
                    break;

                    }
                }
                if(matched) {
                    break;
                }
            }
        }

        console.log('no match ...')
    }
    console.log('=================================================================');
    console.log(matches);
}

command();
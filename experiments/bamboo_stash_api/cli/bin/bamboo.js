const BambooService = require('../src/bamboo_service');

const command = async function() {
    const service = new BambooService( 'https://bamboo.traviangames.com', 'MjkxNTA1ODk3MzQ3OjGMJzhqwr4t+qsypH1k4+L0Z2ZE');
    // console.log(await service.getAllPlans());
    // console.log(await service.getAllDeploymentPlans());
    // console.log(await service.getDeploymentEnvironmentResults('190480461'));
    console.log(await service.getPlanResult('BIL2-BIL-81'));

}

command();
const axios = require('axios');
const _ = require('lodash');


module.exports = class BambooService {
    constructor(bambooUrl, accessToken) {
        bambooUrl += '/rest/api/latest/';

        this._BambooClient = axios.create({
            baseURL: bambooUrl,
            headers: {
                'Authorization': `Bearer ${accessToken}`
            },
            params: {
                'max-results': 100
            }
        });
    }

    async getAllPlans() {
        try {
            const response = await this._BambooClient.get('plan');
            return this._toReturnData(response, 'data.plans.plan');
        } catch (error) {
            return this._toReturnData(error);
        }
    }

    async getAllDeploymentPlans() {
        try {
            const response = await this._BambooClient.get('deploy/project/all');
            return this._toReturnData(response, 'data');
        } catch (error) {
            return this._toReturnData(error);
        }
    }

    async getDeploymentEnvironmentResults(envId) {
        try {
            const response = await this._BambooClient.get(`deploy/environment/${envId}/results`);
            return this._toReturnData(response, 'data.results');
        } catch (error) {
            return this._toReturnData(error);
        }
    }


    async getPlanResults(planId) {
        try {
            const response = await this._BambooClient.get(`result/${planId}`);
            return this._toReturnData(response, 'data');
        } catch (error) {
            return this._toReturnData(error);
        }
    }


    async getPlanBuildResult(planId, buildNumber) {
        try {
            const response = await this._BambooClient.get(`result/${planId}-${buildNumber}`);
            return this._toReturnData(response, 'data');
        } catch (error) {
            return this._toReturnData(error);
        }
    }

    async getDeployForPlan(planId) {
        try {
            let response = await this._BambooClient.get(`deploy/project/forPlan?planKey=${planId}`);
            const deployId = response.data[0].id;
            response = await this._BambooClient.get(`deploy/project/${deployId}`);
            return this._toReturnData(response, 'data');
        } catch (error) {
            return this._toReturnData(error);
        }
    }

 

    _toReturnData(response, pathToData) {
        if (response.status === 200) {
            return {
                success: true,
                error: null,
                data: _.get(response, pathToData)
            };
        }
        return {
            success: false,
            error: {
                statusCode: response.response.status,
                detail: response.message
            }
        };
    }

}

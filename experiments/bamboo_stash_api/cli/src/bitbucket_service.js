const axios = require('axios');


module.exports = class BitBucketService {
    constructor(bitbucketUrl, accessToken) {
        bitbucketUrl += '/rest/api/1.0/';

        this._bitBucketClient = axios.create({
            baseURL: bitbucketUrl,
            headers: {
                'Authorization': `Bearer ${accessToken}`
            },
            params: {
                limit: 1000
            }
        });
    }

    async getAllProjects() {
        try {
            const response = await this._bitBucketClient.get('projects');
            return this._toReturnData(response);
        } catch (error) {
            return this._toReturnData(error);
        }
    }

    async getAllRepositories(project) {
        try {
            const response = await this._bitBucketClient.get(`projects/${project}/repos`);
            return this._toReturnData(response);
        } catch (error) {
            return this._toReturnData(error);
        }
    }

    async getAllPullRequests(project, repo) {
        try {
            const response = await this._bitBucketClient.get(`projects/${project}/repos/${repo}/pull-requests?state=ALL`);
            return this._toReturnData(response);
        } catch (error) {
            return this._toReturnData(error);
        }
    }

    async getMergeCommits(project, repo) {
        try {
            const response = await this._bitBucketClient.get(`projects/${project}/repos/${repo}/commits?merge=only`);
            return this._toReturnData(response);
        } catch (error) {
            return this._toReturnData(error);
        }
    }

    async getCommit(project, repo, commitId) {
        try {
            const response = await this._bitBucketClient.get(`projects/${project}/repos/${repo}/commits/${commitId}`);
            return this._toReturnData(response);
        } catch (error) {
            return this._toReturnData(error);
        }
    }

    async getPullRequest(project, repo, prId) {
        try {
            const response = await this._bitBucketClient.get(`projects/${project}/repos/${repo}/pull-requests/${prId}`);
            return this._toReturnData(response);
        } catch (error) {
            return this._toReturnData(error);
        }
    }

    _toReturnData(response) {
        if (response.status === 200) {
            return {
                success: true,
                error: null,
                data: response.data.values === undefined ? response.data : response.data.values
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

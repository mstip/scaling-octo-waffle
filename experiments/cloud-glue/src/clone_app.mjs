export default class CloneApp {
    constructor(doToken, output) {
        this.doToken = doToken;
        this.output = output;
        this.cloneAppAreaEl = document.querySelector('.cloneAppArea');
        this.cloneAppToggleEl = document.getElementById('cloneAppToggle');
        this.cloneAppToggleEl.onclick = event => this.#onToggle(event);
        this.cloneAppIdEl = document.getElementById('cloneAppId');
        this.cloneAppIdEl.onkeyup = () => this.#onIdChange();
        this.cloneAppNameEl = document.getElementById('cloneAppName');
        this.cloneAppBranchEl = document.getElementById('cloneAppBranch');
        this.cloneAppDomainEl = document.getElementById('cloneAppDomain');
        this.cloneAppEl = document.getElementById('cloneApp');
        this.cloneAppEl.onclick = () => this.#cloneApp();
        this.fetchedApp = null;
    }

    #onToggle() {
        if (this.cloneAppAreaEl.style.display === 'none') {
            this.cloneAppAreaEl.style.display = 'block';
        } else {
            this.cloneAppAreaEl.style.display = 'none';
            this.#resetInputs();
        }
    }

    async #onIdChange() {
        const id = this.cloneAppIdEl.value;
        if (id.length < 36) {  // uuid length is 36
            this.output.setText(`${id} is not a valid uuid`);
            return;
        }

        const token = this.doToken.getToken();
        if (token.length < 3) {
            this.output.setText('Error: please set DO Token');
            return;
        }
        try {
            this.output.setText('request sent, please wait ...');
            const result = await fetch('https://api.digitalocean.com/v2/apps/' + id,
                {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`
                    }
                })
                .then(response => {
                    if (!response.ok) {
                        throw new Error('HTTP status ' + response.status);
                    }
                    return response.json();
                });
            this.fetchedApp = result.app;

            this.output.setText({
                id: this.fetchedApp.id,
                name: this.fetchedApp.spec.name,
                repo: this.fetchedApp.spec.services[0].github.repo,
                branch: this.fetchedApp.spec.services[0].github.branch,
                domains: this.fetchedApp.spec.domains ? this.fetchedApp.spec.domains.map(domain => domain.domain) : []
            });
            this.#setInputs(
                this.fetchedApp.spec.name,
                this.fetchedApp.spec.services[0].github.branch,
                this.fetchedApp.spec.domains ? this.fetchedApp.spec.domains.map(domain => domain.domain)[0] : ''
            );
        } catch (e) {
            this.fetchedApp = null;
            this.output.setText(e);
            this.#resetInputs();
        }

    }

    #resetInputs() {
        this.cloneAppIdEl.value = '';
        this.cloneAppNameEl.value = '';
        this.cloneAppBranchEl.value = '';
        this.cloneAppDomainEl.value = '';
        this.cloneAppNameEl.disabled = true;
        this.cloneAppBranchEl.disabled = true;
        this.cloneAppDomainEl.disabled = true;
        this.cloneAppEl = true;
    }

    #setInputs(name, branch, domain) {
        this.cloneAppNameEl.value = name;
        this.cloneAppBranchEl.value = branch;
        this.cloneAppDomainEl.value = domain;
        this.cloneAppNameEl.disabled = false;
        this.cloneAppBranchEl.disabled = false;
        this.cloneAppDomainEl.disabled = false;
        this.cloneAppEl.disabled = false;
    }

    async #cloneApp() {
        const newAppName = this.cloneAppNameEl.value;
        if (!newAppName.length) {
            this.output.setText('please set a name');
            return;
        }

        if (newAppName === this.fetchedApp.spec.name) {
            this.output.setText('please choose a new name');
            return;
        }

        const newAppBranch = this.cloneAppBranchEl.value;
        if (!newAppBranch.length) {
            this.output.setText('please set a branch');
            return;
        }

        const newAppDomain = this.cloneAppDomainEl.value;
        if (newAppDomain && newAppDomain === this.fetchedApp.spec.domains.map(domain => domain.domain)[0]) {
            this.output.setText('please set a new domain or leave it empty');
            return;
        }

        const token = this.doToken.getToken();
        if (token.length < 3) {
            this.output.setText('Error: please set DO Token');
            return;
        }

        const newApp = { spec: JSON.parse(JSON.stringify(this.fetchedApp.spec)) };
        newApp.spec.name = newAppName;
        newApp.spec.services[0].github.branch = newAppBranch;

        if (newAppDomain) {
            this.fetchedApp.spec.domains = [
                {
                    domain: newAppDomain,
                    type: 'PRIMARY',
                    zone: newAppDomain.substr(newAppDomain.indexOf('.')+1)
                }
            ];
        } else {
            delete this.fetchedApp.spec.domains;
        }

        try {
            this.output.setText('request sent, please wait ...');
            const result = await fetch('https://api.digitalocean.com/v2/apps',
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`
                    },
                    body: JSON.stringify(newApp)
                })
                .then(response => {
                    if (!response.ok) {
                        throw new Error('HTTP status ' + response.status);
                    }
                    return response.json();
                });
            this.output.setText(result);
        } catch (e) {
            this.output.setText(e);
        }
    }
}
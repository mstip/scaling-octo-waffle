export default class ListAllDoApps {
    constructor(doToken, output) {
        this.doToken = doToken;
        this.output = output;
        this.el = document.getElementById('listAllDoApps');
        this.el.onclick = event => this.#onClick(event);
    }

    async #onClick(event) {
        event.preventDefault();
        const token = this.doToken.getToken();
        if (token.length < 3) {
            this.output.setText('Error: please set DO Token');
            return;
        }

        try {
            this.output.setText('request sent, please wait ...');
            const result = await fetch('https://api.digitalocean.com/v2/apps',
                {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`
                    }
                })
                .then(response => response.json());
            this.output.setText(result.apps.map(item => ({
                id: item.id,
                name: item.spec.name,
                repo: item.spec.services[0].github.repo,
                branch: item.spec.services[0].github.branch,
                domains: item.spec.domains ? item.spec.domains.map(domain => domain.domain) : []
            })));
        } catch (e) {
            console.log(e);
            this.output.setText(e);
        }
    }
}
import CloneApp from './clone_app.mjs';
import DoToken from './do_token.mjs';
import ListAllDoApps from './list_all_do_apps.mjs';
import Output from './output.mjs';

class CloudGlue {
    constructor() {
        this.output = new Output();
        this.doToken = new DoToken();
        this.listAllDoApps = new ListAllDoApps(this.doToken, this.output);
        this.cloneApp = new CloneApp(this.doToken, this.output);
    }

    start() {
        this.output.setText('ready!');
    }
}

const cG = new CloudGlue();
cG.start();
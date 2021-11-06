import Time from "./time.js";
import view from "./view.js";
import Workshop from "./workshop.js";

export default class Game {

    constructor() {
        this.workshop = new Workshop();
        this.$main = document.querySelector('.main');
        this.$main.onclick = e => this.onClickEvent(e);
        this.time = new Time();
        this.dialog = null;
    }

    run() {
        this.time.inc();
        setTimeout(() => this.run(), 100);
    }

    draw() {
        if (this.dialog === null) {
            view(this.$main, this.workshop, this.time);
        } else {
            dialog(this.$main, this.dialog, this.workshop, this.time);
        }
    }

    onClickEvent(e) {
        console.log(e);
    }
}
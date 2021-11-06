export default class Output {
    constructor() {
        this.el = document.getElementById('output');
    }

    setText(text) {
        if (typeof text === 'object') {
            text = JSON.stringify(text, 0, 2);
        }
        this.el.innerText = text;
    }
}
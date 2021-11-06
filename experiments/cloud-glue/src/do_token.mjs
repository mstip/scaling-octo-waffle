export default class DoToken {
    constructor() {
        this.el = document.getElementById('doToken');
        this.el.value = localStorage.getItem('doToken');
        this.el.onchange = () => this.#onChange();
    }

    getToken() {
        return this.el.value;
    }

    #onChange() {
        localStorage.setItem('doToken', this.el.value);
    }
}
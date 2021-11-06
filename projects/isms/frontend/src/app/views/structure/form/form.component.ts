import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
    selector: 'app-structure-form',
    templateUrl: './form.component.html',
    styleUrls: ['./form.component.css']
})
export class FormComponent implements OnInit {

    @Output() close = new EventEmitter();
    @Input() element: any = {};

    constructor() {
    }

    ngOnInit(): void {
    }

    onClose() {
        this.close.emit();
    }
}

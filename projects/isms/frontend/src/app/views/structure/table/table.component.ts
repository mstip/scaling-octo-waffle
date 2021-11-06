import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
    selector: 'app-structure-table',
    templateUrl: './table.component.html',
    styleUrls: ['./table.component.css']
})
export class TableComponent implements OnInit {

    @Input() assets: any[] = [];
    @Output() onElement = new EventEmitter<number>();

    constructor() {
    }

    ngOnInit(): void {
    }

    onElementClick(id: number) {
        this.onElement.emit(id);
    }
}

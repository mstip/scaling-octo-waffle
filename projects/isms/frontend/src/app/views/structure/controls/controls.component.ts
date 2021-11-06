import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
    selector: 'app-structure-controls',
    templateUrl: './controls.component.html',
    styleUrls: ['./controls.component.css']
})
export class ControlsComponent implements OnInit {

    @Input() isFormVisible = false;
    @Input() isTreeMode = false;

    @Output() onShowTree = new EventEmitter<boolean>();
    @Output() onAdd = new EventEmitter();

    constructor() {
    }

    ngOnInit(): void {
    }

    showTree() {
        this.onShowTree.emit(true);
    }

    onAddClick() {
        this.onAdd.emit();
    }

    showTable() {
        this.onShowTree.emit(false);
    }
}

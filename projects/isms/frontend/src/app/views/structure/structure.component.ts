import { Component, OnInit } from '@angular/core';
import * as faker from 'faker';

@Component({
    selector: 'app-structure',
    templateUrl: './structure.component.html',
    styleUrls: ['./structure.component.css']
})
export class StructureComponent implements OnInit {

    assets: any[] = [];

    isFormVisible = false;
    isTreeMode = false;


    selectedElement: any = {
        number: '',
        name: '',
        location: '',
        type: '',
        title: ''
    };

    constructor() {
    }

    ngOnInit(): void {
        this.assets.push({
            number: 1337,
            name: 'Parkhaus',
            type: 'Haus',
            location: faker.address.cityName(),
        });
        for (let i = 0; i < 8; i++) {
            this.assets.push({
                number: i,
                name: faker.vehicle.vehicle(),
                type: 'Auto',
                location: faker.address.cityName(),
                parent: 'Parkhaus'
            });
        }

        this.assets.push({
            number: 42,
            name: 'Lenkrad',
            type: 'AutoInhalt',
            location: faker.address.cityName(),
            parent: this.assets[this.assets.length - 1].name
        });
    }

    add() {
        this.isFormVisible = true;
        this.selectedElement.number = '';
        this.selectedElement.name = '';
        this.selectedElement.location = '';
        this.selectedElement.type = '';
        this.selectedElement.formTitle = 'Hinzufügen';
    }

    onElementClick(index: number) {
        this.isFormVisible = true;
        this.selectedElement.number = this.assets[index].number;
        this.selectedElement.name = this.assets[index].name;
        this.selectedElement.location = this.assets[index].location;
        this.selectedElement.type = this.assets[index].type;
        this.selectedElement.formTitle = 'Bearbeiten';
    }

    onClose() {
        this.isFormVisible = false;
    }


    showTree(treeMode: boolean) {
        this.isTreeMode = treeMode;
    }
}

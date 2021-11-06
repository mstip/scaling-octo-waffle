import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';

@Component({
    selector: 'app-breadcrumb',
    templateUrl: './breadcrumb.component.html',
    styleUrls: ['./breadcrumb.component.css']
})
export class BreadcrumbComponent implements OnInit {

    breadcrumbs: string[] = [];

    private routeBreadcrumbMap = {
        '/tasks': ['Aufgaben'],
        '/progress': ['Fortschritt'],
        '/structure': ['Strukturbaum'],
        '/settings': ['Einstellungen'],
        '/profile': ['Profil'],
    };


    constructor(private route: ActivatedRoute, private location: Location) {
    }

    ngOnInit(): void {
        this.location.onUrlChange(route => {
            // @ts-ignore
            this.breadcrumbs = this.routeBreadcrumbMap[route];
            if (this.breadcrumbs === undefined) {
                this.breadcrumbs = [];
            }
        });
    }

}

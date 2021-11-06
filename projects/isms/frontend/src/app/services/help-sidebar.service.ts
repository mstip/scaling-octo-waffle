import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class HelpSidebarService {
    get isSidebarVisible(): boolean {
        return this._isSidebarVisible;
    }

    sidebarVisible$: Subject<boolean> = new Subject<boolean>();

    private _isSidebarVisible: boolean = false;

    constructor() {
    }

    toggle() {
        this._isSidebarVisible = !this.isSidebarVisible;
        this.sidebarVisible$.next(this._isSidebarVisible);
    }

}

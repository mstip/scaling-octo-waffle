import { Component, OnInit } from '@angular/core';
import { HelpSidebarService } from '../../services/help-sidebar.service';

@Component({
    selector: 'app-tasks',
    templateUrl: './tasks.component.html',
    styleUrls: ['./tasks.component.css']
})
export class TasksComponent implements OnInit {

    isHelpSidebarVisible: boolean = false;

    constructor(private helpSidebarService: HelpSidebarService) {
    }


    ngOnInit(): void {
        this.isHelpSidebarVisible = this.helpSidebarService.isSidebarVisible;
        this.helpSidebarService.sidebarVisible$.subscribe(visible => this.isHelpSidebarVisible = visible);
    }

}

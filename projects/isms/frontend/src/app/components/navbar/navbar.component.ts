import { Component, OnInit } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { Router } from '@angular/router';
import { HelpSidebarService } from '../../services/help-sidebar.service';

@Component({
    selector: 'app-navbar',
    templateUrl: './navbar.component.html',
    styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
    isAuthenticated: boolean = false;


    constructor(
        private authService: AuthService,
        private router: Router,
        private helpSidebarService: HelpSidebarService
    ) {
    }

    ngOnInit(): void {
        this.isAuthenticated = this.authService.isAuthenticated();
        this.authService.auth$.subscribe(authenticated => this.isAuthenticated = authenticated);
    }

    logout() {
        this.authService.logout();
        this.router.navigate(['/']);
    }

    toggleHelpSidebar() {
        this.helpSidebarService.toggle();
    }
}

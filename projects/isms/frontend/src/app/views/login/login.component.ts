import { Component, OnInit } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { Router } from '@angular/router';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

    userName: string = '';
    password: string = '';
    loginError: boolean = false;

    constructor(private authService: AuthService, private router: Router) {
    }

    ngOnInit(): void {
        if (this.authService.isAuthenticated()) {
            this.router.navigate(['tasks']);
        }
    }

    login() {
        this.authService.login(this.userName, this.password).subscribe(
            () => this.router.navigate(['tasks']),
            () => this.loginError = true
        );
    }
}

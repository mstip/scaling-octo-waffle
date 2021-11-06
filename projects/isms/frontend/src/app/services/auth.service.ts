import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { tap } from 'rxjs/operators';
import { Observable, Subject } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class AuthService {

    auth$: Subject<boolean> = new Subject<boolean>();


    set token(value: string | null) {
        window.localStorage['access_token'] = value;
        this._token = value;
    }

    get token(): string | null {
        return this._token;
    }


    private _token: string | null = null;

    constructor(private httpClient: HttpClient) {
        if (window.localStorage['access_token']) {
            this._token = window.localStorage['access_token'];
        }
    }

    login(username: string, password: string): Observable<any> {
        return this.httpClient.post(environment.apiUrl + 'auth/login', {username, password}).pipe(
            tap(
                (data: any) => {
                    this.token = data.access_token;
                    this.auth$.next(this.isAuthenticated());
                }
            )
        );
    }

    isAuthenticated(): boolean {
        return this._token !== null;
    }

    authenticated(): Observable<boolean> {
        return new Observable<boolean>();
    }

    logout() {
        window.localStorage.removeItem('access_token');
        this._token = null;
        this.auth$.next(this.isAuthenticated());
    }
}

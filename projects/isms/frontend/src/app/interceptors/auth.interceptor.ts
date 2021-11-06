import { Injectable } from '@angular/core';
import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthService } from '../services/auth.service';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {

    constructor(private auth: AuthService) {
    }

    intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
        const accessToken = this.auth.token;

        if (accessToken === null) {
            // TODO: handle
        }

        const authReq = request.clone({
            headers: request.headers.set('Authorization', 'Bearer ' + accessToken)
        });

        return next.handle(authReq);
    }
}

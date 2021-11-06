import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment';

@Component({
    selector: 'app-hello',
    templateUrl: './hello.component.html',
    styleUrls: ['./hello.component.css']
})
export class HelloComponent implements OnInit {

    constructor(private httpClient: HttpClient) {
    }

    ngOnInit(): void {
    }

    getHello() {
        this.httpClient.get(environment.apiUrl + 'hello').subscribe((data: any) => alert(data.data));
    }

}

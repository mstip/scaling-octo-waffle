import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class ShortiesService {

  constructor(private http: HttpClient) { }

  getAll() : Observable<any> {
    return this.http.get('http://localhost:3000/message');
  }
}

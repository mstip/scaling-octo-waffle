import {Component, OnInit} from '@angular/core';
import {ShortiesService} from "../services/shorties.service";

@Component({
  selector: 'app-shorties',
  templateUrl: './shorties.component.html',
  styleUrls: ['./shorties.component.css']
})
export class ShortiesComponent implements OnInit {

  messages: any[] = [];

  constructor(private shortiesService: ShortiesService) {
  }


  ngOnInit(): void {
    this.shortiesService.getAll().subscribe(data => {
        this.messages = data;
        console.log(this.messages);
      },
      error => console.log(error)
    );
  }

}

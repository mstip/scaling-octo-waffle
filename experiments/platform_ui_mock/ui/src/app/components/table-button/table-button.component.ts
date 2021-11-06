import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-table-button',
  templateUrl: './table-button.component.html',
  styleUrls: ['./table-button.component.scss']
})
export class TableButtonComponent implements OnInit {
  @Input() color = 'primary';
  @Input() link: string|null = null;

  constructor() {
  }

  ngOnInit(): void {
  }

}

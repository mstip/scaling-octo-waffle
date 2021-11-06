import { TestBed } from '@angular/core/testing';

import { ShortiesService } from './shorties.service';

describe('ShortiesService', () => {
  let service: ShortiesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ShortiesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

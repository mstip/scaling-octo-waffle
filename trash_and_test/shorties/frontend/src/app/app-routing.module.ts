import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {ShortiesComponent} from "./shorties/shorties.component";

const routes: Routes = [
  {
    path: '',
    component: ShortiesComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}

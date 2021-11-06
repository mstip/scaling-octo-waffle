import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MonitorComponent } from './views/monitor/monitor.component';
import { ClusterDetailComponent } from './views/cluster-detail/cluster-detail.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { BreadcrumbComponent } from './components/breadcrumb/breadcrumb.component';
import { ContentCardComponent } from './components/content-card/content-card.component';
import { TextInputComponent } from './components/text-input/text-input.component';
import { CheckboxComponent } from './components/checkbox/checkbox.component';
import { TextareaComponent } from './components/textarea/textarea.component';
import {FormsModule} from '@angular/forms';
import { ButtonComponent } from './components/button/button.component';
import { MessageComponent } from './components/message/message.component';
import { TableComponent } from './components/table/table.component';
import { ViewComponent } from './components/view/view.component';
import { ClustersComponent } from './views/clusters/clusters.component';
import { ProjectsComponent } from './views/projects/projects.component';
import { SettingsComponent } from './views/settings/settings.component';
import { TableButtonComponent } from './components/table-button/table-button.component';

@NgModule({
  declarations: [
    AppComponent,
    MonitorComponent,
    ClusterDetailComponent,
    NavbarComponent,
    BreadcrumbComponent,
    ContentCardComponent,
    TextInputComponent,
    CheckboxComponent,
    TextareaComponent,
    ButtonComponent,
    MessageComponent,
    TableComponent,
    ViewComponent,
    ClustersComponent,
    ProjectsComponent,
    SettingsComponent,
    TableButtonComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

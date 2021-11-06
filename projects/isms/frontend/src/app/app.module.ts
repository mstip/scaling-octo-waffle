import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { LoginComponent } from './views/login/login.component';
import { HelloComponent } from './views/hello/hello.component';
import { FormsModule } from '@angular/forms';
import { AuthInterceptor } from './interceptors/auth.interceptor';
import { TasksComponent } from './views/tasks/tasks.component';
import { ProgressComponent } from './views/progress/progress.component';
import { StructureComponent } from './views/structure/structure.component';
import { SettingsComponent } from './views/settings/settings.component';
import { ProfileComponent } from './views/profile/profile.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { BreadcrumbComponent } from './components/breadcrumb/breadcrumb.component';
import { ControlsComponent } from './views/structure/controls/controls.component';
import { TableComponent } from './views/structure/table/table.component';
import { TreeComponent } from './views/structure/tree/tree.component';
import { FormComponent } from './views/structure/form/form.component';

@NgModule({
    declarations: [
        AppComponent,
        LoginComponent,
        HelloComponent,
        TasksComponent,
        ProgressComponent,
        StructureComponent,
        SettingsComponent,
        ProfileComponent,
        NavbarComponent,
        BreadcrumbComponent,
        ControlsComponent,
        TableComponent,
        TreeComponent,
        FormComponent
    ],
    imports: [
        BrowserModule,
        HttpClientModule,
        AppRoutingModule,
        FormsModule
    ],
    providers: [[
        {provide: HTTP_INTERCEPTORS, useClass: AuthInterceptor, multi: true},
    ]],
    bootstrap: [AppComponent]
})
export class AppModule {
}

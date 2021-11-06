import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './views/login/login.component';
import { HelloComponent } from './views/hello/hello.component';
import { AuthGuard } from './guards/auth.guard';
import { TasksComponent } from './views/tasks/tasks.component';
import { ProgressComponent } from './views/progress/progress.component';
import { StructureComponent } from './views/structure/structure.component';
import { SettingsComponent } from './views/settings/settings.component';
import { ProfileComponent } from './views/profile/profile.component';

const routes: Routes = [
    {path: '', component: LoginComponent},
    {path: 'hello', component: HelloComponent, canActivate: [AuthGuard]},
    {path: 'tasks', component: TasksComponent, canActivate: [AuthGuard]},
    {path: 'progress', component: ProgressComponent, canActivate: [AuthGuard]},
    {path: 'structure', component: StructureComponent, canActivate: [AuthGuard]},
    {path: 'settings', component: SettingsComponent, canActivate: [AuthGuard]},
    {path: 'profile', component: ProfileComponent, canActivate: [AuthGuard]},
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {
}

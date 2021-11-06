import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {MonitorComponent} from './views/monitor/monitor.component';
import {ClusterDetailComponent} from './views/cluster-detail/cluster-detail.component';
import {ProjectsComponent} from './views/projects/projects.component';
import {ClustersComponent} from './views/clusters/clusters.component';
import {SettingsComponent} from './views/settings/settings.component';

const routes: Routes = [
  {path: '', component: MonitorComponent},
  {path: 'monitor', component: MonitorComponent},
  {path: 'projects', component: ProjectsComponent},
  {path: 'clusters', component: ClustersComponent},
  {path: 'clusters/:id', component: ClusterDetailComponent},
  {path: 'settings', component: SettingsComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}

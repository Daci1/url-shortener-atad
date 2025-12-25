import { Routes } from '@angular/router';
import {SigninPageComponent} from './components/sign-in/sign-in.component';
import {HomePageComponent} from './components/home-page/home-page.component';
import {SignUpComponent} from './components/sign-up/sign-up.component';

export const routes: Routes = [
  // ...other routes
  {path: '', component: HomePageComponent, pathMatch: 'full'},
  { path: 'signin', component: SigninPageComponent },
  { path: 'signup', component: SignUpComponent },
];

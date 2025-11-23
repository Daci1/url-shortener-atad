import { Component } from '@angular/core';
import { HomePageComponent } from './components/home-page/home-page.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [HomePageComponent],
  template: '<app-home-page></app-home-page>',
})
export class AppComponent {}

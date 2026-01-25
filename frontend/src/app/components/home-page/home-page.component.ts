import {Component, Input, OnDestroy, OnInit} from '@angular/core';
import { NavbarComponent } from "../navbar/navbar.component";
import { FormsModule } from '@angular/forms';
import {CommonModule} from '@angular/common';
import {AuthService} from '../../services/auth.service';
import {UrlService} from '../../services/url.service';
import {Subscription} from 'rxjs';
import {environment} from '../../../environments/environment';
import {QRCodeComponent} from 'angularx-qrcode';

@Component({
  selector: 'app-home-page',
  imports: [CommonModule, NavbarComponent, FormsModule, QRCodeComponent],
  standalone: true,
  templateUrl: './home-page.component.html',
  styleUrl: './home-page.component.css'
})
export class HomePageComponent implements OnInit, OnDestroy{

  isLogged: boolean = false;
  urlToShorten: string = '';
  customUrl: string = '';
  shortenedUrl: string | null = null;
  token: string | null = null;
  loggedUserSubscription: Subscription | undefined;
  endpointPrefix: string = environment.url + 'api/v1/urls/'

  @Input() showUrlInput: boolean = false;

  onGetStarted(): void {
    this.showUrlInput = true;
  }

  constructor(private authService: AuthService, private urlService: UrlService) { }

  ngOnInit() {
    this.loggedUserSubscription = this.authService.loggedUserData.subscribe(userData => {
      this.isLogged = !!userData;
      this.token = userData ? userData.token : null;
    });
  }

  ngOnDestroy(): void {
    this.loggedUserSubscription?.unsubscribe();
  }

  shortenUrl(): void {
    const trimmedUrl = this.urlToShorten.trim();
    if (!trimmedUrl) {
      alert('Please enter a valid URL.');
      return;
    }

    if(this.isLogged) {
      const sub = this.authService.getLoggedSub();

      const customUrlTrimmed = this.customUrl.trim();
      if(customUrlTrimmed) {
        this.urlService.loggedCreateCustomShortUrl(this.urlToShorten, sub, customUrlTrimmed, this.token!).subscribe({
          next: (res) => {
            this.shortenedUrl = res.data.attributes.shortUrl;
          },
          error: (err) => {
            console.error('Error creating custom short URL:', err);
          }
        })
      } else {
        this.urlService.loggedCreateShortUrl(this.urlToShorten, sub, this.token!).subscribe({
          next: (res) => {
            this.shortenedUrl = res.data.attributes.shortUrl;
          }
        })
      }
    } else {
      this.urlService.createShortUrl(trimmedUrl).subscribe({
        next: (res) => {
          this.shortenedUrl = res.data.attributes.shortUrl;
          console.log(this.shortenedUrl);
        }
      });
    }
  }

  onNavigateToShortUrl() {
    window.open(this.endpointPrefix + this.shortenedUrl, '_blank', 'noopener,noreferrer')
  }
}

import { Component, Input } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { NavbarComponent } from "../navbar/navbar.component";
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-home-page',
  imports: [NavbarComponent, FormsModule],
  templateUrl: './home-page.component.html',
  styleUrl: './home-page.component.css'
})
export class HomePageComponent {
  urlToShorten: string = ''; // To store the URL input by the user
  shortenedUrl: string | null = null;
  @Input() showUrlInput: boolean = false; // To toggle the visibility of the Get Started flow

  onGetStarted(): void {
    this.showUrlInput = true; // Make the URL input section visible
  }

  constructor(private http: HttpClient) { }

  shortenUrl(): void {
    if (!this.urlToShorten.trim()) {
      alert('Please enter a valid URL.');
      return;
    }

    // Example hardcoded API call
    this.http.post<{ shortUrl: string }>('https://api.example.com/shorten', { originalUrl: this.urlToShorten })
      .subscribe({
        next: (response) => {
          this.shortenedUrl = response.shortUrl;
          alert(`Shortened URL: ${this.shortenedUrl}`);
        },
        error: (err) => {
          console.error('Error while shortening URL:', err);
          alert('Failed to shorten the URL. Please try again.');
        }
      });
  }
}

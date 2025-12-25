import {Component} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {LoginRequest} from '../../models/user.models';
import {REFRESH_TOKEN_KEY, TOKEN_KEY, USERNAME_KEY} from '../../constants';
import {AuthService} from '../../services/auth.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-signin-page',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './sign-in.component.html',
  styleUrl: './sign-in.component.css'
})
export class SigninPageComponent {
  email: string = '';
  password: string = '';

  constructor(private authService: AuthService, private router: Router) {
  }

  onSignIn(): void {
    const body: LoginRequest = {
      data: {
        type: 'users',
        attributes: {
          email: this.email,
          password: this.password
        }
      }
    };

    this.authService.logIn(body)
      .subscribe({
        next: (res) => {
          const token = res.data.attributes.token;
          const refreshToken = res.data.attributes.refreshToken;
          const username = res.data.attributes.username;
          localStorage.setItem(TOKEN_KEY, token);
          localStorage.setItem(REFRESH_TOKEN_KEY, refreshToken);
          localStorage.setItem(USERNAME_KEY, username);
          this.authService.loggedUserData.next({token, refreshToken, username});
          this.router.navigate(['/'])
        },
        error: (err) => {
          alert('Error during sign in');
        }
      });
  }
}

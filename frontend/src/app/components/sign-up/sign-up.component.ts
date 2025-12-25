import { Component } from '@angular/core';
import {LoginRequest, RegisterRequest} from '../../models/user.models';
import {AuthService} from '../../services/auth.service';
import {Router} from '@angular/router';
import {FormsModule} from '@angular/forms';

@Component({
  selector: 'app-sign-up',
  imports: [FormsModule],
  templateUrl: './sign-up.component.html',
  styleUrl: './sign-up.component.css'
})
export class SignUpComponent {
  email: string = '';
  username: string = '';
  password: string = '';

  constructor(private authService: AuthService, private router: Router) {

  }


  onRegister(): void {
    const body: RegisterRequest = {
      data: {
        type: 'users',
        attributes: {
          email: this.email,
          username: this.username,
          password: this.password
        }
      }
    };

    this.authService.register(body)
      .subscribe({
        next: () => {
          this.router.navigate(['/signin'])
        },
        error: (err) => {
          console.log(err);
          alert('Error during sign in');
        }
      });
  }
}

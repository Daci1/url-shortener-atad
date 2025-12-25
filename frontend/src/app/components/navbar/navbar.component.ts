import {Component, OnInit} from '@angular/core';
import {Router, RouterLink} from '@angular/router';
import {Subscription} from 'rxjs';
import {AuthService} from '../../services/auth.service';

@Component({
  selector: 'app-navbar',
  imports: [RouterLink],
  standalone: true,
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.css'
})
export class NavbarComponent implements OnInit {
  isLogged: boolean = false;
  username: string | null = null;
  loggedUserSubscription: Subscription | undefined;

  constructor(private authService: AuthService, private router: Router) {

  }

  ngOnInit() {
    this.loggedUserSubscription = this.authService.loggedUserData.subscribe(userData => {
      this.isLogged = !!userData;
      this.username = userData ? userData.username : null;
    });
  }
  onLogOut() {
    this.authService.logOut();
    this.router.navigate(['/']);
  }

  ngOnDestroy(): void {
    this.loggedUserSubscription?.unsubscribe();
  }
}

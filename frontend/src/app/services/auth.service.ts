import { Injectable } from '@angular/core';
import {BehaviorSubject} from 'rxjs';
import {LoginOrRegisterResponse, LoginRequest, RegisterRequest, UserData} from '../models/user.models';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../environments/environment';
import {REFRESH_TOKEN_KEY, TOKEN_KEY, USERNAME_KEY} from '../constants';
import {jwtDecode} from 'jwt-decode';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  loggedUserData = new BehaviorSubject<UserData | null>(null);

  constructor(private http: HttpClient) {}

  register(registerData: RegisterRequest) {
    return this.http.post<LoginOrRegisterResponse>(environment.url + 'api/v1/users', registerData);
  }
  logIn(loginData: LoginRequest) {
    return this.http.post<LoginOrRegisterResponse>(environment.url + 'api/v1/users/login', loginData);
  }

  logOut() {
    this.clearUserFromLocalStorage();
    this.loggedUserData.next(null);
  }

  clearUserFromLocalStorage() {
    localStorage.removeItem(TOKEN_KEY);
    localStorage.removeItem(REFRESH_TOKEN_KEY);
    localStorage.removeItem(USERNAME_KEY);
  }

  loadUserFromLocalStorage() {
    try {
      const loggedUserData: UserData = {
        token: this.getFromLocalStorageOrThrow(TOKEN_KEY),
        refreshToken: this.getFromLocalStorageOrThrow(REFRESH_TOKEN_KEY),
        username: this.getFromLocalStorageOrThrow(USERNAME_KEY),
      }
      return loggedUserData;
    } catch (e: any) {
      return null;
    }
  }
  private getFromLocalStorageOrThrow(key: string): string {
    const value = localStorage.getItem(key);
    if (!value) {
      throw new Error(`Missing ${key} in local storage`);
    }
    return value;
  }

  getLoggedSub() {
    const userData = this.loggedUserData.getValue();
    if (userData?.token) {
      return jwtDecode<{ sub: string }>(userData.token).sub;
    }

    return '';
  }

  autoLogin() {
    const loggedUserData = this.loadUserFromLocalStorage();
    if (loggedUserData) {
      this.loggedUserData.next(loggedUserData);
    }
  }
}

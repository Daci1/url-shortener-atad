import {Injectable} from '@angular/core';
import {CreateUrlRequest, CreateUrlResponse} from '../models/url.models';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../environments/environment';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UrlService {

  constructor(private http: HttpClient) {
  }

  createShortUrl(originalUrl: string): Observable<CreateUrlResponse> {
    return this.http.post<CreateUrlResponse>(environment.url + 'api/v1/urls', {
      data: {
        type: 'urls',
        attributes: {originalUrl}
      }
    } as CreateUrlRequest)

  }

  loggedCreateShortUrl(originalUrl: string, username: string, token: string): Observable<CreateUrlResponse> {
    return this.http.post<CreateUrlResponse>(environment.url + `api/v1/urls/users/${username}`, {
      data: {
        type: 'urls',
        attributes: {originalUrl}
      }
    } as CreateUrlRequest,
      { headers: {Authorization: `Bearer ${token}`}})
  }
}

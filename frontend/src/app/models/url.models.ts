export type CreateUrlRequest = {
  data: {
    type: 'urls';
    attributes: {
      originalUrl: string;
      username?: string;
    }
  }
}

export type CreateUrlResponse = {
  data: {
    type: 'urls';
    attributes: {
      "shortUrl": string;
      "originalUrl": string;
      "createdAt": string;
    }
  }
}

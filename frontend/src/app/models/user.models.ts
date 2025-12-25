export type RegisterRequest = {
  data: {
    type: 'users';
    attributes: {
      email: string;
      username: string;
      password: string;
    }
  }
}

export type LoginRequest = {
  data: {
    type: 'users';
    attributes: {
      email: string;
      password: string;
    }
  }
}
export type LoginOrRegisterResponse = {
  data: {
    attributes: UserData;
    type: string;
  }
};

export type UserData = {
  username: string;
  token: string;
  refreshToken: string;
}

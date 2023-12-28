import HTTP from './http';

const session_apis = {
  async login(email, password) {
    const response = await HTTP.post(`signin`, {
      email: email,
      password: password
    });
    console.log(response.data)
  }
}

export default session_apis;

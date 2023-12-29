import HTTP from './http';
import router from '../router/index'

const session_apis = {
  async login(email, password) {
    const response = await HTTP.post(`signin`, {
      email: email,
      password: password
    }).then(function (response) {
      localStorage.setItem('token', response.data.token)
      router.push({ path: '/' })
    }).catch(function (error) {
      console.log(error)
    });
  }
}

export default session_apis;

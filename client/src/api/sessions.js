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
      // TODO: Do error handler later
      console.log(error)
    });
  },

  async signup(email, password) {
    const response = await HTTP.post(`signup`, {
      email: email,
      password: password,
    }).then(function (response) {
      router.push({path: '/login'})
    }).catch(function(error) {
      // TODO: Do error handler later
      console.log(error)
    })
  }
}

export default session_apis;

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
      if(error.response.data.needVerify) {
        router.push({ name: 'verify', query: { email: email } })
      }
    });
  },

  async signup(email, name, password) {
    const response = await HTTP.post(`signup`, {
      name: name,
      email: email,
      password: password,
    }).then(function (response) {
      router.push({path: '/login'})
    }).catch(function(error) {
      // TODO: Do error handler later
      console.log(error)
    })
  },

  async logout() {
    const token = localStorage.getItem('token')
    const response = await HTTP.delete(`logout`, {
      headers: {
        Token: token
      }
    }).catch(function(error) {
      // TODO: Do error handler later
      console.log(error)
    }).finally(function(error) {
      localStorage.removeItem('token')
      router.push({path: '/login'})
    })
  },

  async verify(email, verifyCode) {
    const response = await HTTP.post(`/verify/activate`, {
      email: email,
      verifyCode: String(verifyCode),
    }).then(function (response) {
      router.push({path: '/login'})
    }).catch(function(error) {
      // TODO: Do error handler later
      console.log(error)
    })
  },

  async sendCode(email) {
    const response = await HTTP.post(`/verify/send`, {
      email: email,
    }).then(function (response) {
      // TODO: Display some message here
    }).catch(function(error) {
      // TODO: Do error handler later
      console.log(error)
    })
  },
}

export default session_apis;

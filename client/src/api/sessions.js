import HTTP from '@/utils/http';
import { notifyError, notifySuccess } from './notification';
import router from '../router/index';

const session_apis = {
  async login(email, password) {
    const response = await HTTP.post(`signin`, {
      email: email,
      password: password
    }).then(function (response) {
      localStorage.setItem('token', response.data.token)
      if(response.data.message) notifySuccess(response.data.message)
      router.push({ path: '/' })
    }).catch(function (error) {
      if(error.response.data.needVerify) {
        router.push({ name: 'verify', query: { email: email } })
      } else if(error.response.data.message) {
        notifyError(error.response.data.message)
      }
    });
  },

  async signup(email, name, password) {
    const response = await HTTP.post(`signup`, {
      name: name,
      email: email,
      password: password,
    }).then(function (response) {
      if(response.data.message) notifySuccess(response.data.message)
      router.push({path: '/login'})
    }).catch(function(error) {
      if(error.response.data.message) notifyError(error.response.data.message)
    })
  },

  async logout() {
    const response = await HTTP.delete(`logout`).catch(function(error) {
      // if(error.response.data.message) notifyError(error.response.data.message)
    }).finally(function() {
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
      if(error.response.data.message) notifyError(error.response.data.message)
    })
  },

  async sendCode(email) {
    const response = await HTTP.post(`/verify/send`, {
      email: email,
    }).then(function (response) {
      if(response.data.message) notifySuccess(response.data.message)
    }).catch(function(error) {
      if(error.response.data.message) notifyError(error.response.data.message)
    })
  },
}

export default session_apis;

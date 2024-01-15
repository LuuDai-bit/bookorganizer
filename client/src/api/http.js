import axios from 'axios';
import session_apis from './sessions';

const HTTP = axios.create({
  baseURL: import.meta.env.VITE_BASE_URL,
  headers: {
    Token: {
      toString() {
        return localStorage.getItem('token')
      }
    }
  }
})

HTTP.interceptors.request.use(function(config) {
  let loading = document.getElementById('loading')
  if(loading && config.method != 'get') {
    loading.classList.remove('hidden')
  }

  return config
},function(error) {
  let loading = document.getElementById('loading')
  if(loading && !loading.classList.contains('hidden')) {
    loading.classList.add('hidden')
  }

  return Promise.reject(error);
})

HTTP.interceptors.response.use(function(response) {
  let loading = document.getElementById('loading')
  if(loading && !loading.classList.contains('hidden')) {
    loading.classList.add('hidden')
  }

  return response
},function(error) {
  let loading = document.getElementById('loading')
  if(loading && !loading.classList.contains('hidden')) {
    loading.classList.add('hidden')
  }

  if(error.response.status === 401) {
    session_apis.logout()
  }

  return Promise.reject(error);
})

const HTTPWithoutInterceptor = axios.create({
  baseURL: import.meta.env.VITE_BASE_URL,
  headers: {
    Token: {
      toString() {
        return localStorage.getItem('token')
      }
    }
  }
})

export default HTTP;
export { HTTPWithoutInterceptor }

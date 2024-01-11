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

HTTP.interceptors.response.use(function(response) {
  return response
},function(error) {
  if(error.response.status === 401) {
    session_apis.logout()
  }

  throw error;
})

export default HTTP;

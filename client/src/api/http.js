import axios from 'axios';

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

export default HTTP;

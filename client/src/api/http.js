import axios from 'axios';

const HTTP = axios.create({
  baseURL: `http://localhost:3000/api/v1`,
  headers: {
    Token: {
      toString() {
        return localStorage.getItem('token')
      }
    }
  }
})

export default HTTP;

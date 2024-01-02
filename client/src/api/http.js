import axios from 'axios';

const token = localStorage.getItem('token') || ''
const HTTP = axios.create({
  baseURL: `http://localhost:3000/api/v1`,
  headers: {
    Token: token
  }
})

export default HTTP;

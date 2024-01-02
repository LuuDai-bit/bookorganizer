import HTTP from './http';
import { notifyError, notifySuccess } from './notification';
import router from '../router/index';

const bookApis = {
  async getBooks(page) {
    const response = await HTTP.get(`/books/${page || 1}`)

    return response
  },

  async createBook(name, author, purchaseDate, startReadAt, finishReadAt, categories) {
    const response = await HTTP.post(`/books/create`, {
      name: name,
      author: author,
      purchaseDate: purchaseDate,
      startReadAt: startReadAt,
      finishReadAt: finishReadAt,
      categories: categories,
    })

    return response
  },
}

export default bookApis

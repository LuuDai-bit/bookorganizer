import HTTP from './http';
import { formatDateWithFormat } from '@/mixins/format_date';

const bookApis = {
  async getBooks(page) {
    const response = await HTTP.get(`/books/${page || 1}`)

    return response
  },

  async createBook(name, author, purchaseDate, startReadAt, finishReadAt, categories) {
    const response = await HTTP.post(`/books/create`, {
      name: name,
      author: author,
      purchaseDate: formatDateWithFormat(purchaseDate, 'YYYY-MM-DD'),
      startReadAt: formatDateWithFormat(startReadAt, 'YYYY-MM-DD hh:mm:ss'),
      finishReadAt: formatDateWithFormat(finishReadAt, 'YYYY-MM-DD hh:mm:ss'),
      categories: categories,
    })

    return response
  },

  async updateBook(id, name, author, purchaseDate, startReadAt, finishReadAt, categories) {
    const response = await HTTP.patch(`/books/update`, {
      id: id,
      name: name,
      author: author,
      purchaseDate: formatDateWithFormat(purchaseDate, 'YYYY-MM-DD'),
      startReadAt: formatDateWithFormat(startReadAt, 'YYYY-MM-DD hh:mm:ss'),
      finishReadAt: formatDateWithFormat(finishReadAt, 'YYYY-MM-DD hh:mm:ss'),
      categories: categories,
    })

    return response
  },
}

export default bookApis

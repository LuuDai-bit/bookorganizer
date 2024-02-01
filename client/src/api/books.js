import HTTP from '@/utils/http';
import { formatDateWithFormat } from '@/mixins/format_date';

const bookApis = {
  async getBooks(page, search="") {
    const response = await HTTP.get(`/books/${page || 1}`, {
      params: { s: search }
    })

    return response
  },

  async createBook(book) {
    const response = await HTTP.post(`/books/create`, {
      name: book.name,
      author: book.author,
      purchaseDate: formatDateWithFormat(book.purchaseDate, 'YYYY-MM-DD'),
      startReadAt: formatDateWithFormat(book.startReadAt, 'YYYY-MM-DD hh:mm:ss'),
      finishReadAt: formatDateWithFormat(book.finishReadAt, 'YYYY-MM-DD hh:mm:ss'),
      key: book.key,
      fileName: book.fileName,
      type: book.type,
      categories: book.categories,
    })

    return response
  },

  async updateBook(book) {
    const response = await HTTP.patch(`/books/update`, {
      id: book.id,
      name: book.name,
      author: book.author,
      purchaseDate: formatDateWithFormat(book.purchaseDate, 'YYYY-MM-DD'),
      startReadAt: formatDateWithFormat(book.startReadAt, 'YYYY-MM-DD hh:mm:ss'),
      finishReadAt: formatDateWithFormat(book.finishReadAt, 'YYYY-MM-DD hh:mm:ss'),
      key: book.key,
      fileName: book.fileName,
      type: book.type,
      categories: book.categories,
    })

    return response
  },

  async download(id) {
    const response = await HTTP.get(`/books/download`, {
      params: { id: id }
    })

    return response
  },
}

export default bookApis

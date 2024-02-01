import { HTTPWithoutInterceptor } from '@/utils/http';

const statisticApis = {
  async countBookByYear(year) {
    const response = await HTTPWithoutInterceptor.get(`/statistic/books/read/${year}`)

    return response
  },
  async getFavoriteCategories() {
    const response = await HTTPWithoutInterceptor.get(`/statistic/categories/favorite`)

    return response
  },
}

export default statisticApis;

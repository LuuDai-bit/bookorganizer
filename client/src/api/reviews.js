import HTTP from '@/utils/http';

const reviewApis = {
  async createReview(bookId, point, review) {
    const response = await HTTP.post(`/reviews/create`, {
      bookId: bookId,
      point: point,
      comment: review,
    })

    return response
  },
}

export default reviewApis;

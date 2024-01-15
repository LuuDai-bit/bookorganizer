import { HTTPWithoutInterceptor } from './http';

const headers = { 'Content-Type': 'multipart/form-data' }
const fileApis = {
  async uploadSingleFile(file) {
    const response = await HTTPWithoutInterceptor.post(`/file/single`, {
      file: file,
    }, {
      headers
    })

    return response
  },
}

export default fileApis;

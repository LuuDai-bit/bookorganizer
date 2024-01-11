import HTTP from './http';
import { notifyError } from './notification';
import router from '../router/index';

const userApis = {
  async getMyInfo() {
    const response = await HTTP.get(`/users/me`)

    return response
  },

  async changePassword(id, password, newPassword) {
    const response = await HTTP.patch(`/users/change_password`, {
      id: id,
      password: password,
      newPassword: newPassword,
    }).then(function(response) {
      localStorage.removeItem('token')
      router.push({path: '/login'})
    }).catch(function(error) {
      if(error.response.data.message) notifyError(error.response.data.message)
    })

    return response
  },
}

export default userApis

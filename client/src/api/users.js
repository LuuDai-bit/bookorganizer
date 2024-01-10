import HTTP from './http';
import session_apis from './sessions';
import { notifyError } from './notification';


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
      session_apis.logout();
    }).catch(function(error) {
      if(error.response.data.message) notifyError(error.response.data.message)
    })

    return response
  },
}

export default userApis

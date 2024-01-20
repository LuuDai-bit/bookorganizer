<template>
  <div id="setting" class="">
    <Breadcrumb paths="Settings" />

    <div class="lg:flex justify-between items-center mb-6">
      <p class="text-2xl font-semibold mb-2 lg:mb-0">Do it your style!</p>
      <FunctionButton text="Logout" @clickHandler="out()"/>
    </div>

    <br>

    <Avatar :url="user.avatar_url"/>

    <div class="inline-block max-w-[60%]">
      <div class="w-48 border-solid border-2 border-grey-900 text-center mr-1 mb-2 inline-block">
        Name
      </div>
      <div class="w-96 border-solid border-2 border-grey-900 text-center mb-2 inline-block">
        {{ user.name || "No name" }}
      </div>
      <div class="w-48 border-solid border-2 border-grey-900 text-center mr-1 mb-2 inline-block">
        Email
      </div>
      <div class="w-96 border-solid border-2 border-grey-900 text-center mb-2 inline-block">
        {{ user.email || "No email" }}
      </div>
      <div class="w-48 border-solid border-2 border-grey-900 text-center mr-1 mb-2 inline-block">
        Register date
      </div>
      <div class="w-96 border-solid border-2 border-grey-900 text-center mb-2 inline-block">
        {{ user.createdAt || "-" }}
      </div>
    </div>

    <div class="mt-10">
      <button class="inline-block hover:text-sky-600 mr-4"
              @click="openChangePasswordModal()">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 inline">
          <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 5.25a3 3 0 0 1 3 3m3 0a6 6 0 0 1-7.029 5.912c-.563-.097-1.159.026-1.563.43L10.5 17.25H8.25v2.25H6v2.25H2.25v-2.818c0-.597.237-1.17.659-1.591l6.499-6.499c.404-.404.527-1 .43-1.563A6 6 0 1 1 21.75 8.25Z" />
        </svg>
        <span>Change password</span>
      </button>
      <button class="inline-block hover:text-sky-600"
              @click="openChangeAvatarModal()">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 inline">
          <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 5.25a3 3 0 0 1 3 3m3 0a6 6 0 0 1-7.029 5.912c-.563-.097-1.159.026-1.563.43L10.5 17.25H8.25v2.25H6v2.25H2.25v-2.818c0-.597.237-1.17.659-1.591l6.499-6.499c.404-.404.527-1 .43-1.563A6 6 0 1 1 21.75 8.25Z" />
        </svg>
        <span>Change avatar</span>
      </button>
    </div>
  </div>
  <ChangePassWordForm v-if="showModal"
                      :id="user.id"
                      @closeChangePasswordModal="closeChangePasswordModal()" />
  <ChangeAvatarForm v-if="changeAvatar"
                    @closeChangeAvatarModal="closeChangeAvatarModal()"
                    @fetchUserDetail="reload()" />
</template>

<script>
import Breadcrumb from '@/components/common/Breadcrumb.vue';
import FunctionButton from '@/components/common/FunctionButton.vue';
import ChangePassWordForm from '@/components/settings/ChangePasswordForm.vue';
import userApis from '@/api/users';
import session_apis from '@/api/sessions';
import Avatar from '@/components/settings/Avatar.vue';
import ChangeAvatarForm from '@/components/settings/ChangeAvatarForm.vue';

export default {
  name: "SettingView",
  components: {
    Breadcrumb,
    FunctionButton,
    ChangePassWordForm,
    Avatar,
    ChangeAvatarForm,
},
  data() {
    return {
      user: {},
      showModal: false,
      changeAvatar: false,
    }
  },
  created() {
    this.fetchUserDetail()
  },
  methods: {
    out() {
      session_apis.logout()
    },
    fetchUserDetail() {
      let self = this
      userApis.getMyInfo().then(function(response) {
        self.user = response.data.user
      }).catch(function(error) {
        if(error.response.data.message) {
          notifyError(error.response.data.message)
        } else {
          notifyError('Something went wrong')
        }
      })
    },
    openChangePasswordModal() {
      this.showModal = true
    },
    closeChangePasswordModal() {
      this.showModal = false
    },
    openChangeAvatarModal() {
      this.changeAvatar = true
    },
    closeChangeAvatarModal() {
      this.changeAvatar = false
    },
    reload() {
      window.location.reload()
    }
  }
}
</script>

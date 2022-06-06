<template>
  <div class="container">

    <div class="list">
      <UsersListComponent :users="data" @selected="onSelect"/>
    </div>

    <div class="dialog">
      <DialogComponent :current-user="currentUser" :messages="messages" :selected-user="selectedUsername" @sent="sent"/>
    </div>

    <div class="errorSpace">
      <transition name="slide-fade">
        <ErrorAlertComponent v-if="showError" :err-message="errorMessage" @close="closeError" class="error"/>
      </transition>
    </div>

  </div>
</template>

<script>
import ErrorAlertComponent from "@/components/utils/ErrorAlertComponent";
import DialogComponent from "@/components/DialogComponent";
import jwtDecode from "jwt-decode";
import fetchUsers from "@/api/fetchUsers";
import fetchMessages from "@/api/fetchMessages";
import UsersListComponent from "@/components/UsersListComponent";

export default {
  components:{
    UsersListComponent,
    DialogComponent,
    ErrorAlertComponent
  },
  name: "UsersView",
  data(){
    return{
      token: localStorage.getItem("JWTToken"),
      data: [],
      showError: false,
      errorMessage: '',
      currentUser: '',
      selectedUsername: '',
      messages:[]
    }
  },

  created(){
    this.token = localStorage.getItem("JWTToken")
    this.currentUser = jwtDecode(this.token).Username
    this.loadUsers()
  },

  methods:{

    sent(){
      this.loadMessages()
    },

    closeError(){
      this.showError=false
      this.errorMessage = ''
    },

    onSelect(username){
      console.log(`log from users ${username}`)
      this.selectedUsername = username
      this.loadMessages()
    },

    loadMessages(){

      if (!this.selectedUsername){
        return null
      }
      fetchMessages(this.selectedUsername, this.token, this)
      console.log("messages request")
    },

    loadUsers() {
      fetchUsers(this.token, this)
    },

    displayError(){
      this.showError=true
    }

  }
}
</script>

<style scoped src="@/assets/css/usersView.css" />
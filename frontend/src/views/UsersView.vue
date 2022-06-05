<template>
  <div class="container">

    <div class="list">
      <div class="left">
        <div class="elem" v-for="(item, i) in data" :key="i">
          <ListItemComponent :item="item" @selected="onSelect"/>
        </div>
      </div>
    </div>

    <div class="dialog">
      <DialogComponent :current-username="currentUsername" :messages="messages" :selected-user="selectedUsername" @sent="sent"/>
    </div>

    <div class="errorSpace">
      <transition name="slide-fade">
        <ErrorAlertComponent v-if="showError" :err-message="errorMessage" @close="closeError" class="error"/>
      </transition>
    </div>

  </div>
</template>

<script>
import ListItemComponent from "@/components/ListItemComponent";
import ErrorAlertComponent from "@/components/utils/ErrorAlertComponent";
import DialogComponent from "@/components/DialogComponent";
import jwtDecode from "jwt-decode";
import fetchUsers from "@/api/fetchUsers";
import fetchMessages from "@/api/fetchMessages";

export default {
  components:{
    DialogComponent,
    ListItemComponent,
    ErrorAlertComponent
  },
  name: "UsersView",
  data(){
    return{
      token: localStorage.getItem("JWTToken"),
      data: [],
      showError: false,
      errorMessage: '',
      currentUsername: '',
      selectedUsername: '',
      messages:[]
    }
  },

  created(){
    this.token = localStorage.getItem("JWTToken")
    this.currentUsername = jwtDecode(this.token).Username
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
      console.log(`selected ${username}`)
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
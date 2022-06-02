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
      <DialogComponent :current-username="currentUsername" :messages="messages" :selected-user="selectedUsername"/>
    </div>
    <div class="errorSpace">
      <transition name="slide-fade">
        <ErrorAlertComponent v-if="showError" :err-message="errorMessage" @close="closeError" class="error"/>
      </transition>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import ListItemComponent from "@/components/ListItemComponent";
import ErrorAlertComponent from "@/components/ErrorAlertComponent";
import DialogComponent from "@/components/DialogComponent";
import jwtDecode from "jwt-decode";

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
      // messages:[
      //   {
      //     from: "user1",
      //     message: "hello world",
      //     date: 1652535907
      //   },
      //   {
      //     from: "default companion",
      //     message: "hello",
      //     date: 1652535909
      //   }
      // ],
    }
  },

  created(){
    this.token = localStorage.getItem("JWTToken")
    this.currentUsername = jwtDecode(this.token).Username
    this.loadUsers()
  },

  methods:{

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

      console.log("messages request")
      axios.post("http://localhost:8080/messages/" + this.selectedUsername,
          {},
          {
            headers: {
              'Authorization': `Bearer ${this.token}`
            }
          }
      ).then(res => {
        if (res.status === 200){

          this.messages = res.data.messages

        }else {
          this.errorMessage = res.data.message
          this.displayError()
        }
      }).catch(err => {
        if (err.response){
          this.errorMessage = "failed to fetch data"
          console.log(err.response.data.error)
          this.displayError()
        }
      })
    },

    loadUsers() {
      axios.post("http://localhost:8080/messages/all",
          {},
          {
            headers: {
              'Authorization': `Bearer ${this.token}`
            }
          }
      ).then(res => {
        if (res.status === 200){
          this.data = res.data.users
        }else {
          this.errorMessage = res.data.message
          this.displayError()
        }
      }).catch(err => {
        if (err.response){
          this.errorMessage = "failed to fetch data"
          console.log(err.response.data.error)
          this.displayError()
        }
      })
    },

    displayError(){
      this.showError=true
    }

  }
}
</script>

<style scoped>
.container{
  display: flex;
  height: 100%;
  width: 100%;
}
.left {
  float: left;
}
.elem{
  margin: 10px;
}

.list{
  position: relative;
  height: 800px;
  overflow:hidden; overflow-y:scroll;
}
.dialog{
  float: right;
  width: 100%;
  height: 800px;
}

/* width */
::-webkit-scrollbar {
  width: 10px;
}

/* Track */
::-webkit-scrollbar-track {
  background: #f1f1f1;
}

/* Handle */
::-webkit-scrollbar-thumb {
  background: #888;
}

/* Handle on hover */
::-webkit-scrollbar-thumb:hover {
  background: #555;
}
</style>
<template>
  <div class="container">
    <div class="list">
      <div class="left">
        <div class="elem" v-for="(item, i) in data" :key="item.id">
          <ListItemComponent :item="item" :key="i" @selected="onSelect"/>
        </div>
      </div>
      <div class="errorSpace">
        <transition name="slide-fade">
          <ErrorAlertComponent v-if="showError" :err-message="errorMessage" @close="closeError" class="error"/>
        </transition>
      </div>
    </div>
    <div class="dialog">
      <DialogComponent :to="currentCompanion"/>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import ListItemComponent from "@/components/ListItemComponent";
import ErrorAlertComponent from "@/components/ErrorAlertComponent";
import DialogComponent from "@/components/DialogComponent";

export default {
  components:{
    DialogComponent,
    ListItemComponent,
    ErrorAlertComponent
  },
  name: "UsersView",
  data(){
    return{
      token: '',
      data: [],
      showError: false,
      errorMessage: '',
      currentCompanion: ''
    }
  },

  created(){
    this.token = localStorage.getItem("JWTToken")
    this.loadUsers()
  },

  methods:{

    closeError(){
      this.showError=false
      this.errorMessage = ''
    },

    onSelect(username){
      this.currentCompanion = username
    },

    loadUsers() {
      axios.post("http://localhost:8080/messages/all",
          {
            username: this.username,
            password: this.password
          },
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
        if (err.response ){
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
  position: relative;
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
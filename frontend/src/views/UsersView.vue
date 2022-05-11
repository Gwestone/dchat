<template>
  <div>
    <div v-for="(item) in data" :key="item.id">
      <ListItemComponent v-bind:name="item.username"/>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import ListItemComponent from "@/components/ListItemComponent";

export default {
  components:{
    ListItemComponent
  },
  name: "UsersView",
  data(){
    return{
      token: '',
      data: [],
      showError: false,
    }
  },

  created(){
    this.token = localStorage.getItem("JWTToken")
    this.loadUsers()
  },

  methods:{

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
          this.errorMessage = err.response.data.message
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

</style>
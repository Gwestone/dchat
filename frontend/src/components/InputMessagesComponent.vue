<template>
  <div>
    <button class="send__button" @click="send">Send</button>
    <input class="message__input" v-model="message">
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'InputMessagesComponent',
  props:{
    selectedUser: String
  },
  data(){
    return{
      token: localStorage.getItem('JWTToken'),
      message: ''
    }
  },
  methods:{
    send(){
      if (!this.message){
        return
      }

      console.log('send messages request')
      console.log(`send to ${this.selectedUser} message: ${this.message}`)
      axios.post('http://localhost:8080/messages/send/' + this.selectedUser,
          {
            'text': this.message
          },
          {
            headers: {
              'Authorization': `Bearer ${this.token}`
            }
          }
      ).then(res => {
        if (res.status === 200){
          this.message = ''
          this.$emit('messageSent')

        }else {
          this.errorMessage = res.data.message
          this.displayError()
        }
      }).catch(err => {
        if (err.response){
          this.errorMessage = 'failed to fetch data'
          console.log(err.response.data.error)
          this.displayError()
        }
      })

    }
  }
}
</script>

<style scoped>
.message__input{
  padding: 1px;
  float: left;
  width: 90%;
  margin-left: 20px;
  font-size:22px;
}
.send__button{
  float: left;
  padding: 5px;
}
</style>
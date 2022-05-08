<template>
  <div class="wrapper fadeInDown">
    <div id="formContent">
      <!-- Tabs Titles -->
      <h2 class="active"> Login </h2>

      <!-- Login Form -->
      <form>
        <input type="text" id="login" class="second" name="login" placeholder="login" v-model="username">
        <input type="text" id="password" class="third" name="login" placeholder="password" v-model="password">
        <input type="submit" class="fourth" value="Log In" @click="register">
      </form>

      <!-- Remind Passowrd -->
      <div id="formFooter">
        <a class="underlineHover" href="#">Forgot Password?</a>
      </div>

    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data(){
    return{
      username: '',
      password: ''
    }
  },
  methods: {

    register(e) {
      e.preventDefault()
      //console.log(this.username, this.password)
      axios.post("http://localhost:8080/auth/login",
          {
            username: this.username,
            password: this.password
          }).then(res => {
            if (res.status === 200){
              localStorage.setItem("JWTToken", res.data.token)
            }
      })

    }
  }
}
</script>

<style scoped src="@/assets/css/login.css" />

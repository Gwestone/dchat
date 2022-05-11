<template>
  <div>
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
    <div class="errorSpace">
      <transition name="slide-fade">
        <ErrorAlertComponent v-if="showError" :err-message="errorMessage" @close="closeError" class="error"/>
      </transition>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import ErrorAlertComponent from "@/components/ErrorAlertComponent";

export default {
  components:{
    ErrorAlertComponent
  },
  data(){
    return{
      username: '',
      password: '',
      showError: false,
      errorMessage: ''
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
              this.$forceUpdate()
              this.$emit("login")
              this.$router.push("/")

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

    closeError(){
      this.showError=false
      this.errorMessage = ''
    },

    displayError(){
      this.showError=true
    }
  }
}
</script>

<style scoped src="@/assets/css/login.css" />

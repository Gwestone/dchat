<template>
  <div>
    <div class="wrapper fadeInDown">
      <div id="formContent">
        <!-- Tabs Titles -->
        <h2 class="active"> Register </h2>

        <!-- Login Form -->
        <form>
          <input type="text" id="login" class="second" name="username" placeholder="username" v-model="username">
          <input type="text" id="password" class="third" name="password" placeholder="password" v-model="password">
          <input type="submit" class="fourth" value="Register" @click="register">
        </form>

      </div>
    </div>
    <div class="errorSpace">
      <transition name="slide-fade">
        <ErrorAlertComponent v-if="showError" :err-message="errorMessage" @close="closeError" class="error"/>
      </transition>
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
import validator from "validator";

export default {
  components:{
    ErrorAlertComponent,
  },
  emits:["register"],
  data(){
    return{
      username: '',
      password: '',
      showError: false,
      errorMessage: ''
    }
  },
  methods:{

    validate(){
      if (!validator.isStrongPassword(this.password)){
        console.log("not strong password")
        this.errorMessage = "password not strong"
        this.displayError()
        return false
      }else{
        return true
      }
    },

    register(e) {

      this.validate()
      e.preventDefault()
      //console.log(this.username, this.password)
      axios.post("http://localhost:8080/auth/register",
          {
            username: this.username,
            password: this.password
          }).then(res => {
        if (res.status === 200){

          localStorage.setItem("JWTToken", res.data.token)
          this.$forceUpdate()
          this.$emit("register")
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

<style scoped src="@/assets/css/register.css" />

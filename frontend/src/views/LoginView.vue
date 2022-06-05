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
          <input type="submit" class="fourth" value="Log In" @click="login">
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
import ErrorAlertComponent from "@/components/utils/ErrorAlertComponent";
import login from "@/api/login";

export default {
  components:{
    ErrorAlertComponent
  },
  emits:["login"],
  data(){
    return{
      username: '',
      password: '',
      showError: false,
      errorMessage: ''
    }
  },
  methods: {

    login(e) {

      e.preventDefault()
      //console.log(this.username, this.password)
      login(this.username, this.password, localStorage, this)

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

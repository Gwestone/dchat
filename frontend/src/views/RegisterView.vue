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
import register from "@/api/register";
import ErrorAlertComponent from "@/components/utils/ErrorAlertComponent";

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

    /*validate(){
      if (!validator.isStrongPassword(this.password)){
        console.log("not strong password")
        // this.errorMessage = "password not strong"
        // this.displayError()
        return false
      }else{
        return true
      }
    },*/

      register(e) {
        //this.validate()
        e.preventDefault()
        //console.log(this.username, this.password)
        register(this.username, this.password, localStorage, this)
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

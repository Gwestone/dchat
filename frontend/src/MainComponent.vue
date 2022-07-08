<template>
  <ul>
    <li><router-link to="/">Home</router-link></li>

    <li v-if="isLogged">
      <router-link to="users">Users</router-link>
    </li>

    <li v-if="isLogged">
      <router-link to="search">Search</router-link>
    </li>

    <li style="float:right" v-if="!isLogged">
      <router-link to="login">Login</router-link>
    </li>

    <li style="float:right" v-if="!isLogged">
      <router-link  to="register">Register</router-link>
    </li>

    <li style="float:right" v-if="isLogged">
      <a @click="logOut">LogOut</a>
    </li>
    <li style="float:right" v-if="isLogged">
      <div>{{ username   }}</div>
    </li>

  </ul>
  <router-view class="routerView" v-on:login="updateToken"/>
</template>

<script>
import jwtDecode from "jwt-decode";

export default {
  name: "MainComponent",
  data(){
    return{
      token: localStorage.getItem("JWTToken"),
      isLogged: !!localStorage.getItem("JWTToken"),
      username: ''
    }
  },

  mounted() {
    this.loadUser()
  },

  methods:{
    logOut(){
      localStorage.removeItem("JWTToken")
      this.isLogged = false
      this.$router.push({name: "home"})
    },

    updateToken(){
      this.token = localStorage.getItem("JWTToken")
      this.isLogged = !!this.token
      this.loadUser()
    },

    loadUser(){
      try {
        this.username = jwtDecode(localStorage.getItem("JWTToken")).Username
      }  catch (e){
        this.username = ''
        console.log(e)
      }
    },

  }
}
</script>

<style scoped src="@/assets/css/mainComponent.css" />
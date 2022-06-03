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
      <a @click="logOut" style="cursor: pointer;">LogOut</a>
    </li>

  </ul>
  <router-view class="routerView" v-on:login="onLogin" v-on:register="onRegister"/>
</template>

<script>
export default {
  name: "NavbarComponent",
  data(){
    return{
      token: localStorage.getItem("JWTToken"),
      isLogged: !!localStorage.getItem("JWTToken")
    }
  },

  mounted() {
    this.token = localStorage.getItem("JWTToken")
    this.isLogged = !!this.token
  },

  methods:{
    logOut(){
      localStorage.removeItem("JWTToken")
      this.isLogged = false
      this.$router.push({name: "home"})
    },

    onLogin(){
      this.token = localStorage.getItem("JWTToken")
      this.isLogged = !!this.token
    },

    onRegister(){
      this.token = localStorage.getItem("JWTToken")
      this.isLogged = !!this.token
    }

  }
}
</script>

<style scoped>
ul {
  list-style-type: none;
  margin: 0;
  padding: 5px;
  overflow: hidden;
  background-color: #333;
}

.routerView{
  height: calc(100% - 57px);
}

li {
  float: left;
}

li a {
  display: block;
  color: white;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
}

/* Change the link color to #111 (black) on hover */
li a:hover {
  background-color: #111;
}
</style>
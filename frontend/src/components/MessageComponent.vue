<template>
  <div class="message">
    {{ message.From }} : {{message.MessageText}} <div class="date">{{new Date(message.Date * 1000).getHours()}}:{{new Date(message.Date * 1000).getMinutes()}} </div>
  </div>
</template>

<script>
export default {
  name: "MessageComponent",
  props:{
    message: {
      type: Object,
      default() {
        return {
          from: "default user1",
          message: "hello world",
          date: 1652535907
        };
      }
    }
  }
}
</script>

<style scoped>
.message{
  font-size: 17px;
  padding: 15px;
  margin: 0;
  display: inline-block;
  position: relative;
  left: 0;
  top: 0;
  background: #f3ffff;
}
.date{
  color: #888888;
}
</style>
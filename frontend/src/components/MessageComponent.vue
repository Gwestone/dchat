<template>
  <div class="message">
    {{ message.from }} : {{message.message}} | {{new Date(message.date).getHours()}}:{{new Date(message.date).getMinutes()}}
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
</style>
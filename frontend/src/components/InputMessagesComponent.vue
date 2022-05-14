<template>
  <div>
    <button class="send__button">Send</button>
    <input class="message__input">
  </div>
</template>

<script>
export default {
  name: "InputMessagesComponent"
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
<template>
  <div class="dialog__component">
    <DisplayMessagesComponent class="message__display" :messages="messages" :current-username="currentUser" :redraw="redraw"/>
    <InputMessagesComponent class="message__input" :selected-user="selectedUser" @message-sent="send"/>
  </div>
</template>

<script>
import InputMessagesComponent from "@/components/InputMessagesComponent";
import DisplayMessagesComponent from "@/components/DisplayMessagesComponent";
export default {
  name: "DialogComponent",
  components: {InputMessagesComponent, DisplayMessagesComponent},
  props: {
    currentUser: String,
    messages: Array,
    selectedUser: String
  },
  data(){
    return{
      redraw: false
    }
  },
  methods:{
    send(){
      this.redraw = !this.redraw
      this.$emit("sent")
    }
  }
}
</script>

<style scoped>
.message__input{
  height: 100px;
  width: 90%;
  position: relative;
}
.message__display{
  height: calc(100% - 100px);
}
.dialog__component{
  bottom: 0;
  left: 0;
  height: 100%;
  position: relative;
}
</style>
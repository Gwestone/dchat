<template>
  <div class="messages__list">
    <div v-for="(item, i) in messages" :key="i" class="message__item" :align="messageOffset(item)">
      <MessageComponent :message="item"/>
    </div>
  </div>
</template>

<script>
import MessageComponent from "@/components/MessageComponent";
export default {
  name: "DisplayMessagesComponent",
  components: {MessageComponent},

  props: {
    messages: Array,
    currentUsername: String,
    redraw: Boolean
  },
  methods:{
    messageOffset(message){
      if(message.From === this.currentUsername){
        return "left"
      }else {
        return "right"
      }
    }
  }
}
</script>

<style scoped src="@/assets/css/displayMessagesComponent.css" />
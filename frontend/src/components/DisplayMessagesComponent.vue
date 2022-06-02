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
      if(message.from === this.currentUsername){
        return "left"
      }else {
        return "right"
      }
    }
  }
}
</script>

<style scoped>
.message__item{
  alignment: left;
}
.messages__list{
  display: block;
  position: relative;
  height: 800px;
  overflow:hidden; overflow-y:scroll;
}
/* width */
::-webkit-scrollbar {
  width: 10px;
}

/* Track */
::-webkit-scrollbar-track {
  background: #f1f1f1;
}

/* Handle */
::-webkit-scrollbar-thumb {
  background: #888;
}

/* Handle on hover */
::-webkit-scrollbar-thumb:hover {
  background: #555;
}

</style>
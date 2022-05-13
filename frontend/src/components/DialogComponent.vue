<template>
  <div v-if="to" class="dialog__component">
    <DisplayMessagesComponent />
    <InputMessagesComponent class="message__input"/>
  </div>
</template>

<script>
import InputMessagesComponent from "@/components/InputMessagesComponent";
import DisplayMessagesComponent from "@/components/DisplayMessagesComponent";
export default {
  name: "DialogComponent",
  components: {InputMessagesComponent, DisplayMessagesComponent},

  props: {
    to: String
  }
}
</script>

<style scoped>
.message__input{
  height: 100px;
  width: 100%;
  position: relative;
}
.dialog__component{
  bottom: 0;
  left: 0;
  position: relative;
}
</style>
<template>
  <div class="left">
    <div class="elem" v-for="(item, i) in users" :key="i">
      <ListItemComponent class="selected" :item="item" @selected="onSelect" v-if="selectedUsername === item.Username"/>
      <ListItemComponent :item="item" @selected="onSelect" v-else/>
    </div>
  </div>
</template>

<script>
import ListItemComponent from "@/components/ListItemComponent";

export default {
  name: "UsersListComponent",
  components:{
    ListItemComponent
  },
  props:{
    users: Array
  },
  methods:{
    onSelect(username){
      console.log(`selected ${username}`)
      this.selectedUsername = username
      this.$emit('selected', username)
    },
  },
  data(){
    return{
      selectedUsername: ''
    }
  }
}
</script>

<style scoped>
.selected{
  background: #2e2eff;
  color: whitesmoke;
}
</style>
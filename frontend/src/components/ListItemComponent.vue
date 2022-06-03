<template>
  <div class="card" @click="$emit('selected', item.Username)">
      {{ item.Username }}
  </div>
</template>

<script>
export default {
  name: "ListItemComponent",
  props:{
    item: {
      type: Object,
      required: true,
      default() {
        return {
          username: "default",
          id: -1
        }
      }
    }
  }
}
</script>

<style scoped>
  .card{
    font-size: 17px;
    padding: 10px;
    padding-left: 30px;
    padding-right: 30px;
    margin: 0;
    background: #f3ffff;
    border-radius: 4px;
    cursor: pointer;
  }
</style>
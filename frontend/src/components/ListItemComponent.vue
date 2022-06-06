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
          Username: "default",
          id: -1
        }
      }
    }
  }
}
</script>

<style scoped src="@/assets/css/listItemComponent.css" />
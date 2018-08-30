<template>
 <nav class="navbar">
  <div>

    <div class="navbar-header">
      <a class="navbar-brand" href="#">jsHarp</a>
    </div>
    
    <form class="form-inline my-2 my-lg-0" style="position: absolute; right: 1%;">
      <button class="btn" @click.prevent="$emit('render-page')">Run</button>
    </form>
  
  </div>
 </nav>
</template>

<script>
export default {

}
</script>

<style scoped>

  navbar {
    background-color: #1e1e1e;
  }

  .btn {
    margin-top: 20%;
    background-color: #1e1e1e;
    border-color: #4286f4;
    color: whitesmoke;
    transition: 0.5s;
  }
 
  .btn:hover {
    border: none;
    color: whitesmoke;
    background-color: #4286f4;
  }
</style>

<template>
  <nav class="navbar navbar-expand-lg border-bottom border-dark">
  <a class="navbar-brand" href="#">jsHarp</a>
  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>

  <div class="collapse navbar-collapse" id="navbarSupportedContent">
    <form class="form-inline my-2 my-lg-0" style="position: absolute; right: 1%;">
      <button class="btn btn-outline-primary my-2 my-sm-0" @click.prevent="$emit('render-page')">Run</button>
    </form>
  </div>
</nav>
</template>

<script>
export default {

}
</script>

<style>
  /* .btn { 
    background-color: #1e1e1e;
  } */
  .navbar {
    background-color: #1e1e1e;
  }
</style>

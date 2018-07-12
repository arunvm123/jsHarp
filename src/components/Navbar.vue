<template>
 <nav class="navbar">
  <div class="container-fluid">
    <!-- Brand and toggle get grouped for better mobile display -->
    <div class="navbar-header">
      <a class="navbar-brand" href="#">jsHarp</a>
    </div>

    <!-- Collect the nav links, forms, and other content for toggling -->
    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
      <ul class="nav navbar-nav navbar-right">
        <form class="form-inline my-2 my-lg-0" style="position: absolute; right: 1%;">
          <button class="btn" @click.prevent="$emit('render-page')">Run</button>
        </form>

      </ul>
    </div><!-- /.navbar-collapse -->
  </div><!-- /.container-fluid -->
</nav>
  <!-- <nav class="navbar navbar-expand-lg border-bottom border-dark">
  <a class="navbar-brand" href="#">jsHarp</a>
  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>

  <div class="collapse navbar-collapse" id="navbarSupportedContent">
    <form class="form-inline my-2 my-lg-0" style="position: absolute; right: 1%;">
      <button class="btn btn-outline-primary my-2 my-sm-0" @click.prevent="$emit('render-page')">Run</button>
    </form>
  </div>
</nav> -->
</template>

<script>
export default {

}
</script>

<style>
  /* .btn { 
    background-color: #1e1e1e;
  } */
  navbar {
    background-color: #1e1e1e;
  }

  .btn {
    margin-top: 20%;
    background-color: #1e1e1e;
    border-color: whitesmoke;
    color: white;
  }
 
  .btn:hover {

  }
</style>

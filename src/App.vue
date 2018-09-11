<template>
  <div id="app">
    <my-nav @render-page="renderPage"></my-nav>
    <div class="row">
      <div class="col-md-6 col-sm-6 col-lg-6 col-xs-6">
        <span style="color: #33FF00;">&lt!-- HTML --&gt</span>
        <my-editor v-model="html" language="html" @codeChange="html = $event"></my-editor>
        <span style="color: #33FF00;">/* CSS */</span>
        <my-editor v-model="css" language="css" @codeChange="css = $event"></my-editor>
        <span style="color: #33FF00;">// JS</span>
        <my-editor v-model="js" language="javascript" @codeChange="js = $event"></my-editor> 
      </div>
       <div class="col-md-6 col-sm-6 col-lg-6 col-xs-6">
        <iframe ref="mold" frameBorder=0 style="position: fixed; height: 240vh; width: 100%;" :src="renderURL">
          Your browser doesn't support iframes
        </iframe>
      </div>  
    </div>
  </div>
</template>

<script>
import Editor from './components/Editor.vue'
import Navbar from './components/Navbar.vue'

export default {
  components: {
    'my-editor': Editor,
    'my-nav': Navbar
  },
  name: 'app',
  data () {
    return {
      html: '',
      css: '',
      js: '',
      renderURL: '',
      bgColor: {
        'background-color': '#ffffff'
      },
      bgColorDisplay: true
    }
  },
  methods: {
    renderPage() {
      fetch('http://localhost:8000/code',{
        method: 'POST',
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          'html': this.html,
          'css': this.css,
          'js': this.js
        })
      }
    ).then(() => {
      this.$refs.mold.style.backgroundColor = '#ffffff'
      this.renderURL = 'http://localhost:8000/render'
      this.$refs.mold.src += ''
    })
      .catch(error => console.error(error))
    }
  }
}
</script>

<style>
  body {
     background-color: #1e1e1e;  
    overflow-x: hidden;
    height: 100%;
    width: 100%;
  }
</style>

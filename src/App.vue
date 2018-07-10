<template>
  <div id="app">
    <div class="row">
      <div class="col">
        <my-editor style="height: 33vh;" v-model="html" language="html" @codeChange="html = $event"></my-editor>
        <my-editor style="height: 33vh;" v-model="css" language="css" @codeChange="css = $event"></my-editor>
        <my-editor style="height: 33vh;" v-model="js" language="javascript" @codeChange="js = $event"></my-editor>
        <button @click="renderPage">Render</button>
       </div>
      <div class="col">
        <iframe ref="mold" style="height: 100%; width: 100%" :src="renderURL">
          Your browser doesn't support iframes
        </iframe>
      </div> 
     </div>
  </div>
</template>

<script>
import Editor from './components/Editor.vue'

export default {
  components: {
    'my-editor': Editor
  },
  name: 'app',
  data () {
    return {
      html: '',
      css: '',
      js: '',
      renderURL: ''
    }
  },
  methods: {
    typeLog() {
      console.log(this.html,this.css,this.js)
    },
    renderPage() {
      fetch('http://localhost:8080/code',{
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
      this.renderURL = 'http://localhost:8080/render'
      this.$refs.mold.src += '' 
    })
      .catch(error => console.log(error))

      // console.log(this.$refs.mold.contentDocument.getElementsByTagName('body'))
      // this.$refs.mold.contentDocument.getElementsByTagName('head')[0].innerHTML = (`<style>`+ this.css +`</style> \n <script>` + this.js + `<` + `/script>`)
      // this.$refs.mold.contentDocument.getElementsByTagName('body')[0].innerHTML = this.html
      // console.log(this.$refs.mold.contentDocument.getElementsByTagName('head'))
    }
  }
}
</script>

<style scoped>
  .fullPage {
    width: 100%;
    height: 100%;
  }
</style>

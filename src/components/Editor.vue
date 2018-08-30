<template>
  <div>
    <div row>
      <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12">
        <monaco style="margin-bottom: 0;"
          :language="language"
          :code="value"
          :editorOptions="options"
          @mounted="onMounted"
          @codeChange="onCodeChange"
          width="800"
          height="500"
          >
        </monaco>
      </div>
    </div> 
  </div>
</template>

<script>
import Monaco from 'vue-monaco-editor'

export default {
  props: {
    'value': String,
    'language': String
  },
  components: {
    Monaco
  },
 data() {
    return {
      options: {
        selectOnLineNumbers: true,
        automaticLayout: false
      }
    };
  },
  methods: {
    onMounted(editor) {
      this.editor = editor;
    },
    onCodeChange(editor) {
      this.$emit('codeChange',editor.getValue()) 
    }
  }
}
</script>

<style scoped>

</style>

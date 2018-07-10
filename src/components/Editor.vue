<template>
  <div>
<monaco
    :language="language"
    :code="value"
    :editorOptions="options"
    @mounted="onMounted"
    @codeChange="onCodeChange"
    >
</monaco>

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
        selectOnLineNumbers: true
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

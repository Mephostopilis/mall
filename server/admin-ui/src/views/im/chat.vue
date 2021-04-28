<template>
  <div class="app-container">
    <JwChat-index
      :tale-list="list"
      :tool-config="toolConfig"
      @enter="bindEnter"
    />
  </div>
</template>

<script>

export default {
  name: 'Chat',
  data() {
    return {
      list: [],
      tableHeader: [],
      toolConfig: {
        // 现在只配置了 ["file", "video", "img", "hongbao", "more", "history"]
        show: ['file', 'history', 'img', ['文件1', '', '美图']], // 二级数组中放自定义名称
        showEmoji: true // 是否显示表情图标
      }
    }
  },
  methods: {
    bindEnter() {
    }
  }
}
</script>

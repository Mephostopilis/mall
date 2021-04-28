<template>
  <div class="upload-container">
    <el-upload
      :data="dataObj"
      :multiple="false"
      :show-file-list="false"
      :on-success="handleImageSuccess"
      :before-upload="beforeUpload"
      class="image-uploader"
      drag
      action="http://192.168.21.95:8080/group1/upload"
    >
      <div v-if="imageUrl.length > 1">
        <!-- <div class="image-preview">
          <img :src="imageUrl">
          <div class="image-preview-action">
            <i class="el-icon-delete" @click="rmImage" />
          </div>
        </div> -->
        <el-image :src="imageUrl" fil="fill" />
      </div>
      <div v-else>
        <i class="el-icon-upload" />
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
      </div>
    </el-upload>
  </div>
</template>

<script>
// import { getToken } from '@/api/qiniu'
// import url from 'url'

export default {
  name: 'SingleImageUpload',
  props: {
    'onSuccess': {
      type: Function,
      default: (pathname) => {}
    },
    'onError': {
      type: Function,
      default: () => {}
    }
  },
  data() {
    return {
      tempUrl: '',
      dataObj: {
        token: '',
        key: '',
        output: '',
        scene: '',
        path: ''
      },
      imageUrl: '',
      pathname: ''
    }
  },
  computed: {
  },
  methods: {
    rmImage() {
      this.emitInput('')
    },
    emitInput(val) {
      this.$emit('input', val)
    },
    handleImageSuccess(response, file, fileList) {
      if (file.status === 'success') {
        // console.log(file)
        const x = new URL(file.response)
        this.imageUrl = x.origin + x.pathname
        this.pathname = x.pathname
        this.onSuccess(x.pathname)
      }
      // this.emitInput(this.tempUrl)
    },
    beforeUpload(file) {
      // const _self = this
      // return new Promise((resolve, reject) => {
      //   getToken()
      //     .then((response) => {
      //       const key = response.data.qiniu_key
      //       const token = response.data.qiniu_token
      //       _self._data.dataObj.token = token
      //       _self._data.dataObj.key = key
      //       this.tempUrl = response.data.qiniu_url
      //       resolve(true)
      //     })
      //     .catch((err) => {
      //       console.log(err)
      //       reject(false)
      //     })
      // })
      const isJPG = file.type === 'image/jpeg'
      const isLt2M = file.size / 1024 / 1024 < 2
      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!')
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!')
      }
      return isJPG && isLt2M
    }
  }
}
</script>

<style lang="scss" scoped>
@import "~@/styles/mixin.scss";
.upload-container {
  width: 100%;
  position: relative;
  @include clearfix;
  .image-uploader {
    width: 60%;
    float: left;
  }
  .image-preview {
    width: 200px;
    height: 200px;
    position: relative;
    border: 1px dashed #d9d9d9;
    float: left;
    margin-left: 50px;
    .image-preview-wrapper {
      position: relative;
      width: 100%;
      height: 100%;
      img {
        width: 100%;
        height: 100%;
      }
    }
    .image-preview-action {
      position: absolute;
      width: 100%;
      height: 100%;
      left: 0;
      top: 0;
      cursor: default;
      text-align: center;
      color: #fff;
      opacity: 0;
      font-size: 20px;
      background-color: rgba(0, 0, 0, 0.5);
      transition: opacity 0.3s;
      cursor: pointer;
      text-align: center;
      line-height: 200px;
      .el-icon-delete {
        font-size: 36px;
      }
    }
    &:hover {
      .image-preview-action {
        opacity: 1;
      }
    }
  }
}
</style>

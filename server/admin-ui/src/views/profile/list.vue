<template>
  <div>
    <div class="crumbs">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item> <i class="el-icon-lx-cascades" /> 快递公司管理 </el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="container">
      <div class="handle-box">
        <el-input v-model="query.name" placeholder="地区" class="handle-input mr10" style="width: 240px" />
        <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>
        <el-button icon="el-icon-lx-add" @click="handleAdd">{{ $t('i18n.btn_add') }}</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        border
        class="table"
        row-key="code"
        header-cell-class-name="table-header"
        :tree-props="{ children: 'children' }"
      >
        <!-- <el-table-column width="55" align="center">
                    <template slot-scope="scope">
                        {{ (scope.$index + 1) | indexFilter }}
                    </template>
                </el-table-column> -->
        <el-table-column prop="code" label="编码" align="center" />
        <el-table-column prop="name" label="名称" align="center" />
        <el-table-column prop="logo" label="logo">
          <template slot-scope="scope">
            <el-image
              v-if="scope.row.level != 0"
              slot="reference"
              :src="filePath + scope.row.icon"
              :alt="filePath + scope.row.icon"
              style="max-height: 50px; max-width: 50px"
            />
          </template>
        </el-table-column>
        <el-table-column prop="website" label="网址" align="center" />
        <el-table-column prop="sort" label="排序" align="center" />
        <el-table-column prop="createUser" label="创建人" align="center" />
        <el-table-column prop="createDate" label="创建时间" align="center" />
        <el-table-column label="操作" align="center">
          <template slot-scope="scope">
            <!-- <el-button type="text" icon="el-icon-view" @click="detailBtn(scope.$index, scope.row)">{{
                            $t('i18n.btn_detial')
                        }}</el-button> -->
            <span v-if="scope.row.level < 3">
              <el-button type="text" icon="el-icon-plus" @click="handleAddFromParent(scope.$index, scope.row)">{{
                $t('i18n.btn_add')
              }}</el-button>
            </span>
            <el-button type="text" icon="el-icon-edit" @click="handleEdit(scope.$index, scope.row)">{{
              $t('i18n.btn_edit')
            }}</el-button>
            <el-button type="text" icon="el-icon-delete" @click="handleDelete(scope.$index, scope.row)">{{
              $t('i18n.btn_delete')
            }}</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination
          background
          layout="total, prev, pager, next"
          :current-page="query.pageIndex"
          :page-size="query.pageSize"
          :total="pageTotal"
          @current-change="handlePageChange"
        />
      </div>
    </div>
    <!-- 编辑弹出框 -->
    <el-dialog
      :close-on-press-escape="false"
      :close-on-click-modal="false"
      :title="$t('i18n.btn_detial')"
      :visible.sync="dialogVisible"
      width="900px"
    >
      <el-form ref="form" :model="form.data" :rules="form.rules" label-width="100px" :disabled="!editAbleble">
        <el-row>
          <el-col :span="8">
            <p style="font-size: 17px; font-weight: bold">快递公司</p>
            <el-divider />
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item label="code" prop="id">
              <el-input v-model="form.data.code" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item label="名称" prop="name">
              <el-input v-model="form.data.name" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item label="logo" prop="logo">
              <el-upload
                :headers="headers()"
                class="avatar-uploader"
                :action="actionUploader"
                :show-file-list="false"
                :on-success="handleAvatarSuccess"
                :before-upload="beforeAvatarUpload"
              >
                <template v-if="form.icon">
                  <img :src="filePath+form.icon" class="avatar">
                </template>
                <template v-else>
                  <i class="el-icon-plus avatar-uploader-icon" />
                  <div class="el-upload__text">将文件拖到此处</br>或<em>点击上传</em></div>
                </template>
              </el-upload>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item label="网址" prop="website">
              <el-input v-model="form.data.website" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item label="排序" prop="sort">
              <el-input v-model="form.data.sort" />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button :disabled="!editAbleble" @click="saveEdit">保存</el-button>
        <el-button @click="closeDialog">{{ $t('i18n.btn_cancel') }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import qs from 'querystring'

export default {
  name: 'Couriercompanylist',
  filters: {},
  data() {
    return {
      module: this.$route.meta.module || 'product',
      query: {
        pageNo: 1,
        pageSize: 10
      },
      form: {
        data: {
          id: '',
          code: '',
          name: '',
          logo: '',
          website: '',
          sort: ''
        },
        rules: {
          code: [{ required: true, message: this.$t('i18n.admin_product_brand_list_name'), trigger: 'blur' }],
          name: [{ required: true, message: this.$t('i18n.admin_product_brand_list_name'), trigger: 'blur' }],
          logo: [{ required: true, message: this.$t('i18n.admin_product_brand_list_englishName'), trigger: 'blur' }],
          website: [{ required: true, message: this.$t('i18n.admin_product_brand_list_firstLetter'), trigger: 'blur' }],
          logo: [{ required: true, message: this.$t('i18n.admin_product_brand_list_logo'), trigger: 'blur' }],
          sort: [{ required: true, message: this.$t('i18n.admin_product_brand_list_brandStory'), trigger: 'blur' }]
        }
      },
      tableData: [],
      editVisible: false,
      editAbleble: false,
      pageTotal: 0,
      resources: [],
      idx: -1,
      actionUploader: this.constant.baseURL + '/file/upload/uploadFile',
      filePath: ''
    }
  },
  computed: {
    dialogVisible: {
      get: function() {
        return this.editVisible
      },
      set: function(value) {
        this.editVisible = value
      }
    }
  },
  watch: {
    editVisible: function(val, oldVla) {
      // console.log(this.$refs);
      if (this.$refs.form) {
        this.$refs.form.resetFields()
      }
    }
  },
  created() {
    this.getData()
    this.getFilePath()
  },
  methods: {
    // 获取 easy-mock 的模拟数据
    getData() {
      this.ajax()
        .get('/' + this.module + '/admin/product/couriercompany/list?' + qs.stringify(this.query))
        .then((res) => {
          // console.log('xx---------', res);
          this.tableData = res.records
          this.pageTotal = Number(res.total)
        })
    },
    // 添加操作
    handleAdd() {
      console.log('handleAdd')
      this.form.data = { id: null, showStatus: 1, lock: false }
      this.editVisible = true
      this.editAbleble = true
      this.editstatus = 'add'
    },
    // 添加操作
    handleAddFromParent(index, row) {
      console.log('handleAdd', index, row)
      this.form.data = { id: null, showStatus: 1, parentCode: row.code, parentName: row.name, level: row.level + 1, lock: true }
      this.editVisible = true
      this.editAbleble = true
      this.editstatus = 'add'
    },
    // 触发搜索按钮
    handleSearch() {
      this.$set(this.query, 'pageNo', 1)
      this.getData()
    },
    // 编辑操作
    handleEdit(index, row) {
      console.log('row:' + JSON.stringify(row))
      this.idx = index
      this.form.data = JSON.parse(JSON.stringify(row))
      this.editVisible = true
      this.editAbleble = true
      this.editstatus = 'update'
    },
    // 查看详情
    detailBtn(index, row) {
      this.idx = index
      this.form.data = JSON.parse(JSON.stringify(row))
      this.editVisible = true
      this.editAbleble = false
    },
    // 保存编辑
    saveEdit() {
      this.$refs.form.validate((valid, f) => {
        if (valid) {
          if (this.editstatus == 'update') {
            this.ajax()
              .put('/' + this.module + '/admin/product/couriercompany/update', this.form.data)
              .then((res) => {
                console.log('SaveEdit:update')
                this.$message.success(`修改第 ${this.idx + 1} 行成功`)
                this.editstatus = ''
                this.getData()
                this.editVisible = false
              })
          }
          if (this.editstatus == 'add') {
            this.ajax()
              .post('/' + this.module + '/admin/product/couriercompany/add', this.form.data)
              .then((res) => {
                console.log('SaveEdit:add')
                this.$message.success(`修改第 ${this.idx + 1} 行成功`)
                this.editstatus = ''
                this.getData()
                this.editVisible = false
              })
          }
        }
      })
    },
    closeDialog() {
      this.editVisible = false
      this.form.data.logo = ''
    },
    handleDelete(index, row) {
      // 二次确认删除
      this.$confirm('确定要删除吗？', '提示', { type: 'warning' })
        .then(() => {
          this.ajax()
            .delete(this.module + '/admin/product/couriercompany/del/' + row.id)
            .then((res) => {
              this.$message.success('删除成功')
              this.getData()
            })
        })
        .catch(() => {})
    },
    // 分页导航
    handlePageChange(val) {
      this.$set(this.query, 'pageNo', val)
      this.getData()
    },
    handleAvatarSuccess(res, file) {
      console.log(res)
      if (res.code == 0) {
        this.form.data.logo = res.data.filePath
        console.log('handleAvatarSuccess:' + this.form.logo)
        console.log(res)
        this.$forceUpdate()
      } else {
        this.$message.error(res.msg)
      }
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === 'image/jpeg'
      const isLt2M = file.size / 1024 / 1024 < 2

      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!')
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!')
      }
      return isJPG && isLt2M
    },
    getFilePath() {
      this.ajax()
        .get('/file/upload/getFilePath')
        .then((res) => {
          this.filePath = res
          console.log('filePath:' + this.filePath)
        })
    }
  }
}
</script>

<style scoped>
.handle-box {
    margin-bottom: 20px;
}

.handle-select {
    width: 120px;
}

.handle-input {
    width: 300px;
    display: inline-block;
}

.table {
    width: 100%;
    font-size: 14px;
}

.red {
    color: #ff0000;
}

.mr10 {
    margin-right: 10px;
}

.table-td-thumb {
    display: block;
    margin: auto;
    width: 40px;
    height: 40px;
}
.avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
}
.avatar-uploader .el-upload:hover {
    border-color: #409eff;
}
.avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
}
.avatar {
    width: 178px;
    height: 178px;
    display: block;
}
</style>

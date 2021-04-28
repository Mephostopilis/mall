
<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form
          ref="queryForm"
          :model="queryParams"
          :inline="true"
          label-width="68px"
        >
          <el-form-item>
            <el-button
              type="primary"
              icon="el-icon-search"
              size="mini"
              @click="handleQuery"
            >搜索</el-button>
            <el-button
              icon="el-icon-refresh"
              size="mini"
              @click="resetQuery"
            >重置</el-button>
          </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['pmsbrand:pmsbrand:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['pmsbrand:pmsbrand:edit']"
              type="success"
              icon="el-icon-edit"
              size="mini"
              :disabled="single"
              @click="handleUpdate"
            >修改
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['pmsbrand:pmsbrand:remove']"
              type="danger"
              icon="el-icon-delete"
              size="mini"
              :disabled="multiple"
              @click="handleDelete"
            >删除
            </el-button>
          </el-col>
        </el-row>

        <el-table
          v-loading="loading"
          :data="pmsbrandList"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column prop="name" label="名称" />
          <el-table-column prop="logo" label="logo">
            <template scope="scope">
              <el-image
                slot="reference"
                :src="baseUrl + scope.row.logo"
                :alt="baseUrl + scope.row.logo"
                style="max-height: 50px; max-width: 50px"
              />
            </template>
          </el-table-column>
          <el-table-column prop="factoryStatus" label="是否是制造商" />
          <el-table-column prop="productCount" label="产品数" />
          <el-table-column
            label="操作"
            align="center"
            class-name="small-padding fixed-width"
          >
            <template slot-scope="scope">
              <el-button
                v-permisaction="['pmsbrand:pmsbrand:edit']"
                size="mini"
                type="text"
                icon="el-icon-info"
                @click="handleDetail(scope.row)"
              >详情
              </el-button>
              <el-button
                v-permisaction="['pmsbrand:pmsbrand:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >修改
              </el-button>
              <el-button
                v-permisaction="['pmsbrand:pmsbrand:remove']"
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDelete(scope.row)"
              >删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>

        <pagination
          v-show="total > 0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="getList"
        />

        <!-- 添加或修改对话框 -->
        <el-dialog :title="title" :visible.sync="open" width="600px">
          <el-form ref="form" :model="form" :rules="rules" label-width="100px" :disabled="isDisabled">
            <el-form-item label="名称" prop="name">
              <el-input v-model="form.name" placeholder="" />
            </el-form-item>
            <el-form-item label="专区大图" prop="bigPic">
              <div v-if="form.logo && form.logo.length > 0">
                <el-image :src="baseUrl + form.logo" />
              </div>
              <div v-else>
                <single-image-upload :on-success="handleBigPicSuccess" />
              </div>
            </el-form-item>
            <el-form-item label="品牌故事" prop="brandStory">
              <el-input v-model="form.brandStory" placeholder="品牌故事" />
            </el-form-item>
            <el-form-item
              label="品牌制造商"
              prop="factoryStatus"
            >
              <el-select v-model="form.factoryStatus" placeholder="请选择">
                <el-option label="是" value="1" />
                <el-option label="否" value="0" />
              </el-select>
            </el-form-item>
            <el-form-item label="首字母" prop="firstLetter">
              <el-input v-model="form.firstLetter" placeholder="首字母" />
            </el-form-item>
            <el-form-item label="品牌logo" prop="logo">
              <div v-if="form.logo && form.logo.length > 0">
                <el-image :src="baseUrl + form.logo" />
              </div>
              <div v-else>
                <single-image-upload :on-success="handleLogoSuccess" />
              </div>
            </el-form-item>
            <el-form-item label="产品评论数量" prop="productCommentCount">
              <el-input
                v-model="form.productCommentCount"
                placeholder="产品评论数量"
              />
            </el-form-item>
            <el-form-item label="产品数量" prop="productCount">
              <el-input v-model="form.productCount" placeholder="产品数量" />
            </el-form-item>
            <el-form-item label="状态" prop="showStatus">
              <el-select v-model="form.showStatus" placeholder="请选择">
                <el-option label="是" value="1" />
                <el-option label="否" value="0" />
              </el-select>
            </el-form-item>
            <el-form-item label="排序" prop="sort">
              <el-input-number v-model="form.sort" placeholder="" />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import {
  addPmsBrand,
  delPmsBrand,
  getPmsBrand,
  listPmsBrand,
  updatePmsBrand
} from '@/api/pms/pmsbrand'
import SingleImageUpload from '@/components/Upload/SingleImage'
import qs from 'qs'

export default {
  name: 'PmsBrand',
  components: {
    SingleImageUpload
  },
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      isEdit: false,
      isDisabled: false,

      // 类型数据字典
      typeOptions: [],
      pmsbrandList: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {},
      // 文件服基础
      baseUrl: 'http://192.168.21.95:8080'
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listPmsBrand(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.pmsbrandList = response.data
          this.total = response.count
          this.loading = false
        }
      )
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {
        id: undefined,
        bigPic: undefined,
        brandStory: undefined,
        factoryStatus: undefined,
        firstLetter: undefined,
        logo: undefined,
        name: undefined,
        productCommentCount: undefined,
        productCount: undefined,
        showStatus: undefined,
        sort: undefined
      }
      this.resetForm('form')
    },

    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = []
      this.resetForm('queryForm')
      this.handleQuery()
    },
    /** 详情 */
    handleDetail(row) {
      this.reset()
      const id = row.id || this.ids
      getPmsBrand(id).then((response) => {
        this.form = response.data[0]
        this.open = true
        this.title = '品牌详情'
        this.isEdit = true
        this.isDisabled = true
      })
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.open = true
      this.title = '添加品牌表'
      this.isEdit = false
      this.isDisabled = false
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map((item) => item.id)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      const id = row.id || this.ids
      getPmsBrand(id).then((response) => {
        this.form = response.data[0]
        this.open = true
        this.title = '修改品牌表'
        this.isEdit = true
        this.isDisabled = false
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updatePmsBrand(qs.stringify(this.form)).then((response) => {
              if (response.code === 0) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addPmsBrand(qs.stringify(this.form)).then((response) => {
              if (response.code === 0) {
                this.msgSuccess('新增成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          }
        }
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const Ids = row.id || this.ids
      this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(function() {
          return delPmsBrand(Ids)
        })
        .then(() => {
          this.getList()
          this.msgSuccess('删除成功')
        })
        .catch(function() {})
    },
    handleBigPicSuccess(pathname) {
      this.form.bigPic = pathname
    },
    handleLogoSuccess(pathname) {
      this.form.logo = pathname
    }
  }
}
</script>

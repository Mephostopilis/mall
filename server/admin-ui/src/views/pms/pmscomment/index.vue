
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
              v-permisaction="['pmscomment:pmscomment:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['pmscomment:pmscomment:edit']"
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
              v-permisaction="['pmscomment:pmscomment:remove']"
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
          :data="pmscommentList"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column
            label="操作"
            align="center"
            class-name="small-padding fixed-width"
          >
            <template slot-scope="scope">
              <el-button
                v-permisaction="['pmscomment:pmscomment:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >修改
              </el-button>
              <el-button
                v-permisaction="['pmscomment:pmscomment:remove']"
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
        <el-dialog :title="title" :visible.sync="open" width="500px">
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">
            <el-form-item label="" prop="appId">
              <el-input v-model="form.appId" placeholder="" />
            </el-form-item>
            <el-form-item label="" prop="collectCouont">
              <el-input v-model="form.collectCouont" placeholder="" />
            </el-form-item>
            <el-form-item label="" prop="content">
              <el-input v-model="form.content" placeholder="" />
            </el-form-item>
            <el-form-item label="" prop="createTime">
              <el-input v-model="form.createTime" placeholder="" />
            </el-form-item>
            <el-form-item label="评论用户头像" prop="memberIcon">
              <el-input v-model="form.memberIcon" placeholder="评论用户头像" />
            </el-form-item>
            <el-form-item label="评价的ip" prop="memberIp">
              <el-input v-model="form.memberIp" placeholder="评价的ip" />
            </el-form-item>
            <el-form-item label="" prop="memberNickName">
              <el-input v-model="form.memberNickName" placeholder="" />
            </el-form-item>
            <el-form-item label="上传图片地址，以逗号隔开" prop="pics">
              <el-input
                v-model="form.pics"
                placeholder="上传图片地址，以逗号隔开"
              />
            </el-form-item>
            <el-form-item label="购买时的商品属性" prop="productAttribute">
              <el-input
                v-model="form.productAttribute"
                placeholder="购买时的商品属性"
              />
            </el-form-item>
            <el-form-item label="" prop="productId">
              <el-input v-model="form.productId" placeholder="" />
            </el-form-item>
            <el-form-item label="" prop="productName">
              <el-input v-model="form.productName" placeholder="" />
            </el-form-item>
            <el-form-item label="" prop="readCount">
              <el-input v-model="form.readCount" placeholder="" />
            </el-form-item>
            <el-form-item label="" prop="replayCount">
              <el-input v-model="form.replayCount" placeholder="" />
            </el-form-item>
            <el-form-item label="" prop="showStatus">
              <el-input v-model="form.showStatus" placeholder="" />
            </el-form-item>
            <el-form-item label="评价星数：0->5" prop="star">
              <el-input v-model="form.star" placeholder="评价星数：0->5" />
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
  addPmsComment,
  delPmsComment,
  getPmsComment,
  listPmsComment,
  updatePmsComment
} from '@/api/pms/pmscomment'

export default {
  name: 'PmsComment',
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
      // 类型数据字典
      typeOptions: [],
      pmscommentList: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {}
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listPmsComment(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.pmscommentList = response.data
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
        appId: undefined,
        collectCouont: undefined,
        content: undefined,
        createTime: undefined,
        id: undefined,
        memberIcon: undefined,
        memberIp: undefined,
        memberNickName: undefined,
        pics: undefined,
        productAttribute: undefined,
        productId: undefined,
        productName: undefined,
        readCount: undefined,
        replayCount: undefined,
        showStatus: undefined,
        star: undefined
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
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.open = true
      this.title = '添加商品评价表'
      this.isEdit = false
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
      getPmsComment(id).then((response) => {
        this.form = response.data
        this.open = true
        this.title = '修改商品评价表'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updatePmsComment(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addPmsComment(this.form).then((response) => {
              if (response.code === 200) {
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
          return delPmsComment(Ids)
        })
        .then(() => {
          this.getList()
          this.msgSuccess('删除成功')
        })
        .catch(function() {})
    }
  }
}
</script>

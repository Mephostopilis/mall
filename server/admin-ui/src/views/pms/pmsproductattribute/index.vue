
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
              v-permisaction="['pmsproductattribute:pmsproductattribute:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['pmsproductattribute:pmsproductattribute:edit']"
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
              v-permisaction="[
                'pmsproductattribute:pmsproductattribute:remove',
              ]"
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
          :data="pmsproductattributeList"
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
                v-permisaction="[
                  'pmsproductattribute:pmsproductattribute:edit',
                ]"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >修改
              </el-button>
              <el-button
                v-permisaction="[
                  'pmsproductattribute:pmsproductattribute:remove',
                ]"
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
            <el-form-item label="" prop="productAttributeCategoryId">
              <el-input
                v-model="form.productAttributeCategoryId"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="" prop="name">
              <el-input v-model="form.name" placeholder="" />
            </el-form-item>
            <el-form-item
              label="属性选择类型：0->唯一；1->单选；2->多选"
              prop="selectType"
            >
              <el-input
                v-model="form.selectType"
                placeholder="属性选择类型：0->唯一；1->单选；2->多选"
              />
            </el-form-item>
            <el-form-item
              label="属性录入方式：0->手工录入；1->从列表中选取"
              prop="inputType"
            >
              <el-input
                v-model="form.inputType"
                placeholder="属性录入方式：0->手工录入；1->从列表中选取"
              />
            </el-form-item>
            <el-form-item label="可选值列表，以逗号隔开" prop="inputList">
              <el-input
                v-model="form.inputList"
                placeholder="可选值列表，以逗号隔开"
              />
            </el-form-item>
            <el-form-item label="排序字段：最高的可以单独上传图片" prop="sort">
              <el-input
                v-model="form.sort"
                placeholder="排序字段：最高的可以单独上传图片"
              />
            </el-form-item>
            <el-form-item
              label="分类筛选样式：1->普通；1->颜色"
              prop="filterType"
            >
              <el-input
                v-model="form.filterType"
                placeholder="分类筛选样式：1->普通；1->颜色"
              />
            </el-form-item>
            <el-form-item
              label="检索类型；0->不需要进行检索；1->关键字检索；2->范围检索"
              prop="searchType"
            >
              <el-input
                v-model="form.searchType"
                placeholder="检索类型；0->不需要进行检索；1->关键字检索；2->范围检索"
              />
            </el-form-item>
            <el-form-item
              label="相同属性产品是否关联；0->不关联；1->关联"
              prop="relatedStatus"
            >
              <el-input
                v-model="form.relatedStatus"
                placeholder="相同属性产品是否关联；0->不关联；1->关联"
              />
            </el-form-item>
            <el-form-item
              label="是否支持手动新增；0->不支持；1->支持"
              prop="handAddStatus"
            >
              <el-input
                v-model="form.handAddStatus"
                placeholder="是否支持手动新增；0->不支持；1->支持"
              />
            </el-form-item>
            <el-form-item label="属性的类型；0->规格；1->参数" prop="type">
              <el-input
                v-model="form.type"
                placeholder="属性的类型；0->规格；1->参数"
              />
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
  addPmsProductAttribute,
  delPmsProductAttribute,
  getPmsProductAttribute,
  listPmsProductAttribute,
  updatePmsProductAttribute
} from '@/api/pms/pmsproductattribute'
import qs from 'qs'
export default {
  name: 'Config',
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
      pmsproductattributeList: [],

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
      listPmsProductAttribute(
        this.addDateRange(this.queryParams, this.dateRange)
      ).then((response) => {
        this.pmsproductattributeList = response.data
        this.total = response.count
        this.loading = false
      })
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
        appId: undefined,
        productAttributeCategoryId: undefined,
        name: undefined,
        selectType: undefined,
        inputType: undefined,
        inputList: undefined,
        sort: undefined,
        filterType: undefined,
        searchType: undefined,
        relatedStatus: undefined,
        handAddStatus: undefined,
        type: undefined
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
      this.title = '添加商品属性参数表'
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
      getPmsProductAttribute(id).then((response) => {
        this.form = response.data
        this.open = true
        this.title = '修改商品属性参数表'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updatePmsProductAttribute(qs.stringify(this.form)).then(
              (response) => {
                if (response.code === 0) {
                  this.msgSuccess('修改成功')
                  this.open = false
                  this.getList()
                } else {
                  this.msgError(response.msg)
                }
              }
            )
          } else {
            addPmsProductAttribute(qs.stringify(this.form)).then((response) => {
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
          return delPmsProductAttribute(Ids)
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

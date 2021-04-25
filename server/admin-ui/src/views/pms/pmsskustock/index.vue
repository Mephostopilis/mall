
<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">

          <el-form-item>
            <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
            <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
          </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['pmsskustock:pmsskustock:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['pmsskustock:pmsskustock:edit']"
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
              v-permisaction="['pmsskustock:pmsskustock:remove']"
              type="danger"
              icon="el-icon-delete"
              size="mini"
              :disabled="multiple"
              @click="handleDelete"
            >删除
            </el-button>
          </el-col>
        </el-row>

        <el-table v-loading="loading" :data="pmsskustockList" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['pmsskustock:pmsskustock:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >修改
              </el-button>
              <el-button
                v-permisaction="['pmsskustock:pmsskustock:remove']"
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
          v-show="total>0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="getList"
        />

        <!-- 添加或修改对话框 -->
        <el-dialog :title="title" :visible.sync="open" width="500px">
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">

            <el-form-item label="" prop="appId">
              <el-input
                v-model="form.appId"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="锁定库存" prop="lockStock">
              <el-input
                v-model="form.lockStock"
                placeholder="锁定库存"
              />
            </el-form-item>
            <el-form-item label="预警库存" prop="lowStock">
              <el-input
                v-model="form.lowStock"
                placeholder="预警库存"
              />
            </el-form-item>
            <el-form-item label="展示图片" prop="pic">
              <el-input
                v-model="form.pic"
                placeholder="展示图片"
              />
            </el-form-item>
            <el-form-item label="" prop="price">
              <el-input
                v-model="form.price"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="" prop="productId">
              <el-input
                v-model="form.productId"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="单品促销价格" prop="promotionPrice">
              <el-input
                v-model="form.promotionPrice"
                placeholder="单品促销价格"
              />
            </el-form-item>
            <el-form-item label="销量" prop="sale">
              <el-input
                v-model="form.sale"
                placeholder="销量"
              />
            </el-form-item>
            <el-form-item label="sku编码" prop="skuCode">
              <el-input
                v-model="form.skuCode"
                placeholder="sku编码"
              />
            </el-form-item>
            <el-form-item label="商品销售属性，json格式" prop="spData">
              <el-input
                v-model="form.spData"
                placeholder="商品销售属性，json格式"
              />
            </el-form-item>
            <el-form-item label="库存" prop="stock">
              <el-input
                v-model="form.stock"
                placeholder="库存"
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
import { addPmsSkuStock, delPmsSkuStock, getPmsSkuStock, listPmsSkuStock, updatePmsSkuStock } from '@/api/pms/pmsskustock'

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
      pmsskustockList: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10

      },
      // 表单参数
      form: {
      },
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
      listPmsSkuStock(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.pmsskustockList = response.data
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
        id: undefined,
        lockStock: undefined,
        lowStock: undefined,
        pic: undefined,
        price: undefined,
        productId: undefined,
        promotionPrice: undefined,
        sale: undefined,
        skuCode: undefined,
        spData: undefined,
        stock: undefined
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
      this.title = '添加sku的库存'
      this.isEdit = false
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      const id =
                row.id || this.ids
      getPmsSkuStock(id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改sku的库存'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updatePmsSkuStock(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addPmsSkuStock(this.form).then(response => {
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
      }).then(function() {
        return delPmsSkuStock(Ids)
      }).then(() => {
        this.getList()
        this.msgSuccess('删除成功')
      }).catch(function() {
      })
    }
  }
}
</script>

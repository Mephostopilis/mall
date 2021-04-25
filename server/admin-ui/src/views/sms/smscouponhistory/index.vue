
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
              v-permisaction="['smscouponhistory:smscouponhistory:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['smscouponhistory:smscouponhistory:edit']"
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
              v-permisaction="['smscouponhistory:smscouponhistory:remove']"
              type="danger"
              icon="el-icon-delete"
              size="mini"
              :disabled="multiple"
              @click="handleDelete"
            >删除
            </el-button>
          </el-col>
        </el-row>

        <el-table v-loading="loading" :data="smscouponhistoryList" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['smscouponhistory:smscouponhistory:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >修改
              </el-button>
              <el-button
                v-permisaction="['smscouponhistory:smscouponhistory:remove']"
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
            <el-form-item label="" prop="couponCode">
              <el-input
                v-model="form.couponCode"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="" prop="couponId">
              <el-input
                v-model="form.couponId"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="" prop="createTime">
              <el-input
                v-model="form.createTime"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="获取类型：0->后台赠送；1->主动获取" prop="getType">
              <el-input
                v-model="form.getType"
                placeholder="获取类型：0->后台赠送；1->主动获取"
              />
            </el-form-item>
            <el-form-item label="" prop="memberId">
              <el-input
                v-model="form.memberId"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="领取人昵称" prop="memberNickname">
              <el-input
                v-model="form.memberNickname"
                placeholder="领取人昵称"
              />
            </el-form-item>
            <el-form-item label="订单编号" prop="orderId">
              <el-input
                v-model="form.orderId"
                placeholder="订单编号"
              />
            </el-form-item>
            <el-form-item label="订单号码" prop="orderSn">
              <el-input
                v-model="form.orderSn"
                placeholder="订单号码"
              />
            </el-form-item>
            <el-form-item label="使用状态：0->未使用；1->已使用；2->已过期" prop="useStatus">
              <el-input
                v-model="form.useStatus"
                placeholder="使用状态：0->未使用；1->已使用；2->已过期"
              />
            </el-form-item>
            <el-form-item label="使用时间" prop="useTime">
              <el-input
                v-model="form.useTime"
                placeholder="使用时间"
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
import { addSmsCouponHistory, delSmsCouponHistory, getSmsCouponHistory, listSmsCouponHistory, updateSmsCouponHistory } from '@/api/sms/smscouponhistory'

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
      smscouponhistoryList: [],

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
      listSmsCouponHistory(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.smscouponhistoryList = response.data
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
        couponCode: undefined,
        couponId: undefined,
        createTime: undefined,
        getType: undefined,
        id: undefined,
        memberId: undefined,
        memberNickname: undefined,
        orderId: undefined,
        orderSn: undefined,
        useStatus: undefined,
        useTime: undefined
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
      this.title = '添加优惠券使用、领取历史表'
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
      getSmsCouponHistory(id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改优惠券使用、领取历史表'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateSmsCouponHistory(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addSmsCouponHistory(this.form).then(response => {
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
        return delSmsCouponHistory(Ids)
      }).then(() => {
        this.getList()
        this.msgSuccess('删除成功')
      }).catch(function() {
      })
    }
  }
}
</script>

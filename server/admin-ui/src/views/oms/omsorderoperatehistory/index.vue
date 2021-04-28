
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
                                v-permisaction="['omsorderoperatehistory:omsorderoperatehistory:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['omsorderoperatehistory:omsorderoperatehistory:edit']"
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
                                v-permisaction="['omsorderoperatehistory:omsorderoperatehistory:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="omsorderoperatehistoryList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/>
                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                            <el-button
                                    v-permisaction="['omsorderoperatehistory:omsorderoperatehistory:edit']"
                                    size="mini"
                                    type="text"
                                    icon="el-icon-edit"
                                    @click="handleUpdate(scope.row)"
                            >修改
                            </el-button>
                            <el-button
                                    v-permisaction="['omsorderoperatehistory:omsorderoperatehistory:remove']"
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
                        
                                    <el-form-item label="订单id" prop="orderId">
                                        <el-input v-model="form.orderId" placeholder="订单id"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="操作人：用户；系统；后台管理员" prop="operateMan">
                                        <el-input v-model="form.operateMan" placeholder="操作人：用户；系统；后台管理员"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="操作时间" prop="createTime">
                                        <el-input v-model="form.createTime" placeholder="操作时间"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单" prop="orderStatus">
                                        <el-input v-model="form.orderStatus" placeholder="订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="备注" prop="note">
                                        <el-input v-model="form.note" placeholder="备注"
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
    import {addOmsOrderOperateHistory, delOmsOrderOperateHistory, getOmsOrderOperateHistory, listOmsOrderOperateHistory, updateOmsOrderOperateHistory} from '@/api/oms/omsorderoperatehistory'

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
                omsorderoperatehistoryList: [],
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
            
        },
            // 表单参数
            form: {
            }
        ,
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
                listOmsOrderOperateHistory(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.omsorderoperatehistoryList = response.data
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
                orderId: undefined,
                operateMan: undefined,
                createTime: undefined,
                orderStatus: undefined,
                note: undefined,
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
                this.title = '添加订单操作历史记录'
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
                getOmsOrderOperateHistory(id).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改订单操作历史记录'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateOmsOrderOperateHistory(this.form).then(response => {
                                if (response.code === 0) {
                                    this.msgSuccess('修改成功')
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addOmsOrderOperateHistory(this.form).then(response => {
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
                }).then(function () {
                    return delOmsOrderOperateHistory(Ids)
                }).then(() => {
                    this.getList()
                    this.msgSuccess('删除成功')
                }).catch(function () {
                })
            }
        }
    }
</script>

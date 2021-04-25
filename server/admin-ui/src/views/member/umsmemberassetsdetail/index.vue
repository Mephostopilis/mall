
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
                                v-permisaction="['umsmemberassetsdetail:umsmemberassetsdetail:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['umsmemberassetsdetail:umsmemberassetsdetail:edit']"
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
                                v-permisaction="['umsmemberassetsdetail:umsmemberassetsdetail:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="umsmemberassetsdetailList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/>
                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                            <el-button
                                    v-permisaction="['umsmemberassetsdetail:umsmemberassetsdetail:edit']"
                                    size="mini"
                                    type="text"
                                    icon="el-icon-edit"
                                    @click="handleUpdate(scope.row)"
                            >修改
                            </el-button>
                            <el-button
                                    v-permisaction="['umsmemberassetsdetail:umsmemberassetsdetail:remove']"
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
                        
                                    <el-form-item label="code" prop="access">
                                        <el-input v-model="form.access" placeholder="code"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="code创建时" prop="accessCreatedAt">
                                        <el-date-picker
                                                    v-model="form.accessCreatedAt"
                                                    type="datetime"
                                                    placeholder="选择日期">
                                            </el-date-picker>
                                    </el-form-item>
                                    <el-form-item label="code创建时" prop="accessExpiresIn">
                                        <el-input v-model="form.accessExpiresIn" placeholder="code创建时"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="客户端id" prop="clientId">
                                        <el-input v-model="form.clientId" placeholder="客户端id"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="code" prop="code">
                                        <el-input v-model="form.code" placeholder="code"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="code创建时" prop="codeCreatedAt">
                                        <el-date-picker
                                                    v-model="form.codeCreatedAt"
                                                    type="datetime"
                                                    placeholder="选择日期">
                                            </el-date-picker>
                                    </el-form-item>
                                    <el-form-item label="code创建时" prop="codeExpiresIn">
                                        <el-input v-model="form.codeExpiresIn" placeholder="code创建时"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="redirect" prop="redirectUri">
                                        <el-input v-model="form.redirectUri" placeholder="redirect"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="code" prop="refresh">
                                        <el-input v-model="form.refresh" placeholder="code"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="code创建时" prop="refreshCreatedAt">
                                        <el-date-picker
                                                    v-model="form.refreshCreatedAt"
                                                    type="datetime"
                                                    placeholder="选择日期">
                                            </el-date-picker>
                                    </el-form-item>
                                    <el-form-item label="code创建时" prop="refreshExpiresIn">
                                        <el-input v-model="form.refreshExpiresIn" placeholder="code创建时"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="范围" prop="scope">
                                        <el-input v-model="form.scope" placeholder="范围"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="用户id" prop="userId">
                                        <el-input v-model="form.userId" placeholder="用户id"
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
    import {addUmsMemberAssetsDetail, delUmsMemberAssetsDetail, getUmsMemberAssetsDetail, listUmsMemberAssetsDetail, updateUmsMemberAssetsDetail} from '@/api/member/umsmemberassetsdetail'

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
                umsmemberassetsdetailList: [],
                
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
                listUmsMemberAssetsDetail(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.umsmemberassetsdetailList = response.data
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
                
                access: undefined,
                accessCreatedAt: undefined,
                accessExpiresIn: undefined,
                clientId: undefined,
                code: undefined,
                codeCreatedAt: undefined,
                codeExpiresIn: undefined,
                id: undefined,
                redirectUri: undefined,
                refresh: undefined,
                refreshCreatedAt: undefined,
                refreshExpiresIn: undefined,
                scope: undefined,
                userId: undefined,
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
                this.title = '添加UmsMemberAssetsDetail'
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
                getUmsMemberAssetsDetail(id).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改UmsMemberAssetsDetail'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateUmsMemberAssetsDetail(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess('修改成功')
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addUmsMemberAssetsDetail(this.form).then(response => {
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
                }).then(function () {
                    return delUmsMemberAssetsDetail(Ids)
                }).then(() => {
                    this.getList()
                    this.msgSuccess('删除成功')
                }).catch(function () {
                })
            }
        }
    }
</script>


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
                                v-permisaction="['omsorderitem:omsorderitem:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['omsorderitem:omsorderitem:edit']"
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
                                v-permisaction="['omsorderitem:omsorderitem:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="omsorderitemList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/>
                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                            <el-button
                                    v-permisaction="['omsorderitem:omsorderitem:edit']"
                                    size="mini"
                                    type="text"
                                    icon="el-icon-edit"
                                    @click="handleUpdate(scope.row)"
                            >修改
                            </el-button>
                            <el-button
                                    v-permisaction="['omsorderitem:omsorderitem:remove']"
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
                                    <el-form-item label="订单编号" prop="orderSn">
                                        <el-input v-model="form.orderSn" placeholder="订单编号"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="productId">
                                        <el-input v-model="form.productId" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="productPic">
                                        <el-input v-model="form.productPic" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="productName">
                                        <el-input v-model="form.productName" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="productBrand">
                                        <el-input v-model="form.productBrand" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="productSn">
                                        <el-input v-model="form.productSn" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="销售价格" prop="productPrice">
                                        <el-input v-model="form.productPrice" placeholder="销售价格"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="购买数量" prop="productQuantity">
                                        <el-input v-model="form.productQuantity" placeholder="购买数量"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="商品sku编号" prop="productSkuId">
                                        <el-input v-model="form.productSkuId" placeholder="商品sku编号"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="商品sku条码" prop="productSkuCode">
                                        <el-input v-model="form.productSkuCode" placeholder="商品sku条码"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="商品分类id" prop="productCategoryId">
                                        <el-input v-model="form.productCategoryId" placeholder="商品分类id"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="商品促销名称" prop="promotionName">
                                        <el-input v-model="form.promotionName" placeholder="商品促销名称"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="商品促销分解金额" prop="promotionAmount">
                                        <el-input v-model="form.promotionAmount" placeholder="商品促销分解金额"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="优惠券优惠分解金额" prop="couponAmount">
                                        <el-input v-model="form.couponAmount" placeholder="优惠券优惠分解金额"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="积分优惠分解金额" prop="integrationAmount">
                                        <el-input v-model="form.integrationAmount" placeholder="积分优惠分解金额"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="该商品经过优惠后的分解金额" prop="realAmount">
                                        <el-input v-model="form.realAmount" placeholder="该商品经过优惠后的分解金额"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="giftIntegration">
                                        <el-input v-model="form.giftIntegration" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="giftGrowth">
                                        <el-input v-model="form.giftGrowth" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]" prop="productAttr">
                                        <el-input v-model="form.productAttr" placeholder="商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]"
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
    import {addOmsOrderItem, delOmsOrderItem, getOmsOrderItem, listOmsOrderItem, updateOmsOrderItem} from '@/api/oms/omsorderitem'

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
                omsorderitemList: [],
                
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
                listOmsOrderItem(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.omsorderitemList = response.data
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
                orderSn: undefined,
                productId: undefined,
                productPic: undefined,
                productName: undefined,
                productBrand: undefined,
                productSn: undefined,
                productPrice: undefined,
                productQuantity: undefined,
                productSkuId: undefined,
                productSkuCode: undefined,
                productCategoryId: undefined,
                promotionName: undefined,
                promotionAmount: undefined,
                couponAmount: undefined,
                integrationAmount: undefined,
                realAmount: undefined,
                giftIntegration: undefined,
                giftGrowth: undefined,
                productAttr: undefined,
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
                this.title = '添加订单中所包含的商品'
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
                getOmsOrderItem(id).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改订单中所包含的商品'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateOmsOrderItem(this.form).then(response => {
                                if (response.code === 0) {
                                    this.msgSuccess('修改成功')
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addOmsOrderItem(this.form).then(response => {
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
                    return delOmsOrderItem(Ids)
                }).then(() => {
                    this.getList()
                    this.msgSuccess('删除成功')
                }).catch(function () {
                })
            }
        }
    }
</script>


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
                                v-permisaction="['omsorder:omsorder:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['omsorder:omsorder:edit']"
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
                                v-permisaction="['omsorder:omsorder:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="omsorderList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/>
                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                            <el-button
                                    v-permisaction="['omsorder:omsorder:edit']"
                                    size="mini"
                                    type="text"
                                    icon="el-icon-edit"
                                    @click="handleUpdate(scope.row)"
                            >修改
                            </el-button>
                            <el-button
                                    v-permisaction="['omsorder:omsorder:remove']"
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
                        
                                    <el-form-item label="" prop="memberId">
                                        <el-input v-model="form.memberId" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="couponId">
                                        <el-input v-model="form.couponId" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="订单编号" prop="orderSn">
                                        <el-input v-model="form.orderSn" placeholder="订单编号"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="提交时间" prop="createTime">
                                        <el-input v-model="form.createTime" placeholder="提交时间"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="用户帐号" prop="memberUsername">
                                        <el-input v-model="form.memberUsername" placeholder="用户帐号"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="订单总金额" prop="totalAmount">
                                        <el-input v-model="form.totalAmount" placeholder="订单总金额"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="应付金额（实际支付金额）" prop="payAmount">
                                        <el-input v-model="form.payAmount" placeholder="应付金额（实际支付金额）"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="运费金额" prop="freightAmount">
                                        <el-input v-model="form.freightAmount" placeholder="运费金额"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="促销优化金额（促销价、满减、阶梯价）" prop="promotionAmount">
                                        <el-input v-model="form.promotionAmount" placeholder="促销优化金额（促销价、满减、阶梯价）"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="积分抵扣金额" prop="integrationAmount">
                                        <el-input v-model="form.integrationAmount" placeholder="积分抵扣金额"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="优惠券抵扣金额" prop="couponAmount">
                                        <el-input v-model="form.couponAmount" placeholder="优惠券抵扣金额"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="管理员后台调整订单使用的折扣金额" prop="discountAmount">
                                        <el-input v-model="form.discountAmount" placeholder="管理员后台调整订单使用的折扣金额"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="支付方式：0->未支付；1->支付宝；2->微信" prop="payType">
                                        <el-input v-model="form.payType" placeholder="支付方式：0->未支付；1->支付宝；2->微信"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="订单来源：0->PC订单；1->app订单" prop="sourceType">
                                        <el-input v-model="form.sourceType" placeholder="订单来源：0->PC订单；1->app订单"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单" prop="status">
                                        <el-input v-model="form.status" placeholder="订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="订单类型：0->正常订单；1->秒杀订单" prop="orderType">
                                        <el-input v-model="form.orderType" placeholder="订单类型：0->正常订单；1->秒杀订单"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="物流公司(配送方式)" prop="deliveryCompany">
                                        <el-input v-model="form.deliveryCompany" placeholder="物流公司(配送方式)"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="物流单号" prop="deliverySn">
                                        <el-input v-model="form.deliverySn" placeholder="物流单号"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="自动确认时间（天）" prop="autoConfirmDay">
                                        <el-input v-model="form.autoConfirmDay" placeholder="自动确认时间（天）"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="可以获得的积分" prop="integration">
                                        <el-input v-model="form.integration" placeholder="可以获得的积分"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="可以活动的成长值" prop="growth">
                                        <el-input v-model="form.growth" placeholder="可以活动的成长值"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="活动信息" prop="promotionInfo">
                                        <el-input v-model="form.promotionInfo" placeholder="活动信息"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="发票类型：0->不开发票；1->电子发票；2->纸质发票" prop="billType">
                                        <el-input v-model="form.billType" placeholder="发票类型：0->不开发票；1->电子发票；2->纸质发票"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="发票抬头" prop="billHeader">
                                        <el-input v-model="form.billHeader" placeholder="发票抬头"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="发票内容" prop="billContent">
                                        <el-input v-model="form.billContent" placeholder="发票内容"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="收票人电话" prop="billReceiverPhone">
                                        <el-input v-model="form.billReceiverPhone" placeholder="收票人电话"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="收票人邮箱" prop="billReceiverEmail">
                                        <el-input v-model="form.billReceiverEmail" placeholder="收票人邮箱"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="收货人姓名" prop="receiverName">
                                        <el-input v-model="form.receiverName" placeholder="收货人姓名"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="收货人电话" prop="receiverPhone">
                                        <el-input v-model="form.receiverPhone" placeholder="收货人电话"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="收货人邮编" prop="receiverPostCode">
                                        <el-input v-model="form.receiverPostCode" placeholder="收货人邮编"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="省份/直辖市" prop="receiverProvince">
                                        <el-input v-model="form.receiverProvince" placeholder="省份/直辖市"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="城市" prop="receiverCity">
                                        <el-input v-model="form.receiverCity" placeholder="城市"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="区" prop="receiverRegion">
                                        <el-input v-model="form.receiverRegion" placeholder="区"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="详细地址" prop="receiverDetailAddress">
                                        <el-input v-model="form.receiverDetailAddress" placeholder="详细地址"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="订单备注" prop="note">
                                        <el-input v-model="form.note" placeholder="订单备注"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="确认收货状态：0->未确认；1->已确认" prop="confirmStatus">
                                        <el-input v-model="form.confirmStatus" placeholder="确认收货状态：0->未确认；1->已确认"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="删除状态：0->未删除；1->已删除" prop="deleteStatus">
                                        <el-input v-model="form.deleteStatus" placeholder="删除状态：0->未删除；1->已删除"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="下单时使用的积分" prop="useIntegration">
                                        <el-input v-model="form.useIntegration" placeholder="下单时使用的积分"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="支付时间" prop="paymentTime">
                                        <el-input v-model="form.paymentTime" placeholder="支付时间"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="发货时间" prop="deliveryTime">
                                        <el-input v-model="form.deliveryTime" placeholder="发货时间"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="确认收货时间" prop="receiveTime">
                                        <el-input v-model="form.receiveTime" placeholder="确认收货时间"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="评价时间" prop="commentTime">
                                        <el-input v-model="form.commentTime" placeholder="评价时间"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="修改时间" prop="modifyTime">
                                        <el-input v-model="form.modifyTime" placeholder="修改时间"
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
    import {addOmsOrder, delOmsOrder, getOmsOrder, listOmsOrder, updateOmsOrder} from '@/api/oms/omsorder'

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
                omsorderList: [],
                
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
                listOmsOrder(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.omsorderList = response.data
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
                memberId: undefined,
                couponId: undefined,
                orderSn: undefined,
                createTime: undefined,
                memberUsername: undefined,
                totalAmount: undefined,
                payAmount: undefined,
                freightAmount: undefined,
                promotionAmount: undefined,
                integrationAmount: undefined,
                couponAmount: undefined,
                discountAmount: undefined,
                payType: undefined,
                sourceType: undefined,
                status: undefined,
                orderType: undefined,
                deliveryCompany: undefined,
                deliverySn: undefined,
                autoConfirmDay: undefined,
                integration: undefined,
                growth: undefined,
                promotionInfo: undefined,
                billType: undefined,
                billHeader: undefined,
                billContent: undefined,
                billReceiverPhone: undefined,
                billReceiverEmail: undefined,
                receiverName: undefined,
                receiverPhone: undefined,
                receiverPostCode: undefined,
                receiverProvince: undefined,
                receiverCity: undefined,
                receiverRegion: undefined,
                receiverDetailAddress: undefined,
                note: undefined,
                confirmStatus: undefined,
                deleteStatus: undefined,
                useIntegration: undefined,
                paymentTime: undefined,
                deliveryTime: undefined,
                receiveTime: undefined,
                commentTime: undefined,
                modifyTime: undefined,
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
                this.title = '添加订单表'
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
                getOmsOrder(id).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改订单表'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateOmsOrder(this.form).then(response => {
                                if (response.code === 0) {
                                    this.msgSuccess('修改成功')
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addOmsOrder(this.form).then(response => {
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
                    return delOmsOrder(Ids)
                }).then(() => {
                    this.getList()
                    this.msgSuccess('删除成功')
                }).catch(function () {
                })
            }
        }
    }
</script>

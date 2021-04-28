
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
              v-permisaction="['pmsproduct:pmsproduct:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['pmsproduct:pmsproduct:edit']"
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
              v-permisaction="['pmsproduct:pmsproduct:remove']"
              type="danger"
              icon="el-icon-delete"
              size="mini"
              :disabled="multiple"
              @click="handleDelete"
            >删除
            </el-button>
          </el-col>
        </el-row>

        <el-table v-loading="loading" :data="pmsproductList" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['pmsproduct:pmsproduct:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >修改
              </el-button>
              <el-button
                v-permisaction="['pmsproduct:pmsproduct:remove']"
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
            <el-form-item label="" prop="brandId">
              <el-input
                v-model="form.brandId"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="" prop="productCategoryId">
              <el-input
                v-model="form.productCategoryId"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="" prop="feightTemplateId">
              <el-input
                v-model="form.feightTemplateId"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="" prop="productAttributeCategoryId">
              <el-input
                v-model="form.productAttributeCategoryId"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="" prop="name">
              <el-input
                v-model="form.name"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="" prop="pic">
              <el-input
                v-model="form.pic"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="货号" prop="productSn">
              <el-input
                v-model="form.productSn"
                placeholder="货号"
              />
            </el-form-item>
            <el-form-item label="删除状态：0->未删除；1->已删除" prop="deleteStatus">
              <el-input
                v-model="form.deleteStatus"
                placeholder="删除状态：0->未删除；1->已删除"
              />
            </el-form-item>
            <el-form-item label="上架状态：0->下架；1->上架" prop="publishStatus">
              <el-input
                v-model="form.publishStatus"
                placeholder="上架状态：0->下架；1->上架"
              />
            </el-form-item>
            <el-form-item label="新品状态:0->不是新品；1->新品" prop="newStatus">
              <el-input
                v-model="form.newStatus"
                placeholder="新品状态:0->不是新品；1->新品"
              />
            </el-form-item>
            <el-form-item label="推荐状态；0->不推荐；1->推荐" prop="recommandStatus">
              <el-input
                v-model="form.recommandStatus"
                placeholder="推荐状态；0->不推荐；1->推荐"
              />
            </el-form-item>
            <el-form-item label="审核状态：0->未审核；1->审核通过" prop="verifyStatus">
              <el-input
                v-model="form.verifyStatus"
                placeholder="审核状态：0->未审核；1->审核通过"
              />
            </el-form-item>
            <el-form-item label="排序" prop="sort">
              <el-input
                v-model="form.sort"
                placeholder="排序"
              />
            </el-form-item>
            <el-form-item label="销量" prop="sale">
              <el-input
                v-model="form.sale"
                placeholder="销量"
              />
            </el-form-item>
            <el-form-item label="" prop="price">
              <el-input
                v-model="form.price"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="促销价格" prop="promotionPrice">
              <el-input
                v-model="form.promotionPrice"
                placeholder="促销价格"
              />
            </el-form-item>
            <el-form-item label="赠送的成长值" prop="giftGrowth">
              <el-input
                v-model="form.giftGrowth"
                placeholder="赠送的成长值"
              />
            </el-form-item>
            <el-form-item label="赠送的积分" prop="giftPoint">
              <el-input
                v-model="form.giftPoint"
                placeholder="赠送的积分"
              />
            </el-form-item>
            <el-form-item label="限制使用的积分数" prop="usePointLimit">
              <el-input
                v-model="form.usePointLimit"
                placeholder="限制使用的积分数"
              />
            </el-form-item>
            <el-form-item label="副标题" prop="subTitle">
              <el-input
                v-model="form.subTitle"
                placeholder="副标题"
              />
            </el-form-item>
            <el-form-item label="商品描述" prop="description">
              <el-input
                v-model="form.description"
                placeholder="商品描述"
              />
            </el-form-item>
            <el-form-item label="市场价" prop="originalPrice">
              <el-input
                v-model="form.originalPrice"
                placeholder="市场价"
              />
            </el-form-item>
            <el-form-item label="库存" prop="stock">
              <el-input
                v-model="form.stock"
                placeholder="库存"
              />
            </el-form-item>
            <el-form-item label="库存预警值" prop="lowStock">
              <el-input
                v-model="form.lowStock"
                placeholder="库存预警值"
              />
            </el-form-item>
            <el-form-item label="单位" prop="unit">
              <el-input
                v-model="form.unit"
                placeholder="单位"
              />
            </el-form-item>
            <el-form-item label="商品重量，默认为克" prop="weight">
              <el-input
                v-model="form.weight"
                placeholder="商品重量，默认为克"
              />
            </el-form-item>
            <el-form-item label="是否为预告商品：0->不是；1->是" prop="previewStatus">
              <el-input
                v-model="form.previewStatus"
                placeholder="是否为预告商品：0->不是；1->是"
              />
            </el-form-item>
            <el-form-item label="以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮" prop="serviceIds">
              <el-input
                v-model="form.serviceIds"
                placeholder="以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮"
              />
            </el-form-item>
            <el-form-item label="" prop="keywords">
              <el-input
                v-model="form.keywords"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="" prop="note">
              <el-input
                v-model="form.note"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="画册图片，连产品图片限制为5张，以逗号分割" prop="albumPics">
              <el-input
                v-model="form.albumPics"
                placeholder="画册图片，连产品图片限制为5张，以逗号分割"
              />
            </el-form-item>
            <el-form-item label="" prop="detailTitle">
              <el-input
                v-model="form.detailTitle"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="" prop="detailDesc">
              <el-input
                v-model="form.detailDesc"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="产品详情网页内容" prop="detailHtml">
              <el-input
                v-model="form.detailHtml"
                placeholder="产品详情网页内容"
              />
            </el-form-item>
            <el-form-item label="移动端网页详情" prop="detailMobileHtml">
              <el-input
                v-model="form.detailMobileHtml"
                placeholder="移动端网页详情"
              />
            </el-form-item>
            <el-form-item label="促销开始时间" prop="promotionStartTime">
              <el-input
                v-model="form.promotionStartTime"
                placeholder="促销开始时间"
              />
            </el-form-item>
            <el-form-item label="促销结束时间" prop="promotionEndTime">
              <el-input
                v-model="form.promotionEndTime"
                placeholder="促销结束时间"
              />
            </el-form-item>
            <el-form-item label="活动限购数量" prop="promotionPerLimit">
              <el-input
                v-model="form.promotionPerLimit"
                placeholder="活动限购数量"
              />
            </el-form-item>
            <el-form-item label="促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购" prop="promotionType">
              <el-input
                v-model="form.promotionType"
                placeholder="促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购"
              />
            </el-form-item>
            <el-form-item label="品牌名称" prop="brandName">
              <el-input
                v-model="form.brandName"
                placeholder="品牌名称"
              />
            </el-form-item>
            <el-form-item label="商品分类名称" prop="productCategoryName">
              <el-input
                v-model="form.productCategoryName"
                placeholder="商品分类名称"
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
import { addPmsProduct, delPmsProduct, getPmsProduct, listPmsProduct, updatePmsProduct } from '@/api/pms/pmsproduct'

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
      pmsproductList: [],

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
      listPmsProduct(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.pmsproductList = response.data
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
        appId: undefined,
        brandId: undefined,
        productCategoryId: undefined,
        feightTemplateId: undefined,
        productAttributeCategoryId: undefined,
        name: undefined,
        pic: undefined,
        productSn: undefined,
        deleteStatus: undefined,
        publishStatus: undefined,
        newStatus: undefined,
        recommandStatus: undefined,
        verifyStatus: undefined,
        sort: undefined,
        sale: undefined,
        price: undefined,
        promotionPrice: undefined,
        giftGrowth: undefined,
        giftPoint: undefined,
        usePointLimit: undefined,
        subTitle: undefined,
        description: undefined,
        originalPrice: undefined,
        stock: undefined,
        lowStock: undefined,
        unit: undefined,
        weight: undefined,
        previewStatus: undefined,
        serviceIds: undefined,
        keywords: undefined,
        note: undefined,
        albumPics: undefined,
        detailTitle: undefined,
        detailDesc: undefined,
        detailHtml: undefined,
        detailMobileHtml: undefined,
        promotionStartTime: undefined,
        promotionEndTime: undefined,
        promotionPerLimit: undefined,
        promotionType: undefined,
        brandName: undefined,
        productCategoryName: undefined
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
      this.title = '添加商品信息'
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
      getPmsProduct(id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改商品信息'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updatePmsProduct(this.form).then(response => {
              if (response.code === 0) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addPmsProduct(this.form).then(response => {
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
      }).then(function() {
        return delPmsProduct(Ids)
      }).then(() => {
        this.getList()
        this.msgSuccess('删除成功')
      }).catch(function() {
      })
    }
  }
}
</script>

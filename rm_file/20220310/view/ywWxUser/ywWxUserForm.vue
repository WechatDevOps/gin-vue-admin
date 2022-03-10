<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label=",-1:">
          <el-input v-model.number="formData.expireTime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="insertTime字段:">
          <el-date-picker v-model="formData.insertTime" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label=",0,1:">
          <el-input v-model.number="formData.isRegistered" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="ins_id:">
          <el-input v-model="formData.lastLoginInsid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="ip:">
          <el-input v-model="formData.lastLoginIp" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="lastLoginTime字段:">
          <el-date-picker v-model="formData.lastLoginTime" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="nameNotes字段:">
          <el-input v-model="formData.nameNotes" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="id:">
          <el-input v-model="formData.openid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="phone字段:">
          <el-input v-model="formData.phone" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="updateTime字段:">
          <el-date-picker v-model="formData.updateTime" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="wxName字段:">
          <el-input v-model="formData.wxName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="wxPhoto字段:">
          <el-input v-model="formData.wxPhoto" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="id:">
          <el-input v-model.number="formData.ywid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'YwWxUser'
}
</script>

<script setup>
import {
  createYwWxUser,
  updateYwWxUser,
  findYwWxUser
} from '@/api/ywWxUser'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        expireTime: 0,
        insertTime: new Date(),
        isRegistered: 0,
        lastLoginInsid: '',
        lastLoginIp: '',
        lastLoginTime: new Date(),
        nameNotes: '',
        openid: '',
        phone: '',
        updateTime: new Date(),
        wxName: '',
        wxPhoto: '',
        ywid: 0,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findYwWxUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reywWxUser
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createYwWxUser(formData.value)
          break
        case 'update':
          res = await updateYwWxUser(formData.value)
          break
        default:
          res = await createYwWxUser(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>

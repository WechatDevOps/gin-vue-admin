<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="id:">
          <el-input v-model="formData.sessionId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="id:">
          <el-input v-model.number="formData.ywid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="loginType字段:">
          <el-input v-model.number="formData.loginType" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="ip:">
          <el-input v-model="formData.srcIp" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="srcPort字段:">
          <el-input v-model.number="formData.srcPort" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="ip:">
          <el-input v-model="formData.ywIp" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="ID:">
          <el-input v-model="formData.insId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="insName字段:">
          <el-input v-model="formData.insName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="url:">
          <el-input v-model="formData.replayUrl" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="insertTime字段:">
          <el-date-picker v-model="formData.insertTime" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="updateTime字段:">
          <el-date-picker v-model="formData.updateTime" type="date" placeholder="选择日期" clearable></el-date-picker>
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
  name: 'YwLoginLog'
}
</script>

<script setup>
import {
  createYwLoginLog,
  updateYwLoginLog,
  findYwLoginLog
} from '@/api/ywLoginLog'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        sessionId: '',
        ywid: 0,
        loginType: 0,
        srcIp: '',
        srcPort: 0,
        ywIp: '',
        insId: '',
        insName: '',
        replayUrl: '',
        insertTime: new Date(),
        updateTime: new Date(),
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findYwLoginLog({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reywLoginLog
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
          res = await createYwLoginLog(formData.value)
          break
        case 'update':
          res = await updateYwLoginLog(formData.value)
          break
        default:
          res = await createYwLoginLog(formData.value)
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

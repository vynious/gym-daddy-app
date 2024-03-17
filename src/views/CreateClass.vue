<!-- admin side -->
<template>
  <div class="background">
    <h1>Create a Class</h1>
    <a-form
      ref="formRef"
      :model="formState"
      :rules="rules"
      :label-col="labelCol"
      :wrapper-col="wrapperCol"
    >
      <a-form-item label="Class name" name="className">
        <a-input v-model:value="formState.className" />
      </a-form-item>
      <a-form-item label="Booking ID" name="bookingId">
        <a-input-number v-model:value="formState.bookingId" />
      </a-form-item>
      <a-form-item label="Activity time" required name="activityTime">
        <a-date-picker
          v-model:value="formState.activityTime"
          show-time
          type="datetime"
          placeholder="Pick a date and time"
          style="width: 100%"
        />
      </a-form-item>
      <a-form-item label="Duration" name="duration">
        <a-input-number v-model:value="formState.duration" />
      </a-form-item>
      <a-form-item label="Suitable level" name="suitableLevel">
        <a-input v-model:value="formState.suitableLevel" />
      </a-form-item>
      <a-form-item label="Max Capacity" name="maxCapacity">
        <a-input-number v-model:value="formState.maxCapacity" />
      </a-form-item>
      <a-form-item label="Activity description" name="activityDescription">
        <a-textarea v-model:value="formState.activityDescription" />
      </a-form-item>
      <a-form-item :wrapper-col="{ span: 14, offset: 5 }">
        <a-button type="primary" @click="onSubmit">Create</a-button>
        <a-button style="margin-left: 10px" @click="resetForm">Reset</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>


<script setup>
import { reactive, ref } from 'vue';
const formRef = ref();
const labelCol = {
  span: 5,
};
const wrapperCol = {
  span: 13,
};
const formState = reactive({
  className: '', // string
  bookingId: 0, // number
  activityTime: undefined, 
  duration: 0, // number
  suitableLevel: '', // string
  maxCapacity: 0,
  activityDescription: '', // string 
});
const rules = {
  className: [
    {
      required: true,
      message: 'Please input class name',
      trigger: 'change',
    },
  ],
  bookingId: [
    {
      required: true,
      message: 'Please input booking ID',
      trigger: 'change',
    },
  ],
  activityTime: [
    {
      required: true,
      message: 'Please pick a date and time',
      trigger: 'change',
      type: 'object',
    },
  ],
  duration: [
    {
      required: true,
      message: 'Please input duration',
      trigger: 'change',
    },
  ],
  suitableLevel: [
    {
      required: true,
      message: 'Please input suitable level',
      trigger: 'change',
    },
  ],
  maxCapacity: [
    {
      required: true,
      message: 'Please input max capacity',
      trigger: 'change',
    },
  ],
  activityDescription: [
    {
      required: true,
      message: 'Please input activity description',
      trigger: 'blur',
    },
  ],
};
const onSubmit = () => {
  formRef.value
    .validate()
    .then(() => {
      console.log('values', formState);
    })
    .catch(error => {
      console.log('error', error);
    });
};
const resetForm = () => {
  formRef.value.resetFields();
};
</script>

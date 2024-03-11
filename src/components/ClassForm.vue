<!-- admin side -->
<template>
    <a-form
      ref="formRef"
      :model="formState"
      :rules="rules"
      :label-col="labelCol"
      :wrapper-col="wrapperCol"
    >
      <a-form-item ref="name" label="Class name" name="name">
        <a-input v-model:value="formState.name" />
      </a-form-item>

      <a-form-item label="Booking ID">
      <a-input-number id="bookingID" v-model:value="bookingID"/>
    </a-form-item>


      <a-form-item label="Activity time" required name="date1">
        <a-date-picker
          v-model:value="formState.date1"
          show-time
          type="date"
          placeholder="Pick a date"
          style="width: 100%"
        />
      </a-form-item>

      <a-form-item label="Duration" name="value">
        <a-input-number id="inputNumber" v-model:value="value" :min="1" :max="50" />
      </a-form-item>

      <!-- <a-form-item label="New class" required name="yes/no">
        <a-radio-group v-model:value="formState.resource">
          <a-radio value="1">Yes</a-radio>
          <a-radio value="2">No</a-radio>
        </a-radio-group>
      </a-form-item> -->

      <a-form-item ref="level" label="Suitable level" name="level">
        <a-input v-model:value="formState.level" />
      </a-form-item>

      <a-form-item label="Activity description" name="desc">
        <a-textarea v-model:value="formState.desc" />
      </a-form-item>

      <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
        <a-button type="primary" @click="onSubmit">Create</a-button>
        <a-button style="margin-left: 10px" @click="resetForm">Reset</a-button>
      </a-form-item>

    </a-form>
  </template>
  <script lang="ts" setup>
  import { Dayjs } from 'dayjs';
  import { reactive, ref, toRaw } from 'vue';
  import type { UnwrapRef } from 'vue';
  import type { Rule } from 'ant-design-vue/es/form';

  const value = ref(0);
  const bookingID = ref(0);

  interface FormState {
    name: string;
    date1: Dayjs | undefined;
    type: string[];
    resource: string;
    desc: string;
    level: string;
  }
  const formRef = ref();
  const labelCol = { span: 5 };
  const wrapperCol = { span: 13 };
  const formState: UnwrapRef<FormState> = reactive({
    name: '',
    value: '',
    date1: undefined,
    type: [],
    resource: '',
    desc: '',
    level: '',
  });
  const rules: Record<string, Rule[]> = {
    name: [
      { required: true, message: 'Please input class name', trigger: 'change' },
    ],
    date1: [{ required: true, message: 'Please pick a date', trigger: 'change', type: 'object' }],
    // desc: [{ required: true, message: 'Please input activity form', trigger: 'blur' }],
    // value:  [{ required: true, message: 'Please input max capacity', trigger: 'blur' }],
  };
  const onSubmit = () => {
    formRef.value
      .validate()
      .then(() => {
        console.log('values', formState, toRaw(formState));
      })
      .catch(error => {
        console.log('error', error);
      });
  };
  const resetForm = () => {
    formRef.value.resetFields();
  };
  </script>
  
  
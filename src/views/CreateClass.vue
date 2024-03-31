<!-- admin side -->
<template>
  <div class="background">
    <h1>Create a Class <router-link to="/classSearch">Class search</router-link></h1>
    <a-form
      ref="formRef"
      :model="formState"
      :rules="rules"
      :label-col="labelCol"
      :wrapper-col="wrapperCol"
    >
      <a-form-item label="Class name" name="className"  style="font-family: 'Poppins Medium';">
        <a-input v-model:value="formState.className" />
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
        <a-button type="primary" @click="onSubmit" style="font-family: 'Poppins Medium';">Create</a-button>
        <a-button style="font-family: 'Poppins Medium'; margin-left: 10px" @click="resetForm" >Reset</a-button>
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

<style>
h1{
  font-family: "Poppins Bold";
}
.background .ant-form-item-label {
  font-family: 'Poppins Medium';
}

@import url(https://db.onlinewebfonts.com/c/0c28006f19928dfd146027cfd7024ca0?family=Poppins+Medium);
@import url(https://db.onlinewebfonts.com/c/07ecc0aa9ce268962dea7356eeff50a6?family=Poppins+Bold);

@font-face {
  font-family: "Poppins Medium";
  src: url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.eot");
  src: url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.eot?#iefix")
      format("embedded-opentype"),
    url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.woff2")
      format("woff2"),
    url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.woff")
      format("woff"),
    url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.ttf")
      format("truetype"),
    url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.svg#Poppins Medium")
      format("svg");
}

@font-face {
  font-family: "Poppins Bold";
  src: url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.eot");
  src: url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.eot?#iefix")
      format("embedded-opentype"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.woff2")
      format("woff2"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.woff")
      format("woff"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.ttf")
      format("truetype"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.svg#Poppins Bold")
      format("svg");
}
</style>
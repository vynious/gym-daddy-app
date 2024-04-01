<!-- admin side -->
<template>
  <div class="background" >
    <a-card style="width: 800px; margin-top: 5%;">
      <h1>Create a Class <router-link to="/classSearch">Class search</router-link></h1>
      <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-item label="Class Id" name="id"  style="font-family: 'Poppins Medium';">
            <a-input v-model:value="formState.id" />
        </a-form-item>
        <a-form-item label="Class name" name="name"  style="font-family: 'Poppins Medium';">
          <a-input v-model:value="formState.name" />
        </a-form-item>
        <a-form-item label="Activity time" required name="date_time">
          <a-date-picker
            v-model:value="formState.date_time"
            show-time
            type="datetime"
            placeholder="Pick a date and time"
            style="width: 100%"
          />
        </a-form-item>
        <a-form-item label="Duration" name="duration">
          <a-input-number v-model:value="formState.duration" />
        </a-form-item>
        <a-form-item label="Suitable level" name="suitable_level">
          <a-input v-model:value="formState.suitable_level" />
        </a-form-item>
        <a-form-item label="Max Capacity" name="max_capacity">
          <a-input-number v-model:value="formState.max_capacity" />
        </a-form-item>
        <a-form-item label="Activity description" name="description">
          <a-textarea v-model:value="formState.description" />
        </a-form-item>
        <a-form-item :wrapper-col="{ span: 14, offset: 5 }">
          <a-button type="primary" @click="onSubmit" style="font-family: 'Poppins Medium';">Create</a-button>
          <a-button style="font-family: 'Poppins Medium'; margin-left: 10px" @click="resetForm" >Reset</a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>


<script setup>
import { reactive, ref } from 'vue';
import axios from 'axios';
const formRef = ref();
const labelCol = {
  span: 5,
};
const wrapperCol = {
  span: 13,
};
const formState = reactive({
  id: 0,
  name: '', // string
  date_time: undefined, 
  duration: 0, // number
  suitable_level: '', // string
  max_capacity: 0,
  description: '', // string 
});
const rules = {
  id: [
    {
      required: true,
      message: 'Please input id',
      trigger: 'change',
    },
  ],
  name: [
    {
      required: true,
      message: 'Please input class name',
      trigger: 'change',
    },
  ],
  date_time: [
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
  suitable_level: [
    {
      required: true,
      message: 'Please input suitable level',
      trigger: 'change',
    },
  ],
  max_capacity: [
    {
      required: true,
      message: 'Please input max capacity',
      trigger: 'change',
    },
  ],
  description: [
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
      axios.post(process.env.VUE_APP_CREATE_CLASS_URL, formState)
        .then(response => {
          console.log('Class created successfully', response.data);
          resetForm();
        })
        .catch(error => {
          console.error('Error creating class:', error);
          // Handle error, show message to user, etc.
        });
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

  .background {
        background-image: url("../assets/background.png");
        background-size: cover;
        background-position: center;
        height: 100vh;
        display: flex;
        flex-direction: column; /* Align items in a column */
        justify-content: flex-start; /* Start from the top */
        align-items: center;
        font-family: "Poppins Medium";
    }
h1{
  font-family: "Poppins Bold";
}
.background .ant-form-item-label {
  font-family: 'Poppins Medium';
}


h1{
  font-family: "Poppins Bold";
}
.background .ant-form-item-label {
  font-family: 'Poppins Medium';
}


</style>
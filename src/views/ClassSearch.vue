<template>
    <div style="font-family: 'Poppins Medium';">
      <a-table :columns="columns" :data-source="data" @change="onChange" style="font-family: 'Poppins Medium';"></a-table>
    </div>
  </template>
  
  <script setup>
  const columns = [
    {
      title: 'Class',
      dataIndex: 'name',
      width: '20%',
      filters: [
        { text: 'Pilates', value: 'Pilates' },
        { text: 'Yoga', value: 'Yoga' },
        { text: 'Dance', value: 'Dance' },
      ],
      onFilter: (value, record) => record.name === value,
    },
    {
      title: 'Date',
      dataIndex: 'date',
      width: '15%',
      sorter: (a, b) => new Date(a.date) - new Date(b.date),
    },
    {
      title: 'Time',
      dataIndex: 'time',
      width: '15%',
      sorter: (a, b) => new Date('1970-01-01T' + a.time) - new Date('1970-01-01T' + b.time),
    },
    {
      title: 'Enrolled',
      dataIndex: 'enrolled',
      width: '10%',
      sorter: (a, b) => a.enrolled - b.enrolled,
    },
    {
      title: 'Max Capacity',
      dataIndex: 'maxCapacity',
      width: '10%',
      sorter: (a, b) => a.maxCapacity - b.maxCapacity,
    },
    {
      title: 'Status',
      dataIndex: 'status',
      width: '20%',
      filters: [
        { text: 'Ongoing', value: 'Ongoing' },
        { text: 'Impending', value: 'Impending' },
        { text: 'Completed', value: 'Completed' },
      ],
      onFilter: (value, record) => record.status === value,
    },
  ];
  
  const data = [
    {
      key: '1',
      name: 'Yoga',
      date: '2024-03-20',
      time: '10:00 AM',
      enrolled: 15,
      maxCapacity: 20,
      status: 'Ongoing',
    },
    {
      key: '2',
      name: 'Pilates',
      date: '2024-03-22',
      time: '2:00 PM',
      enrolled: 18,
      maxCapacity: 25,
      status: 'Impending',
    },
    {
      key: '3',
      name: 'Dance',
      date: '2024-03-11',
      time: '3:00 PM',
      enrolled: 25,
      maxCapacity: 25,
      status: 'Completed',
    },
    // need to research how to link to backend 
  ];
  
  function onChange(pagination, filters, sorter, extra) {
    console.log('params', pagination, filters, sorter, extra);
  }
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
</style>
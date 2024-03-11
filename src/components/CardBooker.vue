<template>
    <div style="padding: 20px">
      <a-row :gutter="16">
        <a-col :span="8">
          <a-card title="Yoga" :bordered="false">
            <a-image :width="300" src="/images/yoga.jpeg"/>
            <p>A 1h session, this class will introduce the fundamentals of a yoga practice, including breath, proper alignment, and mindfulness practices.</p>
            <b>Difficulty: EASY</b>
            <br>
            <router-link to="/bookyoga">
              <a-button type="primary">Book Now</a-button>
            </router-link>
          </a-card>
        </a-col>
        <a-col :span="8">
          <a-card title="Dance" :bordered="false">
            <a-image :width="300" src="../images/dance.jpeg"/>
            <p>A 1h session where instructors will share their tips and tricks accumulated from experience to teach the trendiest dance moves</p>
            <b>Difficulty: INTERMEDIATE</b>
            <br>
            <a-button type="primary">Book Now</a-button>
          </a-card>
        </a-col>
        <a-col :span="8">
          <a-card title="Pilates" :bordered="false">
            <a-image :width="300" src="../images/pilates.jpeg"/>
            <p>A 1h session where each class will work towards balancing all muscle groups’ strength and flexibility, stimulating the core muscles.</p>
            <b>Difficulty: HARD</b>
            <br>
            <a-button type="primary">Book Now</a-button>
          </a-card>
        </a-col>
      </a-row>
    </div>
  </template>
  
  
  
<!-- <template>
	<div style="padding: 20px">
		<a-row :gutter="16">
			<a-col :span="8">
				<a-card
					title="Yoga"
					:bordered="false"
					style="font-family: 'Poppins Bold'"
				>
					<p style="font-family: 'Poppins Medium'">
						A 1h session, this class will introduce the fundamentals
						of a yoga practice, including breath, proper alignment,
						and mindfulness practices.
					</p>
					<b>Difficulty: EASY</b>
					<br />
					<router-link to="/bookyoga">
						<a-button
							type="primary"
							style="font-family: 'Poppins Medium'"
							>Book Now</a-button
						>
					</router-link>
				</a-card>
			</a-col>
			<a-col :span="8">
				<a-card
					title="Dance"
					:bordered="false"
					style="font-family: 'Poppins Bold'"
				>
					<p style="font-family: 'Poppins Medium'">
						A 1h session where instructors will share their tips and
						tricks accumulated from experience to teach the
						trendiest dance moves.
					</p>
					<b>Difficulty: INTERMEDIATE</b>
					<br />
					<router-link to="/bookdance">
						<a-button
							type="primary"
							style="font-family: 'Poppins Medium'"
							>Book Now</a-button
						>
					</router-link>
				</a-card>
			</a-col>
			<a-col :span="8">
				<a-card
					title="Pilates"
					:bordered="false"
					style="font-family: 'Poppins Bold'"
				>
					<p style="font-family: 'Poppins Medium'">
						A 1h session where each class will work towards
						balancing all muscle groups’ strength and flexibility,
						stimulating the core muscles.
					</p>
					<b>Difficulty: HARD</b>
					<br />
					<router-link to="/bookpilates">
						<a-button
							type="primary"
							style="font-family: 'Poppins Medium'"
							>Book Now</a-button
						>
					</router-link>
				</a-card>
			</a-col>
		</a-row>
	</div>
</template> -->

<template>
	<div style="padding: 20px">
		<a-row :gutter="16">
			<a-col :span="8" v-for="classItem in classes" :key="classItem.id">
				<a-card :bordered="false" style="font-family: 'Poppins Bold'">
					<div class="card-title" style="font-family: 'Poppins Bold'">
						{{ classItem.name }}
					</div>
					<p style="font-family: 'Poppins Medium'">
						{{ classItem.description }}
					</p>
					<b
						>Difficulty:
						{{ classItem.suitable_level.toUpperCase() }}</b
					>
					<br />
					<router-link
						:to="`/book${classItem.name
							.toLowerCase()
							.replace(/\s+/g, '')}`"
					>
						<a-button
							type="primary"
							style="font-family: 'Poppins Medium'"
							>Book Now</a-button
						>
					</router-link>
				</a-card>
			</a-col>
		</a-row>
	</div>
</template>

<script>
import axios from "axios";

export default {
	data() {
		return {
			classes: [],
		};
	},
	created() {
		this.fetchClasses();
	},
	methods: {
		fetchClasses() {
			const baseURL = "http://localhost:8000/api/classes";
			axios
				.get(baseURL)
				.then((response) => {
					// Assuming the server response directly contains an array of class objects
					this.classes = response.data.classes;
				})
				.catch((error) => {
					console.error(
						"There was an error fetching the classes data:",
						error
					);
				});
		},
	},
};
</script>

<style scoped>
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
a-card .ant-card-head-title {
	font-family: "Poppins Bold";
}
</style>

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

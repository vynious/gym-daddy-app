<template>
	<div style="padding: 20px">
		<a-row :gutter="16">
			<a-col :span="8" v-for="classItem in classes" :key="classItem.id">
				<a-card
					:bordered="false"
					style="font-family: 'Poppins Bold'; margin-bottom: 10px"
				>
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
					<b>Date and Time: {{ classItem.date_time }}</b>
					<br />
					<b>Current Capacity: {{ classItem.capacity }}</b>
					<br />
					<b>Max Capacity: {{ classItem.max_capacity }}</b>
					<br />
					<a-button
						type="primary"
						style="font-family: 'Poppins Medium'; margin-right: 5px"
						@click="bookClass(classItem.id)"
					>
						Book Now
					</a-button>
					<a-button
						v-if="isAdmin"
						type="danger"
						style="
							font-family: 'Poppins Medium';
							background-color: #f5222d;
							color: white;
						"
						@click="cancelClass(classItem.id)"
					>
						Delete
					</a-button>
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
			userId: 1, // This should be dynamically set based on the logged-in user
			isAdmin: false,
		};
	},
	created() {
		this.fetchClasses();
		this.checkAdminStatus();
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
		checkAdminStatus() {
			this.isAdmin = localStorage.getItem("isAdmin") === "true";
		},
		bookClass(classId) {
			const bookingURL = "http://localhost:8000/api/booking";
			const bookingData = {
				user_id: this.userId,
				class_id: classId,
			};

			axios
				.post(bookingURL, bookingData)
				.then((response) => {
					console.log(response.data);
					alert("Booking successful!");
				})
				.catch((error) => {
					console.error(
						"There was an error creating the booking:",
						error
					);
					alert("Booking failed. Please try again.");
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

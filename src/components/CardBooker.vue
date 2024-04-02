<template>
	<div style="padding: 20px">
		<a-row :gutter="16">
			<a-col :span="8" v-for="classItem in classes" :key="classItem.id">
				<a-card :bordered="false" style="font-family: 'Poppins Bold'; margin-bottom: 10px">
					<div class="card-title" style="font-family: 'Poppins Bold'">
						{{ classItem.name }}
					</div>
					<p style="font-family: 'Poppins Medium'">
						{{ classItem.description }}
					</p>
					<b>Difficulty:
						{{ classItem.suitable_level.toUpperCase() }}</b>
					<br />
					<b>Date and Time: {{ classItem.date_time }}</b>
					<br />
					<b>Current Capacity: {{ classItem.capacity }}</b>
					<br />
					<b>Max Capacity: {{ classItem.max_capacity }}</b>
					<br />
					<a-button type="primary" style="font-family: 'Poppins Medium'; margin-right: 5px"
						:disabled="isClassBooked(classItem.id)" @click="bookClass(classItem.id)">
						{{ isClassBooked(classItem.id) ? 'Booked!' : 'Book Now' }}
					</a-button>
					<a-button v-if="isAdmin" type="danger" style="
							font-family: 'Poppins Medium';
							background-color: #f5222d;
							color: white;
						" @click="deleteClass(classItem.id)">
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
			isAdmin: false,
			userId: null, // This should be dynamically set based on the logged-in user
			bookedClasses: []
		};
	},
	created() {
		this.fetchUserId();
		this.fetchClasses();
		this.checkAdminStatus();
		this.fetchBookedClasses();
	},
	computed: {
		isClassBooked() {
			return (classId) => this.bookedClasses.includes(classId.toString());
		},
	},
	methods: {
		fetchUserId() {
			// This should be dynamically set based on the logged-in user
			const userId = localStorage.getItem("user_id");
			this.userId = userId ? JSON.parse(userId) : null;
		},
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
				class_id: classId.toString(),
			};

			axios
				.post(bookingURL, bookingData)
				.then((response) => {
					console.log(response.data);
					alert("Booking successful!");
					// Find the booked class in your classes array and decrement its capacity
					const bookedClass = this.classes.find(c => c.id === classId);
					if (bookedClass) {
						// Directly decrement the capacity of the class
						bookedClass.capacity--;
						// Vue should react to this change if 'classes' is a reactive data property
					}
					console.log(this.bookedClasses);
					this.bookedClasses.push(classId.toString());
				})
				.catch((error) => {
					console.error(
						"There was an error creating the booking:",
						error
					);
					alert("Booking failed. Please try again.");
				});
		},
		deleteClass(classId) {
			const deleteURL =
				"http://localhost:8000/api/classes/" + classId.toString();

			axios
				.delete(deleteURL)
				.then((response) => {
					console.log(response.data);
					alert("delete successful!");
				})
				.catch((error) => {
					console.error(
						"There was an error deleting the class:",
						error
					);
					alert("delete failed. Please try again.");
				});
		},
		fetchBookedClasses() {
			const bookingURL = `http://localhost:8000/api/booking/user/${this.userId}`;
			axios.get(bookingURL, { headers: { Authorisation: `Bearer ${JSON.parse(localStorage.getItem("token"))}` } })
				.then((response) => {
					this.bookedClasses = response.data.bookings.map(booking => booking.class_id);
				})
				.catch((error) => {
					console.error(
						"There was an error fetching the booked classes data:", error
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
	src: url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.eot?#iefix") format("embedded-opentype"),
		url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.woff2") format("woff2"),
		url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.woff") format("woff"),
		url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.ttf") format("truetype"),
		url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.svg#Poppins Medium") format("svg");
}

@font-face {
	font-family: "Poppins Bold";
	src: url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.eot");
	src: url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.eot?#iefix") format("embedded-opentype"),
		url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.woff2") format("woff2"),
		url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.woff") format("woff"),
		url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.ttf") format("truetype"),
		url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.svg#Poppins Bold") format("svg");
}

a-card .ant-card-head-title {
	font-family: "Poppins Bold";
}
</style>

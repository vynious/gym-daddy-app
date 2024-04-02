<template>
    <div class="default">
        <div class="bookingsummary">
            <div>
                <!-- Dynamically generate booking sections -->
                <div class="bookingsect" v-for="(booking, index) in bookings" :key="index">
                    <button class="booking" @click="showDetails(booking)">
                        <p>Booking Number&nbsp;&nbsp;#{{ booking.id }}</p>
                        <p>Date&nbsp;&nbsp;:&nbsp;&nbsp;{{ formatDate(booking.class.date_time) }}</p>
                        <p>Time&nbsp;&nbsp;:&nbsp;&nbsp;{{ formatTime(booking.class.date_time) }}</p>
                        <p class="details">See Details</p>
                    </button>
                    <button class="delete" @click="deleteBooking(booking.id, index)">
                        <img src="../assets/delete.png" alt="">
                    </button>
                </div>
            </div>
            <!-- Modal structure -->
            <div class="modal" v-if="selectedBooking">
                <div class="modal-content">
                    <span class="close" @click="selectedBooking = null">&times;</span>
                    <h2>Booking Details</h2>
                    <p><strong>Class:</strong> {{ selectedBooking.class.name }}</p>
                    <p><strong>Date:</strong> {{ formatDate(selectedBooking.class.date_time) }}</p>
                    <p><strong>Time:</strong> {{ formatTime(selectedBooking.class.date_time) }}</p>
                    <p><strong>Duration:</strong> {{ selectedBooking.class.duration }} minutes</p>
                    <p><strong>Capacity:</strong> {{ selectedBooking.class.capacity }}</p>
                    <p><strong>Max Capacity:</strong> {{ selectedBooking.class.max_capacity }}</p>
                    <p><strong>Level:</strong> {{ selectedBooking.class.suitable_level }}</p>
                </div>
            </div>
        </div>

    </div>
</template>

<script>
import axios from 'axios';

export default {
    name: 'ProfileBookings',
    data() {
        return {
            bookings: [],
            selectedBooking: null
        }
    },
    created() {
        this.fetchBookings();
    },
    methods: {
        fetchBookings() {
            // The URL needs to be adjusted according to your API endpoint and user ID
            const userId = JSON.parse(localStorage.getItem("user_id"));
            const url = `http://localhost:8000/api/booking/user/${userId}`;
            console.log(url);

            axios.get(url, {
                headers: {
                    Authorisation: `Bearer ${JSON.parse(localStorage.getItem('token'))}`
                }
            })
                .then(response => {
                    this.bookings = response.data.bookings; // Store the bookings in the component's data
                    console.log('Bookings:', this.bookings);
                })
                .catch(error => {
                    console.error('There was an error fetching the bookings:', error);
                });
        },
        formatDate(dateTimeString) {
            const date = new Date(dateTimeString);
            return date.toLocaleDateString('en-SG');
        },

        formatTime(dateTimeString) {
            const date = new Date(dateTimeString);
            return date.toLocaleTimeString('en-SG', { hour: '2-digit', minute: '2-digit' });
        },
        showDetails(booking) {
            this.selectedBooking = booking;
            console.log('Selected booking:', this.selectedBooking);
        },
        deleteBooking(bookingId, index) {
            axios.delete(`http://localhost:8000/api/booking/${bookingId}`, {
                headers: {
                    Authorisation: `Bearer ${JSON.parse(localStorage.getItem('token'))}`
                }
            }).then(() => {
                this.bookings.splice(index, 1); // Remove the booking from the array
            }).catch(error => {
                console.error('Error deleting booking:', error);
            });
        }

    }
}
</script>

<style scoped>
.default {
    background: black;
    height: 100vh;
    width: 70%;
    display: flex;
    justify-content: center;
    align-items: center;
}

.bookingsummary {
    width: 90%;
    height: 90vh;
    overflow-y: auto;
    background-color: #D9D9D920;
    color: white;
    display: flex;
    justify-content: center;
    font-size: 28px;
}

.bookingsect {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 50px;
}

.booking {
    background-color: #D9D9D920;
    border-radius: 12px;
    width: 700px;
    color: white;
    padding: 15px 40px;
    margin-right: 30px;
    font-size: 28px;
    /* font-family: 'Poppins'; */
    border: 0px;
    border-radius: 12px;
}

p {
    margin: 3px;
}

.details {
    font-size: 20px;
    text-decoration: underline;
    cursor: pointer;
}

img {
    width: 70px;
    padding-left: auto;
    padding-right: auto;
}

.delete {
    background-color: #D9D9D920;
    border-radius: 12px;
    width: 100px;
    height: 100px;
    color: white;
    padding: 15px 40px;
    margin: 15px;
    border: 0px;
    border-radius: 12px;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
}

.delete:hover {
    /* background-color: #FFFFFF50; */
    background-color: #00000050;

}

.modal {
    display: flex; /* Added to center the modal-content */
    align-items: center; /* Added to center the modal-content */
    justify-content: center; /* Added to center the modal-content */
    position: fixed;
    z-index: 1000; /* Increased z-index */
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: hidden; /* Changed from auto to hidden */
    background-color: rgba(0, 0, 0, 0.75); /* Made the background darker for visibility */
}

.modal-content {
    background-color: #fefefe;
    margin: auto; /* Centered the modal-content */
    padding: 20px;
    border: 1px solid #888;
    width: 50%; /* Adjust the width as needed */
    color: black;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); /* Optional: added shadow for visibility */
    animation: fadeIn 0.3s; /* Optional: added animation for a fade-in effect */
}

@keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
}


.close {
    color: #aaa;
    float: right;
    font-size: 28px;
    font-weight: bold;
}

.close:hover,
.close:focus {
    color: black;
    text-decoration: none;
    cursor: pointer;
}
</style>
<template>
  <div class="captcha-container">
    <div class="flex flex-row justify-between mb-3">
      <div v-for="icon in icons" :key="icon.id" >
        <img :src="icon.image" :alt="icon.name" width="30px" height="30px"/>
      </div>
    </div>
    <div class="captcha-image">
      <img src="https://via.placeholder.com/300x200" alt="captcha" class="" />
      <!-- Loop through icons -->
      <transition-group name="icon-transition" tag="div">
        <div v-for="icon in icons" :key="icon.id" class="icon"
          :style="{ top: icon.top, left: icon.left, backgroundColor: selectedOrder.includes(icon.order) ? '#d1ffd1' : 'rgba(255, 255, 255, 0.8)' }"
          @click="selectIcon(icon.order)">
          <img :src="icon.image" :alt="icon.name" />
        </div>
      </transition-group>
    </div>
    <p class="message">{{ message }}</p>
    <button class="apply-button" @click="validateOrder">Apply</button>
  </div>
</template>

<script setup>
import { ref, watch } from "vue";

// List of icons with positions and details
const icons = ref([
  { id: 1, order: "1", name: "alarm", top: "10%", left: "10%", image: "https://img.icons8.com/color/48/000000/alarm.png" },
  { id: 2, order: "2", name: "calendar", top: "40%", left: "50%", image: "https://img.icons8.com/color/48/000000/calendar.png" },
  { id: 3, order: "3", name: "cart", top: "70%", left: "80%", image: "https://img.icons8.com/color/48/000000/shopping-cart.png" },
  { id: 4, order: "4", name: "camera", top: "20%", left: "80%", image: "https://img.icons8.com/color/48/000000/camera.png" },
  { id: 5, order: "5", name: "music", top: "60%", left: "20%", image: "https://img.icons8.com/color/48/000000/music.png" },
]);

const selectedOrder = ref([]);
const message = ref("Select the icons in the correct order.");

// Handle selecting an icon
const selectIcon = (order) => {
  if (!selectedOrder.value.includes(order)) {
    selectedOrder.value.push(order);
    message.value = `You selected ${selectedOrder.value.length} icon(s).`;
  } else {
    message.value = "Icon already selected!";
  }
};

// Validate the order of selected icons
const validateOrder = () => {
  const correctOrder = ["1", "2", "3", "4", "5"];
  if (JSON.stringify(selectedOrder.value) === JSON.stringify(correctOrder)) {
    message.value = "Verification successful! 🎉";
  } else {
    message.value = "Incorrect order! Resetting...";
    resetSelection();
  }
};

// Reset selected icons
const resetSelection = () => {
  setTimeout(() => {
    selectedOrder.value = [];
    message.value = "Select the icons in the correct order.";
  }, 1000); // Delay reset for better feedback
};
</script>

<style scoped>
body {
  font-family: Arial, sans-serif;
  background-color: #f0f0f0;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  margin: 0;
}

.captcha-container {
  /*position: relative;*/
  width: 300px;
  text-align: center;
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.captcha-image {
  position: relative;
  width: 100%;
  height: auto;
  border-radius: 8px;
}

.icon {
  position: absolute;
  width: 50px;
  height: 50px;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  transition: transform 0.2s;
}

.icon img {
  width: 30px;
  height: 30px;
}

.icon:hover {
  transform: scale(1.1);
}

.apply-button {
  margin-top: 20px;
  background-color: #007bff;
  color: #fff;
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s;
}

.apply-button:hover {
  background-color: #0056b3;
}

.message {
  margin: 15px 0;
  font-size: 14px;
  color: #333;
  font-weight: bold;
}

.icon-transition-enter-active,
.icon-transition-leave-active {
  transition: all 0.5s;
}

.icon-transition-enter-from {
  transform: scale(0);
  opacity: 0;
}

.icon-transition-enter-to {
  transform: scale(1);
  opacity: 1;
}

.icon-transition-leave-from {
  transform: scale(1);
  opacity: 1;
}

.icon-transition-leave-to {
  transform: scale(0);
  opacity: 0;
}
</style>
<template>
    <div class="flex flex-col items-center justify-center h-screen bg-gray-100">
      <h1 class="text-2xl font-bold text-gray-800 mb-4">Generate and Save QR Code</h1>
      <!-- QR Code -->
      <qrcode-vue
        :value="text"
        :size="200"
        class="shadow-lg"
        ref="qrcodeRef"
      />
      <!-- Input for QR Code -->
      <input
        v-model="text"
        type="text"
        placeholder="Enter text for QR code"
        class="mt-4 p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
      <!-- Save as PNG Button -->
      <button
        @click="saveQRCode('png')"
        class="mt-4 bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
      >
        Save as PNG
      </button>
      <!-- Save as JPEG Button -->
      <button
        @click="saveQRCode('jpeg')"
        class="mt-2 bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:ring-2 focus:ring-green-400"
      >
        Save as JPEG
      </button>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  import QrcodeVue from 'qrcode.vue';
  
  const text = ref('Hello, QR Code!'); // Default text for QR code
  const qrcodeRef = ref(null); // Reference to the QR Code component
  
  const saveQRCode = (format) => {
    // Access the QR Code canvas
    const canvas = qrcodeRef.value?.$el.querySelector('canvas');
    
    if (!canvas) {
      alert('QR Code not found or failed to generate!');
      return;
    }
  
    // Convert canvas to the desired format and create a download link
    const quality = 1.0; // Set image quality for JPEG (1.0 = highest)
    const imageType = format === 'jpeg' ? 'image/jpeg' : 'image/png';
    const image = canvas.toDataURL(imageType, quality);
  
    // Create a download link
    const link = document.createElement('a');
    link.href = image;
  
    // Set filename based on format
    link.download = `qrcode.${format}`;
    link.click();
  };


  
  
  
  
  
  
  </script>
  
  <style scoped>
  /* Optional: Add custom styling here */
  </style>
  
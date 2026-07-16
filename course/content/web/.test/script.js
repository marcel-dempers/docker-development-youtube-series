console.log('Hello from JavaScript');

document.addEventListener('DOMContentLoaded', () => {
  const buttonElement = document.getElementById('api-button');
  const messageElement = document.getElementById('message');
  if (buttonElement) {
      buttonElement.addEventListener('click', () => {
        // Fetch data from the API
        fetch('http://localhost:8080/weather')
          .then(response => {
            if (!response.ok) {
              throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json(); // Parse the JSON response
          })
          .then(data => {
            // Display the JSON content in the message element
            messageElement.textContent = JSON.stringify(data, null, 2);
          })
          .catch(error => {
            console.error('Error fetching weather data:', error);
            messageElement.textContent = 'Failed to load weather data.';
          });
      });
  }
});
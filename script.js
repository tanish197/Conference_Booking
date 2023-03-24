// Get references to the form and message elements
const bookingForm = document.getElementById('booking-form');
const messageDiv = document.getElementById('message');

// Add an event listener to the form for when it's submitted
bookingForm.addEventListener('submit', function(event) {
  // Prevent the default form submission behavior
  event.preventDefault();

  // Get references to the form fields
  const firstNameInput = document.getElementById('first-name');
  const lastNameInput = document.getElementById('last-name');
  const emailInput = document.getElementById('email');
  const ticketNumberInput = document.getElementById('ticket-number');

  // Validate the form fields
  if (!firstNameInput.value || !lastNameInput.value || !emailInput.value || !ticketNumberInput.value) {
    messageDiv.textContent = 'Please fill out all the required fields.';
    return;
  }

  // Check if there are enough tickets available
  const remainingTickets = document.getElementById('remaining-tickets').textContent;
  if (ticketNumberInput.value > remainingTickets) {
    messageDiv.textContent = `Sorry, there are only ${remainingTickets} tickets left.`;
    return;
  }

  // Submit the form data to the server
  messageDiv.textContent = 'Booking in progress...';
  const formData = new FormData(bookingForm);
  fetch('https://example.com/bookings', {
    method: 'POST',
    body: formData
  })
  .then(response => response.text())
  .then(data => {
    messageDiv.textContent = data;
    bookingForm.reset();
  })
  .catch(error => {
    messageDiv.textContent = 'An error occurred. Please try again later.';
  });
});

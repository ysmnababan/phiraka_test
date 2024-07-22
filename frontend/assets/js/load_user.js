document.addEventListener('DOMContentLoaded', function() {
    // Fetch and display user data
    fetch('http://localhost:8080/users') // Replace with your API URL
      .then(response => response.json()) // Parse JSON response
      .then(data => {
        const tbody = document.querySelector('#userTable tbody');
        tbody.innerHTML = ''; // Clear existing rows
        var count = 0;
        data.forEach(user => {
            // console.log(user)
            count++
          const row = document.createElement('tr');
  
          row.innerHTML = `
            <td>${count}</td>
            <td>${user.username}</td>
            <td>********</td>
            <td>${new Date(user.create_time).toLocaleString()}</td>
            <td>
              <button class="btn edit-btn" onclick="editUser(${user.id})">Edit</button>
              <button class="btn delete-btn" onclick="deleteUser(${user.id})">Delete</button>
            </td>
          `;
  
          tbody.appendChild(row);
        });
      })
      .catch(error => {
        console.error('Error fetching user data:', error);
      });
  });
  
  // Function to handle editing a user
  function editUser(userId) {
    // Redirect to an edit page or open a modal with user details
    window.location.href = `edit.html?id=${userId}`; // Example redirect
  }
  
  // Function to handle deleting a user
  function deleteUser(userId) {
    if (confirm('Are you sure you want to delete this user?')) {
      fetch(`http://localhost:8080/api/users/${userId}`, {
        method: 'DELETE',
      })
      .then(response => response.json())
      .then(result => {
        console.log('User deleted:', result);
        // Refresh the table after deletion
        location.reload();
      })
      .catch(error => {
        console.error('Error deleting user:', error);
      });
    }
  }
  
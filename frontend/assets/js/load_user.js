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
          row.setAttribute('data-user-id', user.user_id);

          row.innerHTML = `
            <td>${count}</td>
            <td class="username">${user.username}</td>
            <td class="password">********</td>
            <td class="create-time">${new Date(user.create_time).toLocaleString()}</td>
            <td class="actions">
              <button class="btn edit-btn" onclick="editUser(${user.user_id})">Edit</button>
              <button class="btn delete-btn" onclick="deleteUser('${user.username}')">Delete</button>
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
    const row = document.querySelector(`tr[data-user-id='${userId}']`);
    const usernameCell = row.querySelector('.username');
    const passwordCell = row.querySelector('.password');
    const actionsCell = row.querySelector('.actions');
  
    const currentUsername = usernameCell.textContent;
    const currentPassword = ""; // You may want to fetch the actual password if allowed
  
    usernameCell.innerHTML = `<input type="text" class="form-control" value="${currentUsername}" required>`;
    passwordCell.innerHTML = `<input type="password" class="form-control" value="${currentPassword}" placeholder="insert your password" required>`;
    actionsCell.innerHTML = `
    <button class="btn save-btn" onclick="saveUser(${userId})">Save</button>
    <button class="btn cancel-btn" onclick="cancelEdit(${userId})">Cancel</button>
  `;
  }
  
  function cancelEdit(userId) {
    // Refresh the user data to cancel the edit
    location.reload();
  }
  
  function saveUser(userId) {
    const row = document.querySelector(`tr[data-user-id='${userId}']`);
  const usernameInput = row.querySelector('.username input');
  const passwordInput = row.querySelector('.password input');
  const errorMessage = document.querySelector('.error-message');

  const updatedUsername = usernameInput.value.trim();
  const updatedPassword = passwordInput.value.trim();

  // Validate that neither the username nor the password is empty
  if (!updatedUsername) {
    errorMessage.textContent = 'Username cannot be empty.';
    errorMessage.style.display = 'block'; // Show the error message
    return; // Prevent saving if the username is empty
  } else if (!updatedPassword) {
    console.log("password not updated")
    // errorMessage.textContent = 'Password cannot be empty.';
    // errorMessage.style.display = 'block'; // Show the error message
    // return; // Prevent saving if the password is empty
  }else if (updatedPassword.length < 5 || updatedPassword.length > 8) {
    errorMessage.textContent = 'Password must be 5-8 characters long.';
    errorMessage.style.display = 'block'; // Show the error message
    return; // Prevent saving if the password length is not between 5 and 8 characters
  } else {
    errorMessage.style.display = 'none'; // Hide the error message if both fields are valid
  }
  
    const formData = {
      user_id: userId,
      username: updatedUsername,
      password: updatedPassword
    };
  
    fetch(`http://localhost:8080/user`, {
      method: 'PUT', // Assuming your API uses PUT for updates
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    })
    .then(response => response.json())
    .then(result => {
      console.log('User updated:', result);
      if (result === "USER UPDATED SUCCESSFULLY") {
        location.reload(); // Refresh the table after updating
      } else {
        console.error('Error updating user:', result);
      }
    })
    .catch(error => {
      console.error('Error updating user:', error);
    });
  }
  
  // Function to handle deleting a user
  function deleteUser(username) {
    console.log(username)
    if (confirm('Are you sure you want to delete this user?')) {
        var formData = {
            username: username,
          };
      fetch(`http://localhost:8080/user`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(formData),
      })
      .then(response => response.json())
      .then(result => {
          console.log('User deleted:', result);
        if (result == "USER DELETED SUCCESSFULLY"){
            // Refresh the table after deletion
            const userData = sessionStorage.getItem('user');
            const user = JSON.parse(userData);
            if (username == user.username){
                // Clear session storage
                sessionStorage.clear();

                // Clear local storage (if used)
                localStorage.clear();

                // Redirect to the login page (in case server logout is not implemented)
                window.location.href = 'index.html';
            } else {
                location.reload();
            }
        }
      })
      .catch(error => {
        console.error('Error deleting user:', error);
      });
    }
  }
  
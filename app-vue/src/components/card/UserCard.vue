<template>
  <div class="user-card">
    <div class="user-info">
      <h3>{{ user.name }}</h3>
      <p><strong>ID:</strong> {{ user.id }}</p>
      <p><strong>Login:</strong> {{ user.login }}</p>
      <p><strong>First Name:</strong> {{ user.firstName }}</p>
      <p><strong>Last Name:</strong> {{ user.lastName }}</p>
      <p><strong>Age:</strong> {{ user.age }}</p>
    </div>
    <div class="user-actions">
      <button class="edit-button" @click="editUser">Edit</button>
      <button class="delete-button" @click="deleteUser">Delete</button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'UserCard',
  props: {
    user: {
      type: Object,
      required: true,
    },
  },
  emits: ['edit', 'delete'],
  setup(props, { emit }) {
    const editUser = () => {
      emit('edit', props.user);
    };

    const deleteUser = () => {
      emit('delete', props.user);
    };

    return {
      editUser,
      deleteUser,
    };
  },
});
</script>

<style scoped>
.user-card {
  border: 1px solid #ddd;
  border-radius: 10px; /* Rounded corners */
  padding: 20px;
  margin: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); /* Add a subtle shadow */
  background-color: #fff;
  max-width: 300px;
}

.user-info h3 {
  margin: 0 0 10px;
  font-size: 1.5rem;
}

.user-info p {
  margin: 5px 0;
}

.user-actions {
  display: flex;
  justify-content: space-between;
  margin-top: 15px;
}

.edit-button,
.delete-button {
  padding: 8px 16px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 0.9rem;
}

.edit-button {
  background-color: #4caf50; /* Green */
  color: white;
}

.edit-button:hover {
  background-color: #45a049;
}

.delete-button {
  background-color: #f44336; /* Red */
  color: white;
}

.delete-button:hover {
  background-color: #e53935;
}
</style>

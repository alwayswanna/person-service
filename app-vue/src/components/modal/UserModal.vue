<template>
  <div class="modal-overlay" v-if="visible">
    <div class="modal">
      <h2>Edit User</h2>
      <form @submit.prevent="saveUser">
        <label for="firstName">First Name:</label>
        <input id="firstName" v-model="editedUser.firstName" />

        <label for="lastName">Last Name:</label>
        <input id="lastName" v-model="editedUser.lastName" />

        <label for="age">Age:</label>
        <input type="number" id="age" v-model="editedUser.age" />

        <label for="login">Login:</label>
        <input id="login" v-model="editedUser.login" />

        <button type="submit">Save</button>
        <button type="button" @click="closeModal">Cancel</button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, watch } from 'vue';

export default defineComponent({
  name: 'EditUserModal',
  props: {
    visible: {
      type: Boolean,
      required: true,
    },
    user: {
      type: Object,
      required: false,
    },
  },
  emits: ['save', 'close'],
  setup(props, { emit }) {
    const editedUser = ref({ ...props.user });

    watch(
      () => props.user,
      (newUser) => {
        editedUser.value = { ...newUser };
      }
    );

    const saveUser = () => {
      emit('save', editedUser.value);
    };

    const closeModal = () => {
      emit('close');
    };

    return {
      editedUser,
      saveUser,
      closeModal,
    };
  },
});
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal {
  background-color: white;
  padding: 20px;
  border-radius: 10px;
  width: 300px;
}

.modal h2 {
  margin-top: 0;
}

.modal form {
  display: flex;
  flex-direction: column;
}

.modal label {
  margin-bottom: 5px;
}

.modal input {
  margin-bottom: 10px;
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 5px;
}

.modal button {
  margin-top: 10px;
  padding: 8px 16px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.modal button[type='submit'] {
  background-color: #4caf50;
  color: white;
}

.modal button[type='button'] {
  background-color: #f44336;
  color: white;
}
</style>

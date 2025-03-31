<template>
  <div class="user-list">
    <h1>User List</h1>
    <div v-show="!isAuthenticated">
      <p>You should login to retrieve users.</p>
    </div>
    <div v-if="loading">Loading...</div>
    <div v-else>
      <div v-if="users.length === 0 && isAuthenticated">No users found.</div>
      <div v-else class="user-grid">
        <UserCard
          v-for="user in users"
          :key="user.id"
          :user="user"
          @edit="openEditUserModal"
          @delete="handleDelete"
        />
      </div>
    </div>

    <div class="button-container" v-if="isAuthenticated && users.length > 0">
      <button class="add-user-button" @click="openAddUserModal">
        Add New User
      </button>
    </div>

    <EditUserModal
      :visible="isEditModalVisible"
      :user="selectedUser"
      @save="handleEdit"
      @close="closeEditUserModal"
    />

    <EditUserModal
      :visible="isAddModalVisible"
      @save="handleAdd"
      @close="closeAddUserModal"
    />
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, ref } from 'vue'
import UserCard from '@/components/card/UserCard.vue'
import { personApi } from '@/api/person-api'
import { useAuthStore } from '@/store/AuthenticationStore'
import EditUserModal from '@/components/modal/UserModal.vue'

export default defineComponent({
  name: 'UsersView',
  components: {
    UserCard,
    EditUserModal
  },
  setup() {
    const users = ref([]);
    const loading = ref(true);
    const isEditModalVisible = ref(false);
    const isAddModalVisible = ref(false);
    const selectedUser = ref(null);

    const authenticationStore = useAuthStore()

    const fetchUsers = async () => {
      try {
        const response = await personApi.loadAllPersons();
        users.value = response.data;
      } catch (error) {
        console.error('Failed to fetch users:', error);
      } finally {
        loading.value = false;
      }
    };

    const openEditUserModal = (user) => {
      selectedUser.value = user;
      isEditModalVisible.value = true;
    };

    const openAddUserModal = () => {
      isAddModalVisible.value = true;
    }

    const closeEditUserModal = () => {
      isEditModalVisible.value = false;
    };

    const closeAddUserModal = () => {
      isAddModalVisible.value = false;
    }

    const isAuthenticated = computed(() => {
      return authenticationStore.accessToken != null && authenticationStore.accessToken !== "";
    })

    const handleDelete = async (user) => {
      try {
        let response = await personApi.deletePerson(user.id);
        if (response.status == 200) {
          users.value = users.value.filter((u) => u.id !== user.id);
        }
      } catch (error) {
        console.error('Failed to delete user:', error);
      }
    };

    const handleEdit = async (user) => {
      try {
        let response = await personApi.editPerson(user);

        if (response.status == 200) {
          const index = users.value.findIndex((u) => u.id === user.id);

          if (index !== -1) {
            users.value[index] = user;
          }
        }

        closeEditUserModal();
      } catch (error) {
        console.error('Failed to update user:', error);
        closeEditUserModal();
      }
    };

    const handleAdd = async (user) => {
      try {
        let response = await personApi.createPerson(user)

        if (response.status == 200) {
          users.value.push(response.data)
        }

        closeAddUserModal();
      } catch (err) {
        console.error('Failed to create user:', err)
        closeAddUserModal()
      }
    }

    onMounted(() => {
      if (authenticationStore.accessToken != null && authenticationStore.accessToken !== "") {
        fetchUsers();
      }
    });

    return {
      users,
      loading,
      handleAdd,
      handleEdit,
      selectedUser,
      handleDelete,
      isAuthenticated,
      openAddUserModal,
      openEditUserModal,
      isAddModalVisible,
      closeAddUserModal,
      closeEditUserModal,
      isEditModalVisible,
    };
  },
});
</script>

<style scoped>
.user-list {
  padding: 20px;
}

.user-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.button-container {
  display: flex;
  justify-content: center;
  width: 100%;
  margin-top: 20px;
}

.add-user-button {
  background-color: #4CAF50;
  color: white;
  border: none;
  padding: 12px 24px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 10px 2px;
  cursor: pointer;
  border-radius: 25px;
  transition: all 0.3s ease;
  box-shadow: 0 2px 5px rgba(0,0,0,0.2);
}
</style>

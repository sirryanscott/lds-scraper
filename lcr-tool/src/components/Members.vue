<template>
  <div class="members">
    <h1>Member List</h1>
    <h2 class="messages" v-if="loading">Loading...</h2>
    <h2 class="messages" v-if="permissionErr && !loading">User doesn't have LCR access<br> or try logging in again</h2>
    <div id="list">
        <tr v-if="!loading && !permissionErr">
            <th>Last Name</th>
            <th>First Name</th>
            <th>Gender</th>
            <th>Age</th>
            <th>Birthday</th>
            <th>Phone</th>
            <th>Email</th>
            <th>Address</th>
            <th>Ward</th>
        </tr>
        <tr v-for="member in members" :key="member.id">
            <!-- name, gender age birthday phone email address -->
            <td>{{ member.nameFamilyPreferredLocal }}</td>
            <td>{{ member.nameGivenPreferredLocal }}</td>
            <td>{{ member.sex }}</td>
            <td>{{ member.age }}</td>
            <td>{{ member.birth.date.display }}</td>
            <td>{{ member.phoneNumber }}</td>
            <td>{{ member.email }}</td>
            <td>{{ member.address.addressLines[0] }} {{ member.address.addressLines[1] }}</td>
            <td>{{ member.unitName }}</td>
        </tr>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Members',
  data () {
    return {
      members: [],
      permissionErr: true,
      loading: true
    }
  },
  created () {
    this.getMembers()
  },
  methods: {
    getMembers: function () {
      var params = new URLSearchParams()

      params.append('username', this.$route.params.username)
      axios.post('/members/', params).then((response) => {
        console.log('Getting member list')
        if (response.status === 200) {
          this.permissionErr = false
          this.loading = false
          this.members = response.data
          console.log(this.members)
        }
      }).catch((error) => {
        if (error.response.status === 403) {
          console.log('User doesn\'t have LCR access')
          this.loading = false
          this.permissionErr = true
          return
        }
        console.log('An unexpected error occured')
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
  text-align: center;
}
ul {
  list-style-type: none;
  padding: 0;
}

table {
}
table, th, td {
  border: 1px solid black;
}

#list {
    border-collapse: collapse;
    margin: 0px auto;
    width: 82%;
    padding-bottom: 80px;
}

td {
    padding: 5px;
}
th {
    padding: 10px;
}
tr:nth-child(even) {
    background-color: #f2f2f2;
}

.messages {
  padding-top: 100px;
  font-style: italic;
}
</style>

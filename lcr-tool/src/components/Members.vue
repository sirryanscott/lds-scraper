<template>
  <div class="members">
    <h1>Member List</h1>
    <div id="table">
    <tr>
        <th>Name</th>
        <th>Gender</th>
        <th>Age</th>
        <th>Birthday</th>
        <th>Phone</th>
        <th>Email</th>
        <th>Address</th>
    </tr>
    <tr v-for="member in members" :key="member.id">
        <!-- name, gender age birthday phone email address -->
        <td>{{ member.nameListPreferredLocal }}</td>
        <td>{{ member.sex }}</td>
        <td>{{ member.age }}</td>
        <td>{{ member.birth.date.display }}</td>
        <td>{{ member.phoneNumber }}</td>
        <td>{{ member.email }}</td>
        <td>{{ member.address.addressLines[0] }} {{ member.address.addressLines[1] }}</td>
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
      permissionErr: false,
      members: []
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
          this.members = response.data
          console.log(this.members)
        }
      }).catch((error) => {
        if (error.response.status === 403) {
          console.log('User doesn\'t have LCR access')
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
.members {
    margin: auto;
    width: 100%;
}

table {
     border-collapse: collapse;
     table-layout: fixed;
     width: 100%;
}

table, th, td {
  border: 1px solid black;
}

#table {
    display: inline-block;
  margin-left: 15%;
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
</style>

<template>
  <div class="login">
    <h1>Login</h1>
    <form @submit.prevent="login">
        Username:<br>
        <input type="text" v-model="username" placeholder="Username">
        <br><br>
        Password:<br>
        <input type="password" v-model="password" placeholder="Password">
        <br><br>
        <input type="submit" value="Submit">
    </form>
    <div v-if="loggedInErr" id="invalid">
        <p>Invalid username or password</p>
    </div>
    <div v-if="permissionErr" id="no-permissions">
        <p>You don't have permission to LCR</p>
     </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Login',
  data () {
    return {
      msg: 'LCR Tool',
      password: '',
      username: '',
      loggedInErr: false,
      permissionErr: false
    }
  },
  methods: {
    login: function () {
      var params = new URLSearchParams()

      params.append('username', this.username)
      params.append('password', this.password)

      axios.post('/login/', params).then((response) => {
        console.log('Logging in')
        if (response.status === 200) {
          this.$router.push({name: 'Members', params: { username: this.username }})
          this.loggedInErr = false
        }
      }).catch((error) => {
        if (error.response.status === 401) {
          console.log('invalid credentials')
          this.loggedInErr = true
          return
        }
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
.login {
    margin-left: 25%;
    margin-right: 25%;
    width: 50%;
}
input {
  width: 95%;
  margin-top:5px;
  padding:3px;
}
form {
    width: 30%;
    margin-left: 35%;
    margin-right: 25%;
    display: inline-block;
}
</style>

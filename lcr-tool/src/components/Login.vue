<template>
  <div class="login">
    <h1>Login</h1>
    <form @submit.prevent="login">
        Username:<br>
        <input type="text" v-model="username" placeholder="Username">
        <br>
        Password:<br>
        <input type="password" v-model="password" placeholder="Password">
        <br><br>
        NEED TO ADD AN IF CHECK BEFORE ROUTING
        <router-link to="/members"><input type="submit" value="Submit"></router-link>
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
}
ul {
  list-style-type: none;
  padding: 0;
}
.login {
    margin: auto;
    width: 50%;
}
</style>

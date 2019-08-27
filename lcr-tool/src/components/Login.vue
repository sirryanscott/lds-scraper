<template>
  <div class="login">
    <h1>Login</h1>
    <form @submit.prevent="login">
        Username:<br>
        <input type="text" name="username" placeholder="Username">
        <br>
        Password:<br>
        <input type="password" name="lastname" placeholder="Password">
        <br><br>
        <input type="submit" value="Submit">
    </form>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Login',
  data () {
    return {
      msg: 'LCR Tool'
    }
  },
  methods: {
    login: function () {
      var params = new URLSearchParams()
      params.append('username', 'sirryanscott')
      params.append('password', 'maVoitan1')
      axios.post('/login/', params).then(response => {
        console.log('Logging in')
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

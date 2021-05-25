/* eslint-disable vue/valid-v-for */
<template>

    <div>
      <b-navbar toggleable="lg" type="dark" variant="success">
        <b-container>
<b-row>
  <b-col >

            <b-navbar-brand   v-for="s in sites" :href=s.url>{{s.name}}</b-navbar-brand>
  </b-col>
</b-row>
            <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
            <b-collapse id="nav-collapse" is-nav>
              <!-- Right aligned nav items -->
              <b-navbar-nav class="ml-auto">
                <b-nav-item-dropdown right>
                  <!-- Using 'button-content' slot -->
                  <template slot="button-content"><em>用户</em></template>
                  <b-dropdown-item href="/login">登入</b-dropdown-item>
                   <b-dropdown-item href="/register">注册</b-dropdown-item>
                  <b-dropdown-item href="#">退出</b-dropdown-item>
                </b-nav-item-dropdown>
              </b-navbar-nav>
            </b-collapse>
          </b-container>
      </b-navbar>
    </div>
</template>
<script>
export default {
  data () {
    return {
      sites: [
        {name: '首页', url: '/'},
        {name: '列表', url: '/job/'}
      ]
    }
  }

}
</script>

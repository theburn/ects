<template>
  <el-container>
    <el-header>
      <el-row :gutter="20">
        <el-col :span="4">
          <div class="grid-content logo">
            <h3 style="line-height: 36px;"><router-link to="/">ECTS</router-link></h3>
          </div>
        </el-col>
        <el-col :span="16">
          <el-menu :default-active="$route.path" class="el-menu-nav" mode="horizontal" router>
            <el-menu-item index="/">概览</el-menu-item>
            <el-menu-item index="/task">任务</el-menu-item>
            <el-menu-item index="/pipeline">流水线</el-menu-item>
            <el-menu-item index="/node">节点</el-menu-item>
            <el-menu-item index="/user">用户</el-menu-item>
            <el-menu-item index="/log">日志</el-menu-item>
            <el-menu-item index="/setting">设置</el-menu-item>
          </el-menu>
        </el-col>
        <el-col :span="4">
          <el-dropdown trigger="click" @command="handleCommand">
          <span class="el-dropdown-link">
            <div class="avatar grid-content" v-html="profile.name.slice(0,1)"></div>
          </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item command="a">个人信息</el-dropdown-item>
              <el-dropdown-item command="signOut">退出登陆</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </el-col>
      </el-row>
    </el-header>
    <el-main>
      <el-row :gutter="20">
        <el-col :span="18" :offset="3">
          <router-view></router-view>
        </el-col>
      </el-row>
    </el-main>
  </el-container>
</template>

<script>
  import {mapState} from 'vuex'

  export default {
    name: 'backend',
    data() {
      return {}
    },
    methods: {
      handleCommand(command) {
        switch (command) {
          case 'signOut':
            this.signOut();
            break;
          case 'profile':
            break;
        }
      },
      signOut(){
        this.$store.commit('SET_ACCESS_TOKEN', false);
        this.$store.commit('SET_PROFILE', false);
        this.$message.success('注销成功');
        this.$router.push("/signin");
      }
    },
    computed: {
      ...mapState({
        profile: state => state.account.profile,
        access_token: state => state.account.access_token
      })
    }
  }
</script>

<style lang="scss">
  #app {
    height: 100%;
    background-color: #f0f2f5;

    .el-header {
      width: 100%;
      background-color: white;
      box-shadow: 0 1px 4px rgba(0, 21, 41, .08);
      .el-dropdown {
        float: right;
      }
      .avatar {
        height: 36px;
        width: 36px;
        background-color: #8c939d;
        float: right;
      }
    }

    .el-menu-nav {
      display: table;
      margin: auto;
      align-items: center;
      text-align: center;

      &.el-menu--horizontal {
        border: none;
      }
    }

    .el-container {
      height: 100%;
      min-height: 100%;

      .grid-content {
        margin: 12px 0;
        border-radius: 4px;
        min-height: 36px;
        line-height: 36px;
        font-size: 26px;
        text-align: center;
        cursor: pointer;
        color: #e9e9e9;
      }
      .logo {
        a {
          color: #8c939d;
        }
        text-align: left;
      }
    }

    .el-main {
      height: 100%;
      display: block;

      .el-row {
        height: 100%;
      }

      .el-col {
        height: 100%;
      }
    }
  }

  .main-content {
    height: 100%;
    border-radius: 4px;
    padding: 0 20px;
  }
</style>

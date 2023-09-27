import { createRouter, createWebHashHistory } from "vue-router";
import { store } from '@/store/index';

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "",
      redirect: "/home",
      children: [
        {
          path: "home",
          name: "home",
          component: () => import("@/views/home.vue"),
        }
      ]
    },
    {
      path: "/login",
      name: "login",
      component: () => import("@/views/login.vue"),
    },
    {
      path: "/denied",
      name: "denied",
      component: () => import("@/views/denied.vue"),
    }
  ],
});

/**
 * 路由守卫，重定向
 */
router.beforeEach(async (to, from, next) => {
  // 如果是访问无权限页面直接放行
  if(to.name === 'denied'){
    next();
    return;
  }
  // 如果没有登录则跳转到登录界面
  if(!store.state.username && to.name !== 'login') {
    next({
      name: 'login',
    })
    return;
  }
  // 如果是login&没有登录则自动放行
  if(to.name === 'login' && !store.state.username){
    next();
    return;
  }
  next();
});

export default router;

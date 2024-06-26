import { createRouter, createWebHashHistory } from "vue-router";
import { store } from '@/store/index';

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "",
      redirect: "/home"
    },
    {
      path: "/home",
      name: "home",
      redirect: "/home/uapply",
      component: () => import("@/views/home.vue"),
      children: [
        {
          path: "/home/recommand",
          name: "recommand",
          component: () => import("@/views/components/admin/admin.vue"),
        },
        {
          path: "/home/uapply",
          name: "uapply",
          component: () => import("@/views/components/uapply/index.vue"),
        },
        {
          path: "/home/view-cv",
          name: "view-cv",
          component: () => import("@/views/components/cv/index.vue"),
        },
        {
          path: "/home/view-pc",
          name: "view-pc",
          component: () => import("@/views/components/pc/pc.vue"),
        }
      ]
    },
    {
      path: "/login",
      name: "login",
      component: () => import("@/views/components/login/login.vue"),
    },
    {
      path: "/denied",
      name: "denied",
      component: () => import("@/views/components/denied/denied"),
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
  if(to.name !== 'login' && !store.state.account) {
    next({
      name: 'login',
    })
    return;
  }
  // 如果是login&没有登录则自动放行
  if(to.name === 'login' && !store.state.account){
    next();
    return;
  }
  // 如果已经登陆过则重定向到home
  if(to.name === 'login' && store.state.account) {
    next({
      name: 'home',
    })
  }
  next();
});

export default router;

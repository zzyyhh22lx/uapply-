import { createStore } from 'vuex';
import { ADMIN, SUPERADMIN, COMMONUSER } from '../libs/constant';

export const store = createStore({
    state () {
        return {
            account: <string> "",
            roleType: <undefined | number> undefined,
            isShow: false, // 侧边栏是否显示
        }
      },
    getters: {
        isDenied: state => {
            return [ADMIN, SUPERADMIN, COMMONUSER].includes(state.roleType as number);
        },
    },
    mutations: {
        setUsername(state, account: string) {
            state.account = account;
        },
        setRoleType(state, roleType: number) {
            state.roleType = roleType;
        },
        setShowType(state, isShow: boolean) {
            state.isShow = isShow;
        }
    },
    actions: {

    }
});

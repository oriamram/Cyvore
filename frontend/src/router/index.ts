import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import { AuthService } from "../services/auth";

const routes: Array<RouteRecordRaw> = [
	{
		path: "/login",
		name: "Login",
		component: () => import("../views/LoginView.vue"),
		meta: { requiresAuth: false },
	},
	{
		path: "/register",
		name: "Register",
		component: () => import("../views/RegisterView.vue"),
		meta: { requiresAuth: false },
	},
	{
		path: "/dashboard",
		name: "Dashboard",
		component: () => import("../views/HomeView.vue"),
		meta: { requiresAuth: true },
	},
	{
		path: "/",
		redirect: "/dashboard",
	},
];

const router = createRouter({
	history: createWebHistory("/"),
	routes,
});

// Navigation guard
router.beforeEach((to, from, next) => {
	const authService = AuthService.getInstance();
	const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);
	const isAuthenticated = authService.isAuthenticated();

	if (requiresAuth && !isAuthenticated) {
		next("/login");
	} else if (!requiresAuth && isAuthenticated) {
		next("/dashboard");
	} else {
		next();
	}
});

export default router;

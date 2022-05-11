import { createRouter, createWebHistory } from "vue-router";
import HomeView from "@/views/HomeView";
import SearchView from "@/views/SearchView";
import LoginView from "@/views/LoginView";
import RegisterView from "@/views/RegisterView";
import UsersView from "@/views/UsersView";

const routes = [
  {
    path: "/",
    name: "home",
    component: HomeView,
  },
  {
    path: "/search",
    name: "search",
    component: SearchView,
  },
  {
    path: "/login",
    name: "login",
    component: LoginView,
  },
  {
    path: "/register",
    name: "register",
    component: RegisterView,
  },
  {
    path: "/users",
    name: "users",
    component: UsersView,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;

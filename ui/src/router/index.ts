import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue'),
      meta: { public: true },
    },
    {
      path: '/',
      component: () => import('../components/AppLayout.vue'),
      children: [
        { path: '', redirect: '/dashboard' },
        { path: 'dashboard', name: 'dashboard', component: () => import('../views/DashboardView.vue') },
        { path: 'profile', name: 'profile', component: () => import('../views/ProfileView.vue') },
        { path: 'vacancies', name: 'vacancies', component: () => import('../views/VacanciesView.vue') },
        { path: 'vacancies/:id', name: 'vacancy-detail', component: () => import('../views/VacancyDetailView.vue') },
        { path: 'cover-letters', name: 'cover-letters', component: () => import('../views/CoverLettersView.vue') },
        { path: 'cover-letters/:id', name: 'cover-letter-editor', component: () => import('../views/CoverLetterEditorView.vue') },
      ],
    },
  ],
})

router.beforeEach((to) => {
  const auth = useAuthStore()
  if (!to.meta.public && !auth.isAuthenticated) {
    return { name: 'login' }
  }
})

export default router

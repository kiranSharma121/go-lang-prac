import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080';

// Create axios instance
const api = axios.create({
    baseURL: API_BASE_URL,
    headers: {
        'Content-Type': 'application/json',
    },
});

// Request interceptor to add JWT token
api.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('jwt');
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// Response interceptor to handle auth errors
api.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response?.status === 401) {
            localStorage.removeItem('jwt');
            localStorage.removeItem('user');
            window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);

// Auth endpoints
export const authAPI = {
    signup: (userData) => api.post('/signup', userData),
    login: (credentials) => api.post('/login', credentials),
};

// Student endpoints
export const studentAPI = {
    getDashboard: () => api.get('/student/dashboard'),
    getCourses: () => api.get('/student/courses'),
    enrollInCourse: (courseId) => api.post(`/student/enroll/${courseId}`),
    getMyCourses: () => api.get('/student/my-courses'),
};

// Tutor endpoints
export const tutorAPI = {
    getDashboard: () => api.get('/tutor/dashboard'),
    createCourse: (courseData) => api.post('/tutor/course', courseData),
    updateCourse: (courseId, courseData) => api.put(`/tutor/course/${courseId}`, courseData),
    deleteCourse: (courseId) => api.delete(`/tutor/course/${courseId}`),
    getCourseStudents: (courseId) => api.get(`/tutor/${courseId}/students`),
};

export default api; 
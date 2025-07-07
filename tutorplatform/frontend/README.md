# Tutor Platform Frontend

A modern React + Tailwind CSS frontend for a tutoring platform with role-based dashboards for students and tutors.

## Features

### 🔐 Authentication
- JWT-based authentication
- Role-based access control (Student/Tutor)
- Protected routes with automatic redirection
- Persistent login state

### 👨‍🎓 Student Features
- Browse available courses
- Enroll in courses
- View enrolled courses
- Search and filter courses
- Course details modal

### 👨‍🏫 Tutor Features
- Create new courses
- Edit existing courses
- Delete courses
- View enrolled students
- Course management dashboard

### 🎨 UI/UX Features
- Modern, responsive design with Tailwind CSS
- Loading states and error handling
- Form validation
- Success/error notifications
- Mobile-friendly navigation

## Tech Stack

- **React 18** - UI framework
- **Vite** - Build tool and dev server
- **Tailwind CSS** - Styling
- **React Router** - Client-side routing
- **Axios** - HTTP client
- **Context API** - State management

## Getting Started

### Prerequisites

- Node.js 16+ 
- npm or yarn

### Installation

1. Install dependencies:
```bash
npm install
```

2. Start the development server:
```bash
npm run dev
```

3. Open your browser and navigate to `http://localhost:5173`

### Backend API

The frontend expects a Go backend running on `http://localhost:8080` with the following endpoints:

#### Authentication
- `POST /signup` - User registration
- `POST /login` - User login

#### Student Routes
- `GET /student/dashboard` - Student dashboard data
- `GET /student/courses` - Available courses
- `POST /student/enroll/:courseID` - Enroll in course
- `GET /student/my-courses` - Enrolled courses

#### Tutor Routes
- `GET /tutor/dashboard` - Tutor dashboard data
- `POST /tutor/course` - Create course
- `PUT /tutor/course/:id` - Update course
- `DELETE /tutor/course/:id` - Delete course
- `GET /tutor/:id/students` - Course students

## Project Structure

```
src/
├── components/          # Reusable components
│   ├── CourseCard.jsx
│   ├── Navbar.jsx
│   └── ProtectedRoute.jsx
├── contexts/           # React contexts
│   └── AuthContext.jsx
├── pages/             # Page components
│   ├── LoginPage.jsx
│   ├── SignupPage.jsx
│   ├── StudentDashboard.jsx
│   └── TutorDashboard.jsx
├── services/          # API services
│   └── api.js
├── App.jsx           # Main app component
└── main.jsx         # App entry point
```

## Key Components

### AuthContext
Manages authentication state, JWT tokens, and user information.

### ProtectedRoute
Route wrapper that checks authentication and role-based access.

### CourseCard
Reusable component for displaying course information with enrollment functionality.

### Navbar
Responsive navigation with role-based menu items and user profile.

## API Configuration

The API base URL is configured in `src/services/api.js`. Update the `API_BASE_URL` constant if your backend runs on a different port.

## Build for Production

```bash
npm run build
```

The built files will be in the `dist/` directory.

## Development

### Adding New Features

1. Create new components in `src/components/`
2. Add new pages in `src/pages/`
3. Update API services in `src/services/api.js`
4. Add routes in `src/App.jsx`

### Styling

The project uses Tailwind CSS for styling. Custom styles can be added in `src/index.css`.

## Contributing

1. Follow the existing code structure
2. Use Tailwind CSS for styling
3. Implement proper error handling
4. Add loading states for async operations
5. Test on different screen sizes

## License

MIT License

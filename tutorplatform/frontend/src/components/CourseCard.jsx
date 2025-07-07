import React, { useState } from 'react';
import { studentAPI } from '../services/api';

const CourseCard = ({ course, isEnrolled = false, onEnroll, onViewDetails }) => {
  const [enrolling, setEnrolling] = useState(false);

  const handleEnroll = async () => {
    if (isEnrolled) return;
    
    setEnrolling(true);
    try {
      await studentAPI.enrollInCourse(course.id);
      onEnroll?.(course);
    } catch (error) {
      console.error('Enrollment failed:', error);
    } finally {
      setEnrolling(false);
    }
  };

  return (
    <div className="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow duration-300 overflow-hidden">
      {/* Course Image */}
      <div className="h-48 bg-gradient-to-br from-blue-400 to-purple-500 relative">
        <div className="absolute inset-0 flex items-center justify-center">
          <div className="text-white text-6xl font-bold opacity-20">
            {course.title?.charAt(0) || 'C'}
          </div>
        </div>
      </div>

      {/* Course Content */}
      <div className="p-6">
        <div className="flex items-start justify-between mb-3">
          <h3 className="text-lg font-semibold text-gray-900 line-clamp-2">
            {course.title}
          </h3>
          {isEnrolled && (
            <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
              Enrolled
            </span>
          )}
        </div>

        <p className="text-gray-600 text-sm mb-4 line-clamp-3">
          {course.description || course.content || 'No description available'}
        </p>

        {/* Course Stats */}
        <div className="flex items-center justify-between mb-4">
          <div className="flex items-center space-x-1">
            <div className="flex items-center">
              {[...Array(5)].map((_, i) => (
                <svg
                  key={i}
                  className={`w-4 h-4 ${
                    i < 4 ? 'text-yellow-400' : 'text-gray-300'
                  }`}
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                </svg>
              ))}
            </div>
            <span className="text-sm text-gray-500 ml-1">(120)</span>
          </div>
          <span className="text-sm text-gray-500">Free</span>
        </div>

        {/* Action Buttons */}
        <div className="flex space-x-2">
          <button
            onClick={onViewDetails}
            className="flex-1 px-4 py-2 text-sm font-medium text-blue-600 bg-blue-50 rounded-md hover:bg-blue-100 transition-colors"
          >
            View Details
          </button>
          <button
            onClick={handleEnroll}
            disabled={isEnrolled || enrolling}
            className={`flex-1 px-4 py-2 text-sm font-medium rounded-md transition-colors ${
              isEnrolled
                ? 'bg-gray-100 text-gray-500 cursor-not-allowed'
                : enrolling
                ? 'bg-blue-100 text-blue-600 cursor-wait'
                : 'bg-blue-600 text-white hover:bg-blue-700'
            }`}
          >
            {enrolling ? (
              <div className="flex items-center justify-center">
                <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-white"></div>
                <span className="ml-2">Enrolling...</span>
              </div>
            ) : isEnrolled ? (
              'Enrolled'
            ) : (
              'Enroll'
            )}
          </button>
        </div>
      </div>
    </div>
  );
};

export default CourseCard; 
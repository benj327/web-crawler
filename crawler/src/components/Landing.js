import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

const Landing = () => {
  const [urls, setUrls] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    const response = await axios.post('http://localhost:8080/startScan', {
      urls,
    });

    navigate('/results', { state: { results: response.data.sensitiveInfo } });
  };

  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center">
      <div className="bg-white p-8 rounded-lg shadow-md w-full md:w-1/2 lg:w-1/3">
        <h1 className="text-2xl font-bold mb-6">Security Web Crawler</h1>
        <form onSubmit={handleSubmit} className="space-y-4">
          <label htmlFor="urls" className="block">
            Enter a URL or list of URLs (comma-separated):
          </label>
          <input
            type="text"
            id="urls"
            value={urls}
            onChange={(e) => setUrls(e.target.value)}
            className="w-full p-2 border border-gray-300 rounded-md"
          />
          <button
            type="submit"
            className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600"
          >
            Scan
          </button>
        </form>
      </div>
    </div>
  );
};

export default Landing;

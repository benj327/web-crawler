import React from 'react';
import { useLocation } from 'react-router-dom';

const Results = () => {
  const location = useLocation();
  const results = location.state.results;

  const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text).then(
      () => {
        alert('Copied to clipboard!');
      },
      (err) => {
        console.error('Could not copy text: ', err);
      },
    );
  };

  return (
    <div className="min-h-screen bg-gray-100 p-6">
      <h1 className="text-2xl font-bold mb-6">Scan Results</h1>
      {Object.keys(results).map((url) => (
        <div key={url} className="bg-white p-4 rounded-lg shadow-md mb-4">
          <h2 className="text-xl font-semibold mb-2">{url}</h2>
          <ul className="list-disc pl-6 space-y-1">
            {results[url].map((email, index) => (
              <li key={index}>{email}</li>
            ))}
          </ul>
          <button
            onClick={() => copyToClipboard(results[url].join(', '))}
            className="bg-blue-500 text-white px-4 py-2 rounded-md mt-2 hover:bg-blue-600"
          >
            Copy to clipboard
          </button>
        </div>
      ))}
    </div>
  );
};

export default Results;

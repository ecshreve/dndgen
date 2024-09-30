import React from 'react';
import './ValueDisplay.css'; // Import CSS styles

const ValueDisplay = ({ value, label }: { value: number; label?: string }) => {
  // Determine the background color based on whether the value is positive or negative, gray if 0
  const backgroundColor = value > 0 ? 'green' : value === 0 ? 'gray' : 'red';

  return (
    <div className="value-display">
      <div
        className="circle"
        style={{ backgroundColor }}
      >
        {value > 0 && '+'}
        {value}
      </div>
      {label && <div className="label">
        {label}
      </div>}
    </div>
  );
};

export default ValueDisplay;

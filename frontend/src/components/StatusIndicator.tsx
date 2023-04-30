import React from 'react';
import '../css/StatusIndicator.css';

function StatusIndicator(props) {
  const { state } = props;
  console.log(state);

  if (state === 'loading') {
    return (
      <div className="status-indicator">
        <div className="status-indicator__loading"></div>
      </div>
    );
  }

  if (state === 'success') {
    return (
      <div className="status-indicator">
        <div className="status-indicator__success show">&#10003;</div>
      </div>
    );
  }

  if (state === 'fail') {
    return (
      <div className="status-indicator">
        <div className="status-indicator__fail show">&#10007;</div>
      </div>
    );
  }

  return null;
}

export default StatusIndicator;
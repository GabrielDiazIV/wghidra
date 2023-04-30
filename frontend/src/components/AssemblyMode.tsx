import React from 'react';
import '../css/AssemblyMode.css';

function AssemblyMode(props) {
  const { body } = props;

  return (
    <div className="assembly-mode">
      <div className="assembly-mode__textbox-container">
        <div className="assembly-mode__textbox">{body}</div>
      </div>
    </div>
  );
}

export default AssemblyMode;

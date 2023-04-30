import React from 'react';

function Sidebar(props) {
  const { functionNames, onClick, selectedOption } = props;

  return (
    <div className="sidebar">
      {functionNames.map((name) => (
        <div className={`sidebar__option ${selectedOption === name ? 'active' : ''}`} key={name} onClick={() => onClick(name)}>
          {name}
        </div>
      ))}
    </div>
  );
}

export default Sidebar;

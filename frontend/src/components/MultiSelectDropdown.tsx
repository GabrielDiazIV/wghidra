import { useState } from 'react';
import '../css/MultiSelectDropdown.css';

function MultiSelectDropdown(props) {
    const {options, setScripts} = props;

  const [isOpen, setIsOpen] = useState(false);
  const [selectedOptions, setSelectedOptions] = useState([]);

  function toggleDropdown() {
    setIsOpen(!isOpen);
  }

  function handleOptionClick(option) {
    if (selectedOptions.includes(option)) {
      setSelectedOptions(selectedOptions.filter((o) => o !== option));
    } else {
      setSelectedOptions([...selectedOptions, option]);
    }
    setScripts(selectedOptions);
  }

  return (
    <div className="dropdown-container">
    <button className="dropdown-button" onClick={toggleDropdown}>
      Select Scripts
    </button>
    {isOpen && (
      <div className="dropdown-menu">
        {options.map((option) => (
          <div
            key={option}
            className={`dropdown-option ${
              selectedOptions.includes(option) ? "selected" : ""
            }`}
            onClick={() => handleOptionClick(option)}
          >
            {option}
          </div>
        ))}
      </div>
    )}
  </div>
);
}

export default MultiSelectDropdown;
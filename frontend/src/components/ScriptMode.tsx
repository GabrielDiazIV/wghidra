import React, { useState, useRef, useEffect } from 'react';
import SyntaxHighlighter from 'react-syntax-highlighter';
import { atelierCaveDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import '../css/ScriptMode.css';


function ScriptMode(props) {
  const [selectedOption, setSelectedOption] = useState('');
  const [output, setOutput] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const { body } = props;
  const textAreaRef = useRef(null);

  const scriptOptions = ["Analyzer", "Decompiler", "Debugger", "Disassembler", "Extractor", "Visualizer", "Tracer", "Profiler", "Inspector", "Archiver"]

  const handleOptionClick = (option) => {
    setSelectedOption(option);
    setOutput('');
  };

  const handleSubmitClick = async () => {
    setIsLoading(true);
    try {
      const response = await fetch(`https://example.com/api/${selectedOption}`);
      const data = await response.text();
      setOutput(data);
      textAreaRef.current.scrollIntoView({ behavior: 'smooth' });
    } catch (error) {
      console.error(error);
      
      setOutput("Lorem ipsum dolor sit amet, consectetur adipiscing elit. In malesuada turpis enim, in placerat odio vestibulum vel. Praesent fermentum sem ut neque tempus mattis. Donec malesuada ligula at tortor tincidunt hendrerit. Morbi lobortis vitae urna ac pharetra. Vestibulum lacinia, tellus eget lacinia mollis, neque sem rutrum ipsum, id ultrices dolor orci sed lectus. Pellentesque molestie sem vitae felis ultrices blandit. Aliquam lobortis condimentum arcu eget suscipit. Praesent vestibulum egestas lorem nec egestas. Proin eget ipsum ex. Nulla ullamcorper gravida ligula. Sed sit amet arcu vehicula, fringilla nisi nec, egestas velit. Integer tristique tempus nisi, id ultrices nunc convallis quis. Phasellus at fringilla ligula.");
      textAreaRef.current.scrollIntoView({ behavior: 'smooth' });
    } finally {
      setIsLoading(false);
    }
  };

  // run when body changes
  useEffect(() => {
    setSelectedOption('');
    setOutput('');
    setIsLoading(false);
  }, [body]);

  return (
    <div className="script-mode">
      <div className="script-mode__textbox-container">
        <SyntaxHighlighter className="script-mode__textbox" language="cpp" style={atelierCaveDark}>
          {body}
        </SyntaxHighlighter>
      </div>
      <div className="script-mode__dropdown-container">
        <select className="script-mode__dropdown" value={selectedOption} onChange={(e) => handleOptionClick(e.target.value)}>
          <option value="" disabled>Select an option</option>
          {scriptOptions.map((option) => (
            <option key={option} value={option}>{option}</option>
          ))}
        </select>
        <button className={`script-mode__button${selectedOption && !isLoading ? ' active' : ''}`} onClick={handleSubmitClick}>
            {isLoading ? 'Loading...' : 'GO'}
        </button>
      </div>
      {output && (
        <div className="script-mode__output-container no-scrollbar">
            <div className="script-mode__output" ref={textAreaRef} >
            {output}
            </div>
        </div>
      )}
    </div>
  );
}

export default ScriptMode;

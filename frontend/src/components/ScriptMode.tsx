import React, { useState, useRef, useEffect } from 'react';
import SyntaxHighlighter from 'react-syntax-highlighter';
import { atelierCaveDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import MultiSelectDropdown from './MultiSelectDropdown.tsx';
import '../css/ScriptMode.css';


function ScriptMode(props) {
  const [selectedOptions, setSelectedOptions] = useState('');
  const [output, setOutput] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const { body, projectId } = props;

  const scriptOptions = ["Analyzer", "Decompiler", "Debugger", "Disassembler", "Extractor", "Visualizer", "Tracer", "Profiler", "Inspector", "Archiver"]

  const setScripts = (options) => {
    setSelectedOptions(options)
  }

  const handleSubmitClick = async () => {
    setIsLoading(true);
    try {
      const requestBody = {
        projectId: projectId,
        scripts: selectedOptions
      };

      const response = await fetch(`http://localhost:6969/api/scripts`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(requestBody)
      });

      const data = await response.json();
      setOutput(data.results);
    } catch (error) {
      console.error(error);

      setOutput("Lorem ipsum dolor sit amet, consectetur adipiscing elit. In malesuada turpis enim, in placerat odio vestibulum vel. Praesent fermentum sem ut neque tempus mattis. Donec malesuada ligula at tortor tincidunt hendrerit. Morbi lobortis vitae urna ac pharetra. Vestibulum lacinia, tellus eget lacinia mollis, neque sem rutrum ipsum, id ultrices dolor orci sed lectus. Pellentesque molestie sem vitae felis ultrices blandit. Aliquam lobortis condimentum arcu eget suscipit. Praesent vestibulum egestas lorem nec egestas. Proin eget ipsum ex. Nulla ullamcorper gravida ligula. Sed sit amet arcu vehicula, fringilla nisi nec, egestas velit. Integer tristique tempus nisi, id ultrices nunc convallis quis. Phasellus at fringilla ligula.");
    } finally {
      setIsLoading(false);
    }
  };
  return (
    <div className="script-mode">
      <div className="script-mode__textbox-container">
        <SyntaxHighlighter className="script-mode__textbox" language="cpp" style={atelierCaveDark}>
          {body}
        </SyntaxHighlighter>
      </div>
      <div className="script-mode__dropdown-container">
        <div>
          <MultiSelectDropdown options={scriptOptions} setScripts={setScripts} />
        </div>
        <button className={`script-mode__button${selectedOptions && !isLoading ? ' active' : ''}`} onClick={handleSubmitClick}>
          {isLoading ? 'Loading...' : 'GO'}
        </button>
      </div>
      {output && output.map((taskOutput) => {
        return (

          <div className="script-mode__output-container no-scrollbar">
            <div className="script-mode__output">
              {taskOutput.error ? taskOutput.error : taskOutput.body}
            </div>
          </div>
        )
      })}
    </div>
  );
}

export default ScriptMode;

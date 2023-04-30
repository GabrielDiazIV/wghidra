import React, { useState, useRef, useEffect } from 'react';
import AceEditor from 'react-ace';
import 'ace-builds/src-noconflict/mode-c_cpp';
import 'ace-builds/src-noconflict/theme-monokai';
import '../css/RunnerMode.css';


function RunnerMode(props) {
    const {params } = props;
  const [body, setBody] = useState(props.body);
  const [output, setOutput] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [textValues, setTextValues] = useState<string[]>(new Array(params.length).fill(''));

  const textAreaRef = useRef(null);

  const handleSubmitClick = async () => {
    setIsLoading(true);
    try {
      const response = await fetch(`https://example.com/api/${textValues}`);
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

  const handleTextChange = (index: number, value: string) => {
    const newValues = [...textValues];
    newValues[index] = value;
    setTextValues(newValues);
  };

  // run when body changes
  useEffect(() => {
    setTextValues(new Array(params.length).fill(''));
    setOutput('');
    setIsLoading(false);
    setBody(props.body);
  }, [props.body]);

  return (
    <div className="runner-mode">
      <div className="runner-mode__textbox-container">
      <AceEditor
        className="runner-mode__textbox"
        mode="c_cpp"
        theme="monokai"
        value={body}
        onChange={(newValue) => setBody(newValue)}
        fontSize={16}
        width="100%"
        height="auto"
        minLines={15}
        maxLines={Infinity}
        wrapEnabled={true}
        editorProps={{ $blockScrolling: true }}
      />
      </div>
      <div className="runner-mode__dropdown-container">
        <div className='runner-mode__params-container'>
            {params.map((param, index) => (
                <div  className='runner-mode__param-input-container' key={index}>
                    <span className='runner-mode__param-input-type'>{param}: </span>
                    <input className = 'runner-mode__param-input-box' type="text" value={textValues[index]}  onChange={(e) => handleTextChange(index, e.target.value)} />
                </div>
            ))}
        </div>
        <button className={`runner-mode__button${!isLoading ? ' active' : ''}`} onClick={handleSubmitClick}>
            {isLoading ? 'Loading...' : 'RUN'}
        </button>
      </div>
      {output && (
        <div className="runner-mode__output-container no-scrollbar">
          <div className="runner-mode__output" ref={textAreaRef} >
            {output}
            </div>
        </div>
      )}
    </div>
  );
}

export default RunnerMode;

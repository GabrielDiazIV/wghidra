import React, { useState, useRef, useEffect } from 'react';
import AceEditor from 'react-ace';
import 'ace-builds/src-noconflict/mode-c_cpp';
import 'ace-builds/src-noconflict/theme-monokai';
import '../css/RunnerMode.css';
import { ApiResponse, ScriptResponse, options } from '../App';


function RunnerMode(props) {
  const {
    params, name,
    editedFunctionList,
    setEditedFunctionList
  } = props;

  const [body, setBody] = useState(props.body);
  const [output, setOutput] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [textValues, setTextValues] = useState<string[]>(new Array(params.length).fill(''));

  const textAreaRef = useRef(null);

  const handleSubmitClick = async () => {
    setIsLoading(true);
    try {
      // use

      for (let i = 0; i < editedFunctionList.length; i++) {
        if (editedFunctionList[i].name === name) {
          editedFunctionList[i].body = body;
          break
        }
      }

      const requestBody = {
        functions: editedFunctionList,
        execute_function: name,
        parameters: textValues
      };

      let req_opts = options()
      req_opts.body = JSON.stringify(requestBody)
      console.log(req_opts.body)

      const response = await fetch(`http://localhost:6969/api/run`, req_opts);

      const resp = await response.json() as ApiResponse<ScriptResponse>;
      setEditedFunctionList([...editedFunctionList])
      setOutput(resp.data.results[0].output.output);
      // textAreaRef.current.scrollIntoView({ behavior: 'smooth' });
    } catch (error) {
      console.error(error);

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

  function bodyChange(newValue) {
    setBody(newValue);
  }

  // run when body changes
  useEffect(() => {
    setTextValues(new Array(params.length).fill(''));
    setOutput('');
    setIsLoading(false);
    setBody(props.body);
  }, [name]);

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
            <div className='runner-mode__param-input-container' key={index}>
              <span className='runner-mode__param-input-type'>{param}: </span>
              <input className='runner-mode__param-input-box' type="text" value={textValues[index]} onChange={(e) => handleTextChange(index, e.target.value)} />
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

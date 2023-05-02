import React, { useState } from 'react';
import Sidebar from './Sidebar';
import logo from '../assets/logo.png';
import '../css/MainPage.css';
import '../index.css'
import ScriptMode from '../components/ScriptMode.tsx'
import RunnerMode from '../components/RunnerMode.tsx'
import AssemblyMode from '../components/AssemblyMode.tsx'

function MainPage(props) {
  const { functionList, assembly, projectId } = props;


  const [selectedMode, setSelectedMode] = useState('script');
  const [selectedOption, setSelectedOption] = useState(functionList[0].name);
  const [selectedFunction, setSelectedFunction] = useState<any>(functionList[0]);
  const [editedFunctionList, setEditedFunctionList] = useState(functionList);
  const [runnableFunction, setRunnableFunction] = useState<any>(functionList[0]);

  console.log(functionList);

  const functionNameList = functionList.map(f => f.name)

  const handleOptionClick = (option) => {
    setSelectedOption(option);

    functionList.forEach(f => {
      if (f.name === option) {
        setSelectedFunction(f);
        return;
      }

    });

    editedFunctionList.forEach(f => {
      if (f.name === option) {
        setRunnableFunction(f);
        return;
      }

    });
  };

  const handleRunnableBodyChange = (fname, newBody) => {
    for (let i = 0; i < editedFunctionList.length; i++) {
      if (editedFunctionList[i].name === fname) {
        editedFunctionList[i].body = newBody;
        return;
      }
    }
  }

  let mainComponent;
  if (selectedMode === 'script') {
    mainComponent = <ScriptMode projectId={projectId} body={selectedFunction.body} />
  }
  else if (selectedMode === 'runner') {
    mainComponent = <RunnerMode projectId={projectId} name={runnableFunction.name} onBodyChange={handleRunnableBodyChange} body={runnableFunction.body} params={runnableFunction.parameters} />
  }
  else if (selectedMode === 'assembly') {
    mainComponent = <AssemblyMode projectId={projectId} body={assembly} />
  }

  return (
    <div className="main-page">
      <div className="topbar">
        <div className={`topbar__option ${selectedMode === 'script' ? 'active' : ''}`} onClick={() => setSelectedMode('script')}>
          Script Mode
        </div>
        <div className={`topbar__option ${selectedMode === 'runner' ? 'active' : ''}`} onClick={() => setSelectedMode('runner')}>
          Runner Mode
        </div>
        <div className={`topbar__option ${selectedMode === 'assembly' ? 'active' : ''}`} onClick={() => setSelectedMode('assembly')}>
          Assembly Mode
        </div>
      </div>
      <h1 className="main-page__title">The Decompilation Destination</h1>
      <img
        className="main-page__logo"
        src={logo}
        alt="App Logo"
        onClick={() => props.homeHandler()}
      />
      <Sidebar functionNames={functionNameList} selectedOption={selectedOption} onClick={handleOptionClick} />
      <div className="main-content">
        {mainComponent}
      </div>
    </div>
  );
}

export default MainPage;

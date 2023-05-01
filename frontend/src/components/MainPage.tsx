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

  const TODOcode = assembly ? assembly : `
  section .text
  global _start

  _start:
      ; Define variables
      mov     eax, 10
      mov     ebx, 20
      
      ; Calculate sum
      add     eax, ebx
      
      ; Print result
      mov     edi, 1
      mov     esi, result
      mov     edx, 10
      xor     eax, eax
      cld
      rep     movsb
      
      ; Exit program
      xor     eax, eax
      mov     ebx, 0
      int     0x80

  section .data
  result:     db  "Result: ", 0
`;
  const TODOlist = functionList ? functionList : [
    {
      name: "Object 1",
      parameters: ["foo", "bar", "baz"],
      body: `#include <iostream>

      int main() {
          std::cout << "Hello World!" << std::endl;
          return 0;
      }`},
    {
      name: "Object 2",
      parameters: ["apple", "banana", "orange"],
      body: `#include <iostream>

      int factorial(int n) {
          if (n == 0) {
              return 1;
          }
          return n * factorial(n - 1);
      }
      
      int main() {
          int num = 5;
          std::cout << "Factorial of " << num << " is " << factorial(num) << std::endl;
          return 0;
      }`},
    {
      name: "Object 3",
      parameters: ["red", "green", "blue"],
      body: `#include <iostream>
      #include <cmath>
      
      int main() {
          int n, p;
          while (std::cin >> n >> p) {
              std::cout << std::fixed << std::pow(p, 1.0 / n) << std::endl;
          }
          return 0;
      }`}
  ];

  const [selectedMode, setSelectedMode] = useState('script');
  const [selectedOption, setSelectedOption] = useState(TODOlist[0].name);
  const [selectedFunction, setSelectedFunction] = useState<any>(TODOlist[0]);
  const [editedFunctionList, setEditedFunctionList] = useState(TODOlist);
  const [runnableFunction, setRunnableFunction] = useState<any>(TODOlist[0]);

  console.log(TODOlist);

  const functionNameList = TODOlist.map(f => f.name)

  const handleOptionClick = (option) => {
    setSelectedOption(option);

    TODOlist.forEach(f => {
        if(f.name === option) {
            setSelectedFunction(f);
            return;
        }
        
    });

    editedFunctionList.forEach(f => {
      if(f.name === option) {
          setRunnableFunction(f);
          return;
      }
      
  });
  };

  const handleRunnableBodyChange =(fname, newBody) => {
    for (let i = 0; i < editedFunctionList.length; i++) {
      if (editedFunctionList[i].name === fname) {
        editedFunctionList[i].body = newBody;
        return;
      }
    }
  }

  let mainComponent;
  if (selectedMode === 'script') {
    mainComponent = <ScriptMode projectId={projectId} body={selectedFunction.body}/>
  }
  else if(selectedMode === 'runner') {
    mainComponent = <RunnerMode projectId={projectId} name={runnableFunction.name} onBodyChange={handleRunnableBodyChange} body={runnableFunction.body} params={['String', 'int', 'int']}/> 
  }
  else if (selectedMode === 'assembly') {
    mainComponent = <AssemblyMode projectId={projectId} body={TODOcode} />
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

import React from 'react';
import SyntaxHighlighter from 'react-syntax-highlighter';
import { atelierCaveDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import '../css/AssemblyMode.css';

function AssemblyMode(props) {
  const { body } = props;

  return (
    <div className="assembly-mode">
      <div className="assembly-mode__textbox-container">
        <SyntaxHighlighter className="assembly-mode__textbox" language="asm" style={atelierCaveDark}>
          {body}
        </SyntaxHighlighter>
      </div>
    </div>
  );
}

export default AssemblyMode;

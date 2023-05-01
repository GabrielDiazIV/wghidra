import React, { useState, useRef } from 'react';
import logo from '../assets/logo.png';
import uploadIcon from '../assets/upload-icon.png';
import '../css/LandingPage.css'
import StatusIndicator from './StatusIndicator';

function LandingPage(props) {
  const { status, error, transition } = props;
  console.log(error);
  const inputFile = useRef<HTMLInputElement | null>(null);

  function handleUploadClick() {
    inputFile.current.click();
  }

  function handleFileSelected(event) {
    const file = event.target.files[0];
    // Do something with the selected file
    props.onFileSelect(file);
  }

  let maincontent;
  console.log(status);
  if(status === 'start') {
    maincontent = (
      <div className="landing-page__upload">
        <button
          className="landing-page__upload-button"
          onClick={handleUploadClick}
        >
          <img
            className="landing-page__upload-icon"
            src={uploadIcon}
            alt="Upload Icon"
          />
          <p className="landing-page__upload-text">Upload a file</p>
        </button>
        <input type='file' id='file' ref={inputFile} style={{display: 'none'}}
            onChange={handleFileSelected}
        />
      </div>
    )
    }
    else if (status === 'fail') {
        maincontent = (
            <div>
                <StatusIndicator state={status}/>
                <p className='landing-page__error'>{error}</p>
            </div>
        )
    }
    else {
        maincontent = <StatusIndicator state={status}/>
    }

  return (
    <div className="landing-page">
      <h1 className="landing-page__title">The Decompilation Destination</h1>
      <img
        className="landing-page__logo"
        src={logo}
        alt="App Logo"
      />
      <div className={'landing-page__content ' + transition}>
        {maincontent}
      </div>
    </div>
  );
}

export default LandingPage;
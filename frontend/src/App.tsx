import { useState } from 'react'
import LandingPage from './components/LandingPage'
import MainPage from './components/MainPage';

function App() {
  const [screen, setScreen] = useState('landing');
  const [status, setStatus] = useState('start');
  const [transition, setTransition] = useState('');
  const [file, setFile] = useState('');
  const [functionList, setFunctionList] = useState([])
  const [errorMessage, setErrorMessage] = useState('')

  function transitionScreens(responseType) {
    console.log("Setting status to " + responseType);
    setStatus(responseType);
    
    setTimeout(() => {
      if(responseType === 'fail') {
        setStatus('start');
        setTransition('fade-in');
      }
      else
        setScreen('main');

    }, 3000);
    setTimeout(() => setTransition('fade-out'), 2600); // Change status to fade-out after 3 seconds
  }

  const homeHandler = () => {
    setStatus('start');
    setTransition('fade-in');
    setScreen('landing');
  }

  const handleFileSelect = async(file)  => {
    setFile(file);
    setStatus('loading');

    const formData = new FormData();
    formData.append('file', file);
  
    fetch('/api/upload', {
      method: 'POST',
      body: formData
    })
    .then(async response => {
      if(response.ok) {
        const data = await response.json();
        setFunctionList(data.functions);
        transitionScreens('success');
      }
      else {
        const data = await response.json();
        setErrorMessage(data.error);
        transitionScreens('fail');
      }
    })
    .catch(error => {
      setErrorMessage(error.message);
      transitionScreens('success');
    });
  }

  let content = screen === 'landing' ? <LandingPage status={status} error={errorMessage} transition={transition} onFileSelect={(file) => {handleFileSelect(file)}}/> : <MainPage homeHandler={homeHandler}/>;
  return (
    <div>
      {content}
    </div>
  )
}

export default App

import { useState } from 'react'
import LandingPage from './components/LandingPage'
import MainPage from './components/MainPage';


export const url = '/api/project';
export const options = () => {
  return {
    method: 'POST',
    body: "",
    headers: {
      Accept: '*/*',
      // 'User-Agent': 'Thunder Client (https://www.thunderclient.com)',
      // 'Content-Type': 'multipart/form-data'
    }
  }
};

export type ApiResponse<T> = {
  data: T,
  error: { msg: string, code: number }
}

export type Function = {
  name: string,
  body: string,
  parameters: string[],

}

export type Project = {
  assembly: string,
  functions: Function[],
  projectId: string,
}


function App() {
  const [screen, setScreen] = useState('landing');
  const [status, setStatus] = useState('start');
  const [transition, setTransition] = useState('');
  const [assembly, setAssembly] = useState('');
  const [projectId, setProjectID] = useState('');
  const [functionList, setFunctionList] = useState<Function[]>([])
  const [errorMessage, setErrorMessage] = useState('')

  function transitionScreens(responseType) {
    console.log("Setting status to " + responseType);
    setStatus(responseType);

    setTimeout(() => {
      if (responseType === 'fail') {
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

  const handleFileSelect = async (file) => {
    setStatus('loading');

    const formData = new FormData();
    formData.append('project', file);
    const controller = new AbortController();
    const timeout = setTimeout(() => {
      controller.abort()
    }, 60_000)

    let request_opt = options()
    request_opt.body = formData
    fetch(url, request_opt)
      .then(async response => {
        console.log(response)
        if (response.ok) {
          const data = await response.json() as ApiResponse<Project>;

          console.log(data)
          const fns = data.data.functions
          fns.sort()
          setFunctionList(fns);
          setAssembly(data.data.assembly);
          setProjectID(data.data.projectId);
          transitionScreens('success');
        }
        else {
          const data = await response.json();
          setErrorMessage(data.error.message);
          transitionScreens('fail');
        }
      })
      .catch(error => {
        console.log(error)
        setErrorMessage(error.message);
        transitionScreens('fail');
      })
      .finally(() => {
        clearTimeout(timeout)
      });
  }

  let content = screen === 'landing' ? <LandingPage status={status} error={errorMessage} transition={transition} onFileSelect={(file) => { handleFileSelect(file) }} /> : <MainPage projectId={projectId} homeHandler={homeHandler} functionList={functionList} assembly={assembly} />;
  return (
    <div>
      {content}
    </div>
  )
}

export default App

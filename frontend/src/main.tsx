import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import LandingPage from './components/LandingPage.tsx'
import StatusIndicator from './components/StatusIndicator.tsx'
import MainPage from './components/MainPage.tsx'
import './index.css'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)

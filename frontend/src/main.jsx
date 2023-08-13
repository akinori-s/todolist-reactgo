import React from 'react'
import ReactDOM from 'react-dom/client'
import { Route, Routes, Link, BrowserRouter, useLocation } from "react-router-dom";
import App from './App.jsx'
import Header from './components/Header.jsx'
import LoginPage from './components/LoginPage.jsx'
import SignUpPage from './components/SignUpPage.jsx'
import ErrorPage from './components/ErrorPage.jsx'
import { AuthProvider } from './contexts/AuthContext.jsx'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
	<BrowserRouter>
		<AuthProvider>
			<Header />
			<Routes>
				<Route path="/" element={<App />} />
				<Route path="/login" element={<LoginPage />} />
				<Route path="/signup" element={<SignUpPage />} />
				<Route path="*" element={<ErrorPage />} />
			</Routes>
		</AuthProvider>
	</BrowserRouter>
  </React.StrictMode>,
)

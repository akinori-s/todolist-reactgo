import React, { useState } from 'react';
import { Link, useNavigate } from "react-router-dom";
import useForm from '../hooks/useForm';
import { login } from '../api/authApi';

function LoginPage() {
	const [loginInfo, handleChange, resetForm] = useForm({ email: '', password: '' });
	const [error, setError] = useState(null);
	const navigate = useNavigate();
	
	const handleSubmit = async () => {
		setError(null);
		console.log("Logging in with", loginInfo); // wip: remove later

		try {
			const response = await login(loginInfo);
		} catch (err) {
			setError(err.message);
		}
		resetForm();
		navigate("/");
		console.log("Logged in successfully!", response.data);
	};

	return (
	<div className="flex flex-col items-center h-screen w-full py-20 space-y-8 min-w-[300px]">
		<h1 className="text-2xl text-center">Login</h1>
		<div className="flex flex-col w-1/2 min-w-[300px] max-w-[400px] space-y-4">
			<p style={{ color: 'red' }}>{error}</p>
			<input 
				type="email" 
				name="email"
				value={loginInfo.email}
				onChange={handleChange}
				className="border rounded-md px-2 py-1 w-full" 
				placeholder="Email Address"
			/>
			<input 
				type="password" 
				name="password"
				value={loginInfo.password}
				onChange={handleChange}
				className="border rounded-md px-2 py-1 w-full" 
				placeholder="Password"
			/>
			<button 
				className="border rounded-md px-4 py-2 bg-blue-500 text-white hover:bg-blue-600 w-full"
				onClick={handleSubmit}
			>
				Login
			</button>
		</div>
		<div className="flex flex-row items-center space-x-4">
			<button 
				className="border rounded-md px-4 py-2 bg-gray-200 text-white hover:bg-gray-400"
				onClick={() => {/* wip:Handle Forgot Password */}}
			>
				Forgot Password?
			</button>
			<Link 
				to="/signup"
				className="border rounded-md px-4 py-2 bg-blue-500 text-white hover:bg-blue-600"
			>
				Sign Up
			</Link>
		</div>
	</div>
	);
}

export default LoginPage;

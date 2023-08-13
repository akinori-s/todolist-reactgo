import React from 'react';
import { Route, Routes, Link, BrowserRouter, useLocation } from "react-router-dom";

function Header() {
	const location = useLocation()

	return (
		<div className="flex justify-between items-center h-14 w-full px-6 border-b min-w-[300px]">
			<Link 
				to="/"
				className="text-2xl text-center" 
			>
				TodoList
			</Link>
			<p className="text-sm" value={[location.pathname]}></p>
			{(location.pathname !== '/login' && location.pathname !== '/signup') ? 
				<div>
					<Link 
						to="/login" 
						className="border rounded-md px-2 py-1 mr-2 bg-blue-500 text-white hover:bg-blue-600"
					>
						Login
					</Link>
					<Link 
						to="/signup"
						className="border rounded-md px-2 py-1 text-blue-500 border-solid border-blue-500 hover:text-white hover:bg-blue-500"
					>
						Sign Up
					</Link>
				</div>
			:
				<div></div>
			}
		</div>
	);
}

export default Header;

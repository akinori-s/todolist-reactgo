import { useState } from 'react'
import './App.css'

function App() {
	const [todos, setTodos] = useState([
		{ text: "Tasdddddddddddddddddddddddddddddssssssssssss dddddddddddsk 1", completed: false },
		{ text: "Task 2", completed: false },
	]);

	const toggleCompletion = (index) => {
		const updatedTodos = [...todos];
		updatedTodos[index].completed = !updatedTodos[index].completed;
		setTodos(updatedTodos);
	};

	return (
		<div className="flex flex-col items-center h-screen w-full py-20 space-y-8 min-w-[300px]">
      <h1 className="text-2xl text-center">TodoList App</h1>
			<div className="flex flex-row items-center space-x-4">
				<input type="text" className="border rounded-md px-2 py-1 w-3/4" placeholder="Enter a task" />
				<button className="border rounded-md px-2 py-1 bg-blue-500 text-white hover:bg-blue-600">Add</button>
			</div>
      <ul className="list-decimal list-inside h-1/2 w-1/2 min-w-[300px] max-w-[500px] overflow-y-auto border rounded-md">
        {todos.map((todo, index) => (
          <li key={index} className="flex justify-between items-center py-2 px-4 border-b last:border-0 hover:bg-gray-50 w-full">
						<div className="flex items-center flex-grow">
							<button 
								className={`mr-4 p-[5px] ${todo.completed ? 'bg-green-500' : 'bg-red-500'} rounded`} 
								onClick={() => toggleCompletion(index)}
							/>
							<span className={`${todo.completed ? 'line-through' : ''} break-all whitespace-normal`}>{todo.text}</span>
  					</div>
						<button className="px-2 py-1 bg-gray-200 text-white rounded hover:bg-red-400">
							Delete
						</button>
					</li>
        ))}
      </ul>
    </div>
	);
}

export default App

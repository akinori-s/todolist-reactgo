import { useState, useEffect } from 'react'
import { addTodo, getTodoList, updateTodo, deleteTodo } from './api/todoApi'
import './App.css'


function App() {
	const [todos, setTodos] = useState([]);
	const [newTodo, setNewTodo] = useState('');

	useEffect(() => {
		async function getTodos() {
			try {
				const data = await getTodoList();
				setTodos(data);
			} catch (error) {
				console.error("Failed to fetch todos:", error);
			}
		}
		getTodos();
	}, []);

	const handleAddTodo = async () => {
		const todoText = newTodo.trim();
		const newId = todos.length + 1;
		if (newTodo.trim()) {
			const todo = {id: newId, text: todoText, completed: false};
			try {
				const response = await addTodo(todo);
				setTodos([...todos, response]);
				setNewTodo('');
			} catch (error) {
				console.error('Failed to add todo:', error);
			}
		}
	};

	const handleDeleteTodo = async (id) => {
    try {
      await deleteTodo(id);
      setTodos(todos.filter(todo => todo.id !== id));
    } catch (error) {
      console.error('Failed to delete todo:', error);
    }
  };

	const handleUpdateTodo = async (id, updatedValues) => {
		try {
			const response = await updateTodo(updatedValues);
			const updatedTodos = todos.map(todo => 
				todo.id === id ? response : todo
			);
			setTodos(updatedTodos);
		} catch (error) {
			console.error("Error updating todo:", error);
		}
};

	const toggleCompletion = (index) => {
		const todo = todos.find(t => t.id === index);
		if (todo) {
				const updatedTodo = { ...todo, completed: !todo.completed };
				handleUpdateTodo(index, updatedTodo);
		}
		console.log(todos);
	};

	return (
		<div className="flex flex-col items-center h-screen w-full py-20 space-y-8 min-w-[300px]">
      <h1 className="text-2xl text-center">TodoList App</h1>
			<div className="flex flex-row items-center space-x-4">
				<input type="text" className="border rounded-md px-2 py-1 w-3/4" placeholder="Enter a task" value={newTodo} onChange={e => setNewTodo(e.target.value)}/>
				<button className="border rounded-md px-2 py-1 bg-blue-500 text-white hover:bg-blue-600" onClick={handleAddTodo}>Add</button>
			</div>
      <ul className="list-decimal list-inside h-1/2 w-1/2 min-w-[300px] max-w-[500px] overflow-y-auto border rounded-md">
        {todos.map((todo) => (
          <li key={todo.id} className="flex justify-between items-center py-2 px-4 border-b last:border-0 hover:bg-gray-50 w-full">
						<div className="flex items-center flex-grow">
							<button 
								className={`mr-4 p-[5px] ${todo.completed ? 'bg-green-500' : 'bg-red-500'} rounded`} 
								onClick={() => toggleCompletion(todo.id)}
							/>
							<span className={`${todo.completed ? 'line-through' : ''} break-all whitespace-normal`}>{todo.text}</span>
  					</div>
						<button className="px-2 py-1 bg-gray-200 text-white rounded hover:bg-red-400" onClick={() => handleDeleteTodo(todo.id)}>
							Delete
						</button>
					</li>
        ))}
      </ul>
    </div>
	);
}

export default App

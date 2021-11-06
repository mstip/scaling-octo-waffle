function onCheck(index) {
	const params = new URLSearchParams();
	params.append('todoIndex', index);
	axios.post('/toggleTodo', params);
	setTimeout(() =>	window.location.href ='/', 100)
}
function onDelete(index) {
	const params = new URLSearchParams();
	params.append('todoIndex', index);
	axios.post('/deleteTodo', params);
	setTimeout(() =>	window.location.href ='/', 100)
}

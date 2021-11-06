import React, { useEffect, useState } from "react";
import TodoItem from "./TodoItem";
import TodoModel from "./TodoModel";

export default function TodoList() {
    const [todos, setTodos] = useState([new TodoModel]);


    useEffect(() => {
        fetch('https://3001-coffee-mastodon-y32x2pog.ws-eu18.gitpod.io/todos')
            .then(response => response.json())
            .then(data => setTodos(data));
    }, []);


    function onNewTodo(event: React.KeyboardEvent<HTMLInputElement>) {
        if (event.key !== 'Enter') {
            return;
        }
        const newTodo = new TodoModel(event.currentTarget.value);
        event.currentTarget.value = '';
        fetch('https://3001-coffee-mastodon-y32x2pog.ws-eu18.gitpod.io/todos', {
            method: 'POST',
            body: JSON.stringify(newTodo),
            headers: {
                'Content-Type': 'application/json'
            },
        }).then(response => {
            if (!response.ok) {
                return
            }
            setTodos([...todos, newTodo]);
        });

    }

    function onToggleDone(index: number) {
        todos[index].done = !todos[index].done;

        fetch('https://3001-coffee-mastodon-y32x2pog.ws-eu18.gitpod.io/todos/' + index , {
            method: 'PUT',
            body: JSON.stringify(todos[index]),
            headers: {
                'Content-Type': 'application/json'
            },
        }).then(response => {
            if (!response.ok) {
                return
            }
            setTodos([...todos]);
        });

    }

    function onDeleteTodo(index: number): void {
        fetch('https://3001-coffee-mastodon-y32x2pog.ws-eu18.gitpod.io/todos/' + index , {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
        }).then(response => {
            if (!response.ok) {
                return
            }
            todos.splice(index, 1);
            setTodos([...todos]);
        });
       
    }

    const todoItems = todos.map((item: TodoModel, index) =>
        <TodoItem
            key={index}
            index={index}
            text={item.text}
            done={item.done}
            onToggleDone={onToggleDone}
            onDeleteTodo={onDeleteTodo} />
    )

    return (
        <div>
            <h1>TODOLIST</h1>
            <input type="text" placeholder="newtodo" onKeyUp={onNewTodo} />
            <ul>
                {todoItems}
            </ul>
        </div>
    );
}


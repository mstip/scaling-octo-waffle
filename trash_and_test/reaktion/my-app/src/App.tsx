import React from 'react';
import { BrowserRouter, Link, Route, Switch } from 'react-router-dom';
import './App.css';
import TodoList from './todolist/TodoList';

function App() {
  return (
    <BrowserRouter>
    <div>
      <nav>
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/todos">Todo</Link>
          </li>
        </ul>
      </nav>

      {/* A <Switch> looks through its children <Route>s and
          renders the first one that matches the current URL. */}
      <Switch>
      <Route path="/todos">
          <TodoList />
        </Route>
        <Route path="/">
          <Home />
        </Route>
      
      </Switch>
    </div>
  </BrowserRouter>
  );
}

function Home() {
  return <h1>Home</h1>;
}

export default App;

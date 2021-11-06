import logo from './logo.svg';
import './App.css';
import { useState } from 'react';


function Title(props) {
  return <h1>{props.text}</h1>;
}

function Btn(props) {
  return <button onClick={() => props.onClick()}>{props.value}</button>
}

function App() {

  const [counter, setCounter] = useState(0);


  return (
    <div>
      <Title text="counter" />
      <button onClick={() => setCounter(counter + 1)}>{counter}</button>
      <button onClick={() => setCounter(0)}>Reset</button>
      <Btn value={counter} onClick={() => setCounter(23)}/>
      </div>

  );
}

export default App;

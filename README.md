import './App.css';
import { useState } from 'react';

function App() {
  const { status, setStatus } = useState('Disconnected')

  // In functional React component

  let socket = new WebSocket("ws://localhost:8080/ws")

  const wsConect = () => {
    socket.onopen = function (e) {
      setStatus('conected')
    }

    socket.onmessage = function (event) {
      console.log(event.data);
    }

    socket.onclose = function (event) {
      if (event.wasClean) {
        setStatus('Disconnected')
      } else {
        setStatus('Disconnected unexpectedly')
      }
    }
  }


  return (
    <div className="App">
      <p>
        Web Socket
      </p>

      <button onClick={wsConect}>conect</button>
      <button>disconect</button>
    </div>
  );
}

export default App;

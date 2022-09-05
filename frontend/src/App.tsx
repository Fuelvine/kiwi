import {useState} from 'react';
import logo from './assets/images/Fuelvine.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import {UpdateCode} from "../wailsjs/go/main/Telemetry";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [codeText, setCodeText] = useState("None");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    function greet() {
        Greet(name).then(updateResultText);
    }

    function updateCode() {
        UpdateCode().then(setCodeText)
    }

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo"/>
            <div id="result" className="result">{codeText}</div>
            <div id="input" className="input-box">
                <button className="btn" onClick={updateCode}>Update</button>
            </div>
        </div>
    )
}

export default App

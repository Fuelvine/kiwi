import {useState} from 'react';
import logo from './assets/images/Fuelvine.png';
import './App.css';
import {GetCode, GetSpeed} from "../wailsjs/go/main/Telemetry";
import {EventsOn} from "../wailsjs/runtime";

function App() {
    const [codeText, setCodeText] = useState("None");
    const [speed, setSpeed] = useState(0);

    function updateCode() {
        GetCode().then(setCodeText)
    }

    function updateSpeed(){
        GetSpeed().then(setSpeed)
    }

    EventsOn("event", updateCode)
    EventsOn("car", updateSpeed)

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo"/>
            <div id="result" className="result">{codeText}</div>
            <div id="result" className="result">{speed}</div>
            <div id="input" className="input-box">
                <button className="btn" onClick={updateCode}>Update</button>
            </div>
        </div>
    )
}

export default App

import {useState} from 'react';
import logo from './assets/images/Fuelvine.png';
import './App.css';
import {GetCarTelemetry} from "../wailsjs/go/telemetry/Telemetry";
import {EventsOn} from "../wailsjs/runtime";
import * as packets from "./Packets";

function App() {
    const [codeText, setCodeText] = useState("None");
    const [speed, setSpeed] = useState(0);
    const [telemetry, setTelemetry] = useState();

    async function updateTelemetry() {
        let carTelemetry: packets.CarTelemetryData = await GetCarTelemetry()
        console.log(carTelemetry.Speed)
        console.log(carTelemetry.TyresPressure)
    }

    EventsOn("event", setCodeText)
    EventsOn("car", setSpeed)

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo"/>
            <div id="result" className="result">{codeText}</div>
            <div id="result" className="result">{speed}</div>
            <div id="result" className="result">{telemetry}</div>
            <div id="input" className="input-box">
                <button className="btn" onClick={updateTelemetry}>Update</button>
            </div>
        </div>
    )
}

export default App

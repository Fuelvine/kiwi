import {useState} from 'react';
import './App.css';
import {FramesJSON} from "../wailsjs/go/telemetry/Telemetry";
import {EventsOn} from "../wailsjs/runtime";
import * as packets from "./Packets";
import {SpeedChart} from "./charts/speed";

function App() {
    const [codeText, setCodeText] = useState("None");
    const [speed, setSpeed] = useState(0);
    const [telemetry, setTelemetry] = useState();

    async function updateTelemetry() {
        // let carTelemetry: packets.CarTelemetryData = await GetCarTelemetry()
        // console.log(carTelemetry.Speed)
        // console.log(carTelemetry.TyresPressure)
    }

    EventsOn("event", setCodeText)
    EventsOn("car", setSpeed)

    return (
        <div id="App">
            {/*<div id="result" className="result">{codeText}</div>*/}
            {/*<div id="result" className="result">{speed}</div>*/}
            {/*<button onClick={FramesJSON}>*/}
            {/*     Export frames*/}
            {/*</button>*/}
            <SpeedChart/>
        </div>
    )
}

export default App

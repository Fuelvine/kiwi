export interface CarTelemetryData {
    Speed: number;
    Throttle: number;
    Steer: number;
    Brake: number;
    Clutch: number;
    Gear: number;
    EngineRPM: number;
    DRS: number;
    RevLightsPercent: number;
    RevLightsBitValue: number;
    BrakesTemperature: number[];
    TyresSurfaceTemperature: number[];
    TyresInnerTemperature: number[];
    EngineTemperature: number;
    TyresPressure: number[];
    SurfaceType: number[];
}
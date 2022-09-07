import React, {useState} from 'react';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import { Line } from 'react-chartjs-2';
import {EventsOn} from "../../wailsjs/runtime";

export function SpeedChart() {

    ChartJS.register(
        CategoryScale,
        LinearScale,
        PointElement,
        LineElement,
        Title,
        Tooltip,
        Legend
    );

    let options = {
        // maintainAspectRatio: false,
        responsive: true,
        plugins: {
            legend: {
                position: 'top' as const,
            },
            title: {
                display: true,
                text: 'Chart.js Line Chart',
            },
        },
    };

    const [msLabel, setLabel] = useState([] as string[])
    const [speedData, setData] = useState([] as number[])

    let data = {
        labels: msLabel,
        datasets: [
            {
                label: 'Dataset 1',
                data: speedData,
                borderColor: 'rgb(255, 99, 132)',
                backgroundColor: 'rgba(255, 99, 132, 0.5)',
            },
        ],
    };

    function updateChart(speed: number) {
        if (speed < 3) {
            return
        }
        setLabel(msLabel.concat([" "]))
        setData(speedData.concat([speed]))
    }

    EventsOn("car",  updateChart)
    return <Line options={options} data={data} />;
}

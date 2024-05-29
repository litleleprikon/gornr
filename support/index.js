const svg = document.querySelector('.line-chart')

const lineLimit = 1.737;
const delayVolume = 0.65;

const t = Array.from(Array(100).keys());
const values = Array.from(t.map((x) => {
    let val = Math.sin(x / (Math.PI * 2));
    val *= lineLimit;
    // if (Math.abs(val) > lineLimit) {
    //     val = lineLimit*(val/Math.abs(val));
    // }
    return val;
}))

const delayValues = Array.from(t.map((x) => {
    let val = Math.sin(x / (Math.PI * 2) + 1);
    val *= lineLimit * delayVolume;
    // if (Math.abs(val) > lineLimit) {
    //     val = lineLimit*(val/Math.abs(val));
    // }
    return val;
}))

const resultValues = Array.from(delayValues.map((e, i) => {
    let val = e + values[i];
    val *= lineLimit / (lineLimit + (lineLimit * delayVolume));
    return val;
}));

new chartXkcd.Line(svg, {
    title: 'Delay',
    xLabel: 'Time, Ms',
    yLabel: 'Voltage',
    data: {
        labels: t,
        datasets: [
            {
                label: 'Original',
                data: values,
            },
            {
                label: 'Delay',
                data: delayValues,
            },
            {
                label: 'Result',
                data: resultValues,
            },
        ],
    },
    options: {}
});
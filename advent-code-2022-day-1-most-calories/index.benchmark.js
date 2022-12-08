const Benchmark = require('benchmark');
const suite = new Benchmark.Suite;
const { getMostCalories, getMostCaloriesFile } = require('./index');

// benchmarks
suite
    .add('#getMostCalories', function() {
        const sample = `1000
        2000
        3000
    
        4000
    
        5000
        6000
    
        7000
        8000
        9000
    
        10000`;
        getMostCalories(sample);
    })
    // removing due to getting "too many files" errors
    // .add('#getMostCaloriesFile', async function() {
    //     await getMostCaloriesFile('./sample.txt');
    // })
    // add listeners
    .on('cycle', function(event) {
        console.log(String(event.target))
    })
    .on('complete', function() {
        console.log('The fastest method is ' + this.filter('fastest').map('name'))
    })
    // run async or not
    .run({ 'async': true });
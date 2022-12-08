const test = require('ava');
const { getMostCalories, getMostCaloriesFile } = require('./index')

test('#getMostCalories when given a valid sample', t => {
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
    const result = getMostCalories(sample);
    const expected = 24000;
    t.is(result, expected);
});

test('#getMostCaloriesFile when given a valid sample file', async t => {
    const result = await getMostCaloriesFile('./advent-code-2022-day-1-most-calories/sample.txt');
    const expected = 24000;
    t.is(result, expected);
});
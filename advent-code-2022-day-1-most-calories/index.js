/**
 * Finds the person with the most snack calories
 * Input will be a string with every line has a snack value of calories in a number
 * If there's a blank line, then that separates a person from one another
 * This will find the person that has the most calories with their snacks
 * @param {String} file 
 * @returns {Number}
 */
const getMostCalories = (input) => {
    let mostCalories = 0;
    // split it by new lines, if there's a gap, it separates people
    const peopleWithItems = input.split('\n\n');
    // each single line separates their snacks, add them up
    // and if their calorie count exceeds the current top, replace it
    peopleWithItems.forEach((items) => {
        const itemsList = items.split('\n');
        const total = itemsList.reduce((cur, next) => {
            return cur + Number(next);
        }, 0);
        if (total > mostCalories) {
            mostCalories = total;
        }
    });
    return mostCalories;
};

const fs = require('fs');
const readline = require('readline');
/**
 * Finds the person with the most snack calories
 * Input will be a file with every line has a snack value of calories in a number
 * If there's a blank line, then that separates a person from one another
 * This will find the person that has the most calories with their snacks
 * @param {String} file 
 * @returns {Number}
 */
const getMostCaloriesFile = (file) => {
    let mostCalories = 0;
    let currentPersonCalories = 0;
    // return a promise so we can a final value from the stream
    return new Promise((resolve, reject) => {
        const rl = readline.createInterface({
            input: fs.createReadStream(file),
            crlfDelay: Infinity
        });
        rl.on('line', (line) => {
            // convert the stream buffer to a string
            const input = Buffer.from(line).toString();
            // we've reached the end of their snacks, determine if their calorie count is the top, and reset
            if (input.length === 0) {
                if (currentPersonCalories > mostCalories) {
                    mostCalories = currentPersonCalories;
                    currentPersonCalories = 0;
                }
            } else {
                currentPersonCalories += Number(input);
            }
        });
        rl.on('error', (e) => {
            reject(e);
        });
        rl.on('close', () => {
            resolve(mostCalories);
        });
    });
};

module.exports = {
    getMostCalories,
    getMostCaloriesFile
};
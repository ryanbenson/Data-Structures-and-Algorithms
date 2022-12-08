const getMostCalories = (input) => {
    let mostCalories = 0;
    const peopleWithItems = input.split('\n\n');
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
console.log(result);

const fs = require('fs');
const readline = require('readline');
const getMostCaloriesFile = (file) => {
    let mostCalories = 0;
    let currentPersonCalories = 0;
    return new Promise((resolve, reject) => {
        const rl = readline.createInterface({
            input: fs.createReadStream(file),
            crlfDelay: Infinity
        });
        rl.on('line', (line) => {
            const input = Buffer.from(line).toString();
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

(async () => {
    const result = await getMostCaloriesFile('sample.txt');
    console.log(result);
})();
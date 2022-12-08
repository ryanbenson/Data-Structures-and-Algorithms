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
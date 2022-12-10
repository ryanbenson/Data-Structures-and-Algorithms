/**
 * Finds the total calories of the top 3 people with the most snack calories
 * Input will be a string with every line has a snack value of calories in a number
 * If there's a blank line, then that separates a person from one another
 * This will find the person that has the most calories with their snacks
 * @param {String} input to process 
 * @returns {Number}
 */
const getCalories = (input) => {
    const calories = [];
    // split it by new lines, if there's a gap, it separates people
    const peopleWithItems = input.split('\n\n');
    // each single line separates their snacks, add them up
    // and if their calorie count exceeds the current top, replace it
    peopleWithItems.forEach((items) => {
        const itemsList = items.split('\n');
        const total = itemsList.reduce((cur, next) => {
            return cur + Number(next);
        }, 0);
        calories.push(total);
    });
    calories.sort((a, b) => a - b);
    // get the top 3, which are at the end of the array, then add them up
    const top3Calories = calories.slice(-3);
    const totalCalories = top3Calories.reduce((cur, next) => {
        return cur + next;
    }, 0);
    return totalCalories;
};


module.exports = {
    getCalories
};

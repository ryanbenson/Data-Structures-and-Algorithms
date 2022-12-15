/**
 * Finds the sum of all matching items
 * Each row is 1 container. Split the line in half,
 * that represents 2 compartments of that container,
 * finds the item that is in both compartments, then
 * finds the "value" of that itme.
 * Returns sum of all item values
 *
 * Lowercase item types a through z have priorities 1 through 26.
 * Uppercase item types A through Z have priorities 27 through 52.
 *
 * @param {String} input to process
 * @returns {Number}
 */
const getItemSum = (input) => {
    let totalValue = 0;
    const containers = input.split('\n');
    containers.forEach((container) => {
        const [first, second] = getContainerCompartments(container);
        const intersection = first.find((item) => second.includes(item));
        const value = getValue(intersection);
        totalValue += value;
    });
    return totalValue;
};

/**
 * Converts a container into two compartments
 * splits the string in half, then converts them into array of strings
 * @param {String} container 
 * @returns {Array}
 */
const getContainerCompartments = (container) => {
    const len = container.length;
    const middle = Math.floor(len / 2);
    const first = container.substring(0, middle);
    const second = container.substring(middle);
    return [first.split(''), second.split('')];
};

// Mapping each letter to their value
const letterValueMap = {
    a: 1,
    b: 2,
    c: 3,
    d: 4,
    e: 5,
    f: 6,
    g: 7,
    h: 8,
    i: 9,
    j: 10,
    k: 11,
    l: 12,
    m: 13,
    n: 14,
    o: 15,
    p: 16,
    q: 17,
    r: 18,
    s: 19,
    t: 20,
    u: 21,
    v: 22,
    w: 23,
    x: 24,
    y: 25,
    z: 26,
    A: 27,
    B: 28,
    C: 29,
    D: 30,
    E: 31,
    F: 32,
    G: 33,
    H: 34,
    I: 35,
    J: 36,
    K: 37,
    L: 38,
    M: 39,
    N: 40,
    O: 41,
    P: 42,
    Q: 43,
    R: 44,
    S: 45,
    T: 46,
    U: 47,
    V: 48,
    W: 49,
    X: 50,
    Y: 51,
    Z: 52,
}

/**
 * Provides the value of a letter
 * @param {String} letter 
 * @returns {Number}
 */
const getValue = (letter) => {
    return letterValueMap[letter]
};

module.exports = {
    getItemSum
};

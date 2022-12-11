/**
 * Finds the total score from playing rock paper scissors
 * Input will be a string with every line with a code representing your opponent's move
 * then your move:
 * A Y
 * B X
 * C Z
 * 
 * A=rock, B=paper, C=scissors, X=rock, Y=paper, Z=scissors
 * Scoring:
 * Per shape = 1 for Rock, 2 for Paper, and 3 for Scissors
 * Outcome = 0 if you lost, 3 if the round was a draw, and 6 if you won
 * @param {String} input to process 
 * @returns {Number}
 */
const getScore = (input) => {
    let score = 0;
    const games = input.split('\n');
    games.forEach((moves) => {
        // moves are split by a space
        const [opponent, player] = moves.split(' ');
        const moveScore = getMoveScore(player);
        const outcomeScore = getOutcomeScore(opponent, player);
        const roundScore = moveScore + outcomeScore;
        score += roundScore;
    });
    return score;
};

/**
 * Determines the move score for the match
 * Scoring:
 * Per shape = 1 for Rock, 2 for Paper, and 3 for Scissors
 * @param {String} player move
 * @returns {Number} score from the move
 */
const getMoveScore = (player) => {
    if (player === 'X') {
        return 1;
    }
    if (player === 'Y') {
        return 2;
    }
    if (player === 'Z') {
        return 3;
    }
};

/**
 * Determines the outcome score using rock paper scissors rules
 * Scoring:
 * Outcome = 0 if you lost, 3 if the round was a draw, and 6 if you won
 * @param {String} opponent move
 * @param {String} player move
 * @returns {Number} score from the outcome
 */
const getOutcomeScore = (opponent, player) => {
    if (opponent === 'A' && player === 'X') {
        return 3;
    }
    if (opponent === 'B' && player === 'Y') {
        return 3;
    }
    if (opponent === 'C' && player === 'Z') {
        return 3;
    }
    if (opponent === 'A' && player === 'Y') {
        return 6;
    }
    if (opponent === 'B' && player === 'Z') {
        return 6;
    }
    if (opponent === 'C' && player === 'X') {
        return 6;
    }
    return 0;
};

module.exports = {
    getScore
};

/**
 * Finds the total score from playing rock paper scissors
 * Input will be a string with every line with a code representing your opponent's move
 * then the other is if you win/lose/tie, determining what move to make:
 * A Y
 * B X
 * C Z
 * 
 * A=rock, B=paper, C=scissors
 * X=lose, Y=draw, Z=win
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
        const [opponentMove, playerOutcome] = moves.split(' ');
        const playerMove = getPlayerMove(opponentMove, playerOutcome);
        const moveScore = getMoveScore(playerMove);
        const outcomeScore = getOutcomeScore(playerOutcome);
        const roundScore = moveScore + outcomeScore;
        score += roundScore;
    });
    return score;
};

/**
 * Gets what move the player needs based on the outcome
 * opponent's moves: A=rock, B=paper, C=scissors
 * outcome: X=lose, Y=draw, Z=win
 * @param {String} opponentMove
 * @param {String} playerOutcome
 * @returns {String} player's move
 */
const getPlayerMove = (opponentMove, playerOutcome) => {
    const winMap = {
        A: 'paper',
        B: 'scissors',
        C: 'rock'
    };
    const drawMap = {
        A: 'rock',
        B: 'paper',
        C: 'scissors'
    };
    const loseMap = {
        A: 'scissors',
        B: 'rock',
        C: 'paper'
    };
    // player needs to lose
    if (playerOutcome === 'X') {
        return loseMap[opponentMove];
    }
    // player needs to draw
    if (playerOutcome === 'Y') {
        return drawMap[opponentMove];
    }
    // player needs to win
    return winMap[opponentMove];
};

/**
 * Determines the move score for the match
 * Scoring:
 * Per shape = 1 for Rock, 2 for Paper, and 3 for Scissors
 * @param {String} player move
 * @returns {Number} score from the move
 */
const getMoveScore = (player) => {
    if (player === 'rock') {
        return 1;
    }
    if (player === 'paper') {
        return 2;
    }
    if (player === 'scissors') {
        return 3;
    }
};

/**
 * Determines the outcome score using rock paper scissors rules
 * Scoring:
 * Outcome = 0 if you lost, 3 if the round was a draw, and 6 if you won
 * X=lose, Y=draw, Z=win
 * @param {String} outcome
 * @returns {Number} score from the outcome
 */
const getOutcomeScore = (outcome) => {
    // draw
    if (outcome === 'Y') {
        return 3;
    }
    // win
    if (outcome === 'Z') {
        return 6;
    }
    // lose
    return 0;
};

module.exports = {
    getScore
};
